package trainticket

import (
	"context"
	"fmt"
	"slices"
	"strconv"
)

type BasicService interface {
	QueryForTravel(ctx context.Context, info Travel) (TravelResult, error)
	QueryForTravels(ctx context.Context, infos []Travel) (map[string]TravelResult, error)
	QueryForStationId(ctx context.Context, stationName string) (Station, error)
}

type BasicServiceImpl struct {
	stationService StationService
	trainService   TrainService
	routeService   RouteService
	priceService   PriceService
}

func NewBasicServiceImpl(ctx context.Context,
	stationService StationService,
	trainService TrainService,
	routeService RouteService,
	priceService PriceService,
) (BasicService, error) {
	return &BasicServiceImpl{
		stationService: stationService,
		trainService:   trainService,
		routeService:   routeService,
		priceService:   priceService,
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

func (b *BasicServiceImpl) QueryForTravels(ctx context.Context, infos []Travel) (map[string]TravelResult, error) {
	var tripInfos = make(map[string]Travel)
	var startTrips = make(map[string][]string)
	var endTrips = make(map[string][]string)
	var routeTrips = make(map[string][]string)
	var typeTrips = make(map[string][]string)
	var stationNames []string
	var trainTypeNames []string
	var routeIds []string
	var avaTrips []string
	for _, info := range infos {
		stationNames = append(stationNames, info.StartPlace)
		stationNames = append(stationNames, info.EndPlace)
		trainTypeNames = append(trainTypeNames, info.Trip.TrainTypeName)
		routeIds = append(routeIds, info.Trip.RouteID)

		tripNumber := info.Trip.TripID
		avaTrips = append(avaTrips, tripNumber)
		tripInfos[tripNumber] = info

		start := info.StartPlace
		var trips []string = startTrips[start]
		trips = append(trips, tripNumber)
		startTrips[start] = trips

		end := info.EndPlace
		trips = endTrips[end]
		trips = append(trips, tripNumber)
		endTrips[end] = trips

		routeId := info.Trip.RouteID
		trips = routeTrips[routeId]
		trips = append(trips, tripNumber)
		routeTrips[routeId] = trips

		trainTypeName := info.Trip.TrainTypeName
		trips = typeTrips[trainTypeName]
		trips = append(trips, tripNumber)
		typeTrips[trainTypeName] = trips

		/* stations, err := b.stationService.FindByIDs(ctx, stationNames)
		if err != nil {
			return nil, err
		} */

		/* for _, station := range stations {
			if value == "" { // equivalent to null check for string pointers in Java
				// station not exist
				if trips, ok := startTrips[key]; ok {
					avaTrips = removeAll(avaTrips, trips)
				}
				if trips, ok := endTrips[key]; ok {
					avaTrips = removeAll(avaTrips, trips)
				}
			}
		} */

		if len(avaTrips) == 0 {
			return nil, fmt.Errorf("no travel info available")
		}

		tts, err := b.trainService.RetrieveByNames(ctx, trainTypeNames)
		if err != nil {
			return nil, err
		}

		var trainTypeMap = make(map[string]TrainType)
		for _, t := range tts {
			trainTypeMap[t.Name] = t
		}

		for typeTripsKey, typeTripsLst := range typeTrips {
			if _, ok := trainTypeMap[typeTripsKey]; !ok {
				removeAll(avaTrips, typeTripsLst)
			}
		}

		if len(avaTrips) == 0 {
			return nil, fmt.Errorf("no travel info available")
		}

		routes, err := b.routeService.GetRouteByIds(ctx, routeIds)
		if err != nil {
			return nil, err
		}
		var routeMap = make(map[string]Route)
		for _, r := range routes {
			routeMap[r.ID] = r
		}

		for routeTripsKey, routeTripsLst := range routeTrips {
			routeId := routeTripsKey
			if _, ok := routeMap[routeId]; !ok {
				removeAll(avaTrips, routeTripsLst)
			} else {
				route := routeMap[routeId]
				trips := routeTripsLst
				for _, t := range trips {
					start := tripInfos[t].StartPlace
					end := tripInfos[t].EndPlace
					if !slices.Contains(route.Stations, start) || !slices.Contains(route.Stations, end) ||
						indexOf(route.Stations, start) >= indexOf(route.Stations, end) {
						avaTrips = remove(avaTrips, t)
					}
				}
			}
		}

		if len(avaTrips) == 0 {
			return nil, fmt.Errorf("no travel info available")
		}

		// TODO: FINALIZE
	}
	return nil, nil
}

func (b *BasicServiceImpl) QueryForStationId(ctx context.Context, stationName string) (Station, error) {
	station, err := b.stationService.FindByID(ctx, stationName)
	if err != nil {
		return Station{}, err
	}
	return station, nil
}

func removeAll(src []string, toRemove []string) []string {
	removeSet := make(map[string]struct{})
	for _, v := range toRemove {
		removeSet[v] = struct{}{}
	}

	var result []string
	for _, v := range src {
		if _, found := removeSet[v]; !found {
			result = append(result, v)
		}
	}
	return result
}

func remove(list []string, target string) []string {
	for i, v := range list {
		if v == target {
			return append(list[:i], list[i+1:]...)
		}
	}
	return list // if not found
}
