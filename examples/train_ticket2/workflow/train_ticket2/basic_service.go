package train_ticket2

import (
	"context"
	"fmt"
	"slices"
	"strconv"
)

type BasicService interface {
	// TODO:
	// - QueryForTravels
	// - QueryForStationId
	QueryForTravel(ctx context.Context, info Travel) (TravelResult, error)

	// REMOVE
	QueryOrderWithAllInfo(ctx context.Context, orderID string) (Order, FoodOrder, Assurance, ConsignRecord, Delivery, error)
}

type BasicServiceImpl struct {
	stationService StationService
	trainService   TrainService
	routeService   RouteService
	priceService   PriceService
	// extra
	orderService     OrderService
	foodService      FoodService
	assuranceService AssuranceService
	consignService   ConsignService
	deliveryService  DeliveryService
}

func NewBasicServiceImpl(ctx context.Context,
	stationService StationService,
	trainService TrainService,
	routeService RouteService,
	priceService PriceService,
	orderService OrderService,
	foodService FoodService,
	assuranceService AssuranceService,
	consignService ConsignService,
	deliveryService DeliveryService,
) (BasicService, error) {
	return &BasicServiceImpl{
		stationService:   stationService,
		trainService:     trainService,
		routeService:     routeService,
		priceService:     priceService,
		orderService:     orderService,
		foodService:      foodService,
		assuranceService: assuranceService,
		consignService:   consignService,
		deliveryService:  deliveryService,
	}, nil
}

func (b *BasicServiceImpl) QueryForTravel(ctx context.Context, info Travel) (TravelResult, error) {
	start := info.StartPlace
	startingPlaceExist, err := b.stationService.Exists(ctx, start)
	if err != nil {
		return TravelResult{}, err
	}
	if !startingPlaceExist {
		return TravelResult{}, fmt.Errorf("start place (%s) does not exist", start)
	}

	end := info.EndPlace
	endPlaceExist, err := b.stationService.Exists(ctx, end)
	if err != nil {
		return TravelResult{}, err
	}
	if !endPlaceExist {
		return TravelResult{}, fmt.Errorf("end place (%s) does not exist", end)
	}

	trainType, err := b.trainService.RetrieveByName(ctx, info.Trip.TrainTypeName)
	if err != nil {
		return TravelResult{}, err
	}

	routeID := info.Trip.RouteID
	route, err := b.routeService.GetRouteById(ctx, routeID)
	if err != nil {
		return TravelResult{}, err
	}

	// 1. check route list for this train
	// 2. check that the required start and arrival stations are in the list of
	// stops that are not on the route, and check that the the location of the
	// start station is before the stop
	// 3. trains that meet the above criteria are added to the return list
	var indexStart, indexEnd int
	if slices.Contains(route.Stations, start) && slices.Contains(route.Stations, end) &&
		indexOf(route.Stations, start) < indexOf(route.Stations, end) {

		indexStart = indexOf(route.Stations, start)
		indexEnd = indexOf(route.Stations, end)
	} else {
		return TravelResult{}, fmt.Errorf("station not correct in route")
	}

	priceConfig, err := b.priceService.FindByRouteIDAndTrainType(ctx, routeID, trainType.Name)
	if err != nil {
		return TravelResult{}, nil
	}
	var prices = make(map[string]string)
	distance := route.Distances[indexEnd] - route.Distances[indexStart]
	priceForEconomyClass := distance * int64(priceConfig.BasicPriceRate)
	priceForConfortClass := distance * int64(priceConfig.FirstClassPriceRate)
	prices["economyClass"] = strconv.FormatInt(priceForEconomyClass, 10)
	prices["confortClass"] = strconv.FormatInt(priceForConfortClass, 10)

	result := TravelResult{
		TrainType: trainType,
		Route:     route,
		Prices:    prices,
		Percent:   1.0,
	}
	return result, nil
}

func (b *BasicServiceImpl) QueryOrderWithAllInfo(ctx context.Context, orderID string) (Order, FoodOrder, Assurance, ConsignRecord, Delivery, error) {
	order, err := b.orderService.GetOrderById(ctx, orderID)
	if err != nil {
		return Order{}, FoodOrder{}, Assurance{}, ConsignRecord{}, Delivery{}, nil
	}
	foodOrder, err := b.foodService.FindFoodOrderByOrderId(ctx, order.ID)
	if err != nil {
		return Order{}, FoodOrder{}, Assurance{}, ConsignRecord{}, Delivery{}, nil
	}

	assurance, err := b.assuranceService.FindAssuranceByOrderId(ctx, order.ID)
	if err != nil {
		return Order{}, FoodOrder{}, Assurance{}, ConsignRecord{}, Delivery{}, nil
	}

	consign, err := b.consignService.FindByOrderId(ctx, order.ID)
	if err != nil {
		return Order{}, FoodOrder{}, Assurance{}, ConsignRecord{}, Delivery{}, nil
	}

	delivery, err := b.deliveryService.FindDelivery(ctx, order.ID)
	if err != nil {
		return Order{}, FoodOrder{}, Assurance{}, ConsignRecord{}, Delivery{}, nil
	}

	return order, foodOrder, assurance, consign, delivery, nil
}
