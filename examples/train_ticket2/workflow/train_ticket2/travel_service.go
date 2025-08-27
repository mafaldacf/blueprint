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
}

type TravelServiceImpl struct {
	basicService BasicService
	travelDB     backend.NoSQLDatabase
}

func NewTravelServiceImpl(ctx context.Context, basicService BasicService, travelDB backend.NoSQLDatabase) (TravelService, error) {
	return &TravelServiceImpl{basicService: basicService, travelDB: travelDB}, nil
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
	//stationList := route.Stations

	//firstClassTotalNum := tr.TrainType.ComfortClass
	//secondClassTotalNum := tr.TrainType.EconomyClass

	var first, second int // TODO getRestTicketNumber() for POST /api/v1/seatservice/seats/left_tickets

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

func setResponse(trip Trip, tr TravelResult, startPlaceName string, endPlaceName string, departureTime string) (TripResponse, error) {
	route := tr.Route
	//stationList := route.Stations

	//firstClassTotalNum := tr.TrainType.ComfortClass
	//secondClassTotalNum := tr.TrainType.EconomyClass

	var first, second int // TODO getRestTicketNumber() for POST /api/v1/seatservice/seats/left_tickets

	indexStart := indexOf(route.Stations, startPlaceName)
	indexEnd := indexOf(route.Stations, endPlaceName)
	distanceStart := route.Distances[indexStart] - route.Distances[0]
	distanceEnd := route.Distances[indexEnd] - route.Distances[0]
	trainType := tr.TrainType
	minutesStart := 60 * distanceStart / trainType.AvgSpeed
	minutesEnd := 60 * distanceEnd / trainType.AvgSpeed

	start, err := time.ParseInLocation(CALENDAR_LAYOUT, trip.StartTime, time.Local)
	if err != nil {
		return TripResponse{}, err
	}
	calendarStart := start.Add(time.Duration(minutesStart) * time.Minute)
	calendarEnd := start.Add(time.Duration(minutesEnd) * time.Minute)

	startTime := calendarStart.Format(CALENDAR_LAYOUT)
	endTime := calendarEnd.Format(CALENDAR_LAYOUT)

	response := TripResponse{
		TripID:               trip.TripID,
		TrainTypeName:        trip.TrainTypeName,
		ConfortClass:         first,
		EconomyClass:         second,
		PriceForConfortClass: tr.Prices["confortClass"],
		PriceForEconomyClass: tr.Prices["economyClass"],
		StartStation:         startPlaceName,
		TerminalStation:      endPlaceName,
		StartTime:            startTime,
		EndTime:              endTime,
	}

	return response, nil
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
