package trainticket

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"
)

type RebookService interface {
	PayDifference(ctx context.Context, info RebookInfo) error
	Rebook(ctx context.Context, info RebookInfo) error
}

type RebookServiceImpl struct {
	seatService          SeatService
	travelService        TravelService
	orderService         OrderService
	trainService         TrainService
	routeService         RouteService
	insidePaymentService InsidePaymentService
}

func NewRebookServiceImpl(ctx context.Context, seatService SeatService, travelService TravelService, orderService OrderService, trainService TrainService, routeService RouteService, insidePaymentService InsidePaymentService) (RebookService, error) {
	return &RebookServiceImpl{seatService: seatService, travelService: travelService, orderService: orderService, routeService: routeService, insidePaymentService: insidePaymentService}, nil
}

func (s *RebookServiceImpl) PayDifference(ctx context.Context, info RebookInfo) error {
	order, err := s.orderService.GetOrderById(ctx, info.OrderId)
	if err != nil {
		return err
	}
	if order.Status == 0 {
		return nil
	}

	gtdi := TripAllDetailInfo{
		From:       order.FromStation,
		To:         order.ToStation,
		TravelDate: info.Date,
		TripID:     info.TripId,
	}
	gtdr, err := s.travelService.GetTripAllDetailInfo(ctx, gtdi)
	if err != nil {
		return err
	}
	ticketPrice := "0"
	if info.SeatType == 2 {
		ticketPrice = gtdr.tripResponse.PriceForConfortClass
	} else if info.SeatType == 3 {
		ticketPrice = gtdr.tripResponse.PriceForEconomyClass
	}

	oldPrice := order.Price
	priceOld := oldPrice
	priceNew := ticketPrice

	var payDifference = true
	if payDifference && priceOld != "" && priceNew != "" {
		return s.updateOrder(ctx, order, info, gtdr, ticketPrice)
	}
	return fmt.Errorf("can't pay the difference, please try again")
}

func (s *RebookServiceImpl) Rebook(ctx context.Context, info RebookInfo) error {
	order, err := s.orderService.GetOrderById(ctx, info.OrderId)
	if err != nil {
		return err
	}

	if order.Status == 1 {
		return fmt.Errorf("order not suitable to rebook!")
	}

	if order.Status == ORDER_STATUS_NOT_PAID {
		return fmt.Errorf("you haven't paid the original ticket!")
	} else if order.Status == ORDER_STATUS_PAID {
		// do nothing
	} else if order.Status == ORDER_STATUS_COLLECTED {
		return fmt.Errorf("you have already collected your ticket and you can change it now.")
	} else {
		return fmt.Errorf("you can't change your ticket")
	}

	if !checkTime(order.TravelDate, order.TravelTime) {
		return fmt.Errorf("you can only change the ticket before the train start or within 2 hours after the train start")
	}

	gtdi := TripAllDetailInfo{
		From:       order.FromStation,
		To:         order.ToStation,
		TravelDate: info.Date,
		TripID:     info.TripId,
	}
	gtdr, err := s.travelService.GetTripAllDetailInfo(ctx, gtdi)
	if err != nil {
		return err
	}
	tripResponse := gtdr.tripResponse
	if info.SeatType == 2 {
		if tripResponse.ConfortClass <= 0 {
			fmt.Errorf("seat not enough")
		}
	} else {
		if tripResponse.EconomyClass == 3 && tripResponse.ConfortClass <= 0 {
			fmt.Errorf("seat not enough")
		}
	}

	//Deal with the difference, more refund less compensation
	//Return the original ticket so that someone else can book the corresponding seat
	ticketPrice := "0"
	if info.SeatType == 2 {
		ticketPrice = gtdr.tripResponse.PriceForConfortClass
	} else if info.SeatType == 3 {
		ticketPrice = gtdr.tripResponse.PriceForEconomyClass
	}
	oldPrice := order.Price
	priceOld, _ := strconv.Atoi(oldPrice)
	priceNew, _ := strconv.Atoi(ticketPrice)
	if priceOld > priceNew {
		difference := strconv.Itoa(priceOld - priceNew)
		err = s.insidePaymentService.Drawback(ctx, info.LoginId, difference)
		if err != nil {
			return fmt.Errorf("can't draw back the difference now, please try again!")
		}
		return s.updateOrder(ctx, order, info, gtdr, ticketPrice)
	} else if priceOld == priceNew {
		// do nothing
		return s.updateOrder(ctx, order, info, gtdr, ticketPrice)
	} else {
		// make up the difference
		difference := strconv.Itoa(priceNew - priceOld)
		return fmt.Errorf("please pay the different money! (%s)", difference)
	}

}

func (s *RebookServiceImpl) updateOrder(ctx context.Context, order Order, info RebookInfo, gtdr TripAllDetail, ticketPrice string) error {
	trip := gtdr.trip
	oldTripId := order.TrainNumber
	order.TrainNumber = info.TripId
	order.BoughtDate = time.Now().Format(time.DateOnly)
	order.Status = ORDER_STATUS_CHANGE
	order.Price = ticketPrice
	order.SeatClass = info.SeatType
	order.TravelDate = info.Date
	order.TravelTime = trip.StartTime

	route, err := s.routeService.GetRouteById(ctx, trip.RouteID)
	if err != nil {
		return err
	}
	trainType, err := s.trainService.RetrieveByName(ctx, trip.TrainTypeName)
	if err != nil {
		return err
	}
	stations := route.Stations
	firstClassTotalNum := trainType.ComfortClass
	secondClassTotalNum := trainType.EconomyClass
	if info.SeatType == 2 {
		ticket, _ := s.seatService.DistributeSeat(ctx, SeatRequest{
			TravelDate:   info.Date,
			TrainNumber:  order.TrainNumber,
			StartStation: order.FromStation,
			DestStation:  order.ToStation,
			SeatType:     2,
			TotalNum:     firstClassTotalNum,
			Stations:     stations,
		})
		order.SeatClass = 2
		order.SeatNumber = strconv.Itoa(ticket.SeatNo)
	} else {
		ticket, _ := s.seatService.DistributeSeat(ctx, SeatRequest{
			TravelDate:   info.Date,
			TrainNumber:  order.TrainNumber,
			StartStation: order.FromStation,
			DestStation:  order.ToStation,
			SeatType:     3,
			TotalNum:     secondClassTotalNum,
			Stations:     stations,
		})
		order.SeatClass = 3
		order.SeatNumber = strconv.Itoa(ticket.SeatNo)
	}
	if (tripGD(oldTripId) && tripGD(info.TripId)) || (!tripGD(oldTripId) && !tripGD(info.TripId)) {
		err = s.orderService.UpdateOrder(ctx, order)
		if err != nil {
			return err
		}
	}
	s.orderService.UpdateOrder(ctx, order)
	/* s.orderService.DeleteOrder(ctx, order.ID)
	s.orderService.CreateNewOrder(ctx, order) */
	return nil
}

func tripGD(tripId string) bool {
	return strings.HasPrefix(tripId, "G") || strings.HasPrefix(tripId, "D")
}

func checkTime(travelDate string, travelTime string) bool {
	now := time.Now()

	date, err := time.Parse(time.DateOnly, travelDate)
	if err != nil {
		return false
	}

	t, err := time.Parse(time.TimeOnly, travelTime)
	if err != nil {
		return false
	}

	travelDateTime := time.Date(
		date.Year(), date.Month(), date.Day(),
		t.Hour(), t.Minute(), 0, 0, time.Local,
	)

	limit := travelDateTime.Add(2 * time.Hour)

	return now.Before(limit) || now.Equal(limit)
}
