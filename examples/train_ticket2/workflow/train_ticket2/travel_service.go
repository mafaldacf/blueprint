package train_ticket2

import (
	"context"
	"fmt"
	"time"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
	"go.mongodb.org/mongo-driver/bson"
)

type TravelService interface {
	GetTripAllDetailInfo(ctx context.Context, gtdi TripAllDetailInfo) (TripAllDetail, error)
	CreateTrip(ctx context.Context, info TravelInfo) (Trip, error)
	DeleteTrip(ctx context.Context, tripID string) error
}

type TravelServiceImpl struct {
	basicService BasicService
	seatService  SeatService
	travelDB     backend.NoSQLDatabase
}

func NewTravelServiceImpl(ctx context.Context, basicService BasicService, seatService SeatService, travelDB backend.NoSQLDatabase) (TravelService, error) {
	return &TravelServiceImpl{basicService: basicService, seatService: seatService, travelDB: travelDB}, nil
}

func (t *TravelServiceImpl) CreateTrip(ctx context.Context, info TravelInfo) (Trip, error) {
	collection, err := t.travelDB.GetCollection(ctx, "travel_db", "trip")
	if err != nil {
		return Trip{}, err
	}

	filter := bson.D{{Key: "TripID", Value: info.TripID}}
	res, err := collection.FindOne(ctx, filter)
	if err != nil {
		return Trip{}, err
	}
	var trip Trip
	ok, err := res.One(ctx, &trip)
	if err != nil {
		return Trip{}, err
	}
	if ok {
		return Trip{}, fmt.Errorf("trip (%s) already exists", info.TripID)
	}
	
	trip = Trip{
		TripID:              info.TripID,
		TrainTypeName:       info.TrainTypeName,
		RouteID:             info.RouteID,
		StartStationName:    info.StartStationName,
		StationsName:        info.StationsName,
		TerminalStationName: info.TerminalStationName,
		StartTime:           info.StartTime,
		EndTime:             info.EndTime,
	}
	err = collection.InsertOne(ctx, trip)
	if err != nil {
		return Trip{}, err
	}
	return trip, err
}

func (t *TravelServiceImpl) DeleteTrip(ctx context.Context, tripID string) error {
	collection, err := t.travelDB.GetCollection(ctx, "travel_db", "trip")
	if err != nil {
		return err
	}

	filter := bson.D{{Key: "TripID", Value: tripID}}
	return collection.DeleteOne(ctx, filter)
}

func (t *TravelServiceImpl) GetTripAllDetailInfo(ctx context.Context, gtdi TripAllDetailInfo) (TripAllDetail, error) {
	collection, err := t.travelDB.GetCollection(ctx, "travel_db", "trip")
	if err != nil {
		return TripAllDetail{}, err
	}
	filter := bson.D{{Key: "TripID", Value: gtdi.TripID}}
	res, err := collection.FindOne(ctx, filter)
	if err != nil {
		return TripAllDetail{}, err
	}
	var trip Trip
	ok, err := res.One(ctx, &trip)
	if err != nil {
		return TripAllDetail{}, err
	}
	if !ok {
		return TripAllDetail{}, fmt.Errorf("trip (%s) not found", gtdi.TripID)
	}

	// get tickets

	startPlaceName := gtdi.From
	endPlaceName := gtdi.To
	departureTime := gtdi.TravelDate
	if !afterToday(departureTime) {
		return TripAllDetail{}, fmt.Errorf("departure time (%s) not valid", departureTime)
	}
	travelQuery := Travel{
		Trip:          trip,
		StartPlace:    startPlaceName,
		EndPlace:      endPlaceName,
		DepartureTime: departureTime,
	}
	travelResult, err := t.basicService.QueryForTravel(ctx, travelQuery)
	if err != nil {
		return TripAllDetail{}, err
	}

	// set trip response
	route := travelResult.Route
	stationList := route.Stations

	firstClassTotalNum := travelResult.TrainType.ComfortClass
	secondClassTotalNum := travelResult.TrainType.EconomyClass

	first, err := t.seatService.GetLeftTicketOfInterval(ctx, SeatRequest{departureTime, trip.TripID, startPlaceName, endPlaceName, 1, firstClassTotalNum, stationList})
	if err != nil {
		return TripAllDetail{}, err
	}

	second, err := t.seatService.GetLeftTicketOfInterval(ctx, SeatRequest{departureTime, trip.TripID, startPlaceName, endPlaceName, 2, secondClassTotalNum, stationList})
	if err != nil {
		return TripAllDetail{}, err
	}

	indexStart := indexOf(route.Stations, startPlaceName)
	indexEnd := indexOf(route.Stations, endPlaceName)
	distanceStart := route.Distances[indexStart] - route.Distances[0]
	distanceEnd := route.Distances[indexEnd] - route.Distances[0]
	trainType := travelResult.TrainType
	minutesStart := 60 * distanceStart / trainType.AvgSpeed
	minutesEnd := 60 * distanceEnd / trainType.AvgSpeed

	start, err := time.ParseInLocation(CALENDAR_LAYOUT, trip.StartTime, time.Local)
	if err != nil {
		return TripAllDetail{}, err
	}
	calendarStart := start.Add(time.Duration(minutesStart) * time.Minute)
	calendarEnd := start.Add(time.Duration(minutesEnd) * time.Minute)

	startTime := calendarStart.Format(CALENDAR_LAYOUT)
	endTime := calendarEnd.Format(CALENDAR_LAYOUT)

	tripResponse := TripResponse{
		TripID:               trip.TripID,
		TrainTypeName:        trip.TrainTypeName,
		ConfortClass:         first,
		EconomyClass:         second,
		PriceForConfortClass: travelResult.Prices["confortClass"],
		PriceForEconomyClass: travelResult.Prices["economyClass"],
		StartStation:         startPlaceName,
		TerminalStation:      endPlaceName,
		StartTime:            startTime,
		EndTime:              endTime,
	}

	gtdr := TripAllDetail{
		trip:         trip,
		tripResponse: tripResponse,
	}

	return gtdr, nil
}

func afterToday(dateStr string) bool {
	d, err := time.ParseInLocation(CALENDAR_LAYOUT, dateStr, time.Local)
	if err != nil {
		return false
	}

	now := time.Now()
	yearA, monthA, dayA := now.Date()
	yearB, monthB, datB := d.Date()

	if yearA > yearB {
		return false
	} else if yearA == yearB {
		if int(monthA) > int(monthB) {
			return false
		} else if int(monthA) == int(monthB) {
			return dayA <= datB
		} else {
			return true
		}
	} else {
		return true
	}
}
