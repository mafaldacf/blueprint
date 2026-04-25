package tests

import (
	"context"
	"testing"

	"github.com/blueprint-uservices/blueprint/examples/trainticket/workflow/trainticket"
	"github.com/blueprint-uservices/blueprint/runtime/core/registry"
	"github.com/blueprint-uservices/blueprint/runtime/plugins/simplenosqldb"
	"github.com/blueprint-uservices/blueprint/runtime/plugins/simplequeue"
	"github.com/stretchr/testify/assert"
)

var preserveServiceRegistry = registry.NewServiceRegistry[trainticket.PreserveService]("preserve_service")

func init() {
	preserveServiceRegistry.Register("local", func(ctx context.Context) (trainticket.PreserveService, error) {
		assuranceDB, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		assuranceService, err := trainticket.NewAssuranceServiceImpl(ctx, assuranceDB)
		if err != nil {
			return nil, err
		}

		stationDB, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		stationService, err := trainticket.NewStationServiceImpl(ctx, stationDB)
		if err != nil {
			return nil, err
		}
		trainDB, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		trainService, err := trainticket.NewTrainServiceImpl(ctx, trainDB)
		if err != nil {
			return nil, err
		}
		routeDB, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		routeService, err := trainticket.NewRouteServiceImpl(ctx, routeDB)
		if err != nil {
			return nil, err
		}
		priceDB, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		priceService, err := trainticket.NewPriceServiceImpl(ctx, priceDB)
		if err != nil {
			return nil, err
		}
		basicService, err := trainticket.NewBasicServiceImpl(ctx, stationService, trainService, routeService, priceService)
		if err != nil {
			return nil, err
		}

		orderDB, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		orderService, err := trainticket.NewOrderServiceImpl(ctx, orderDB)
		if err != nil {
			return nil, err
		}

		configDB, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		configService, err := trainticket.NewConfigServiceImpl(ctx, configDB)
		if err != nil {
			return nil, err
		}
		seatService, err := trainticket.NewSeatServiceImpl(ctx, orderService, configService)
		if err != nil {
			return nil, err
		}

		travelDB, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		travelService, err := trainticket.NewTravelServiceImpl(ctx, basicService, seatService, routeService, trainService, travelDB)
		if err != nil {
			return nil, err
		}

		consignPriceDB, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		consignPriceService, err := trainticket.NewConsignPriceServiceImpl(ctx, consignPriceDB)
		if err != nil {
			return nil, err
		}
		consignDB, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		consignService, err := trainticket.NewConsignServiceImpl(ctx, consignPriceService, consignDB)
		if err != nil {
			return nil, err
		}

		contactsDB, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		contactsService, err := trainticket.NewContactsServiceImpl(ctx, contactsDB)
		if err != nil {
			return nil, err
		}

		trainFoodDB, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		trainFoodService, err := trainticket.NewTrainFoodServiceImpl(ctx, trainFoodDB)
		if err != nil {
			return nil, err
		}
		stationFoodDB, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		stationFoodService, err := trainticket.NewStationFoodServiceImpl(ctx, stationFoodDB)
		if err != nil {
			return nil, err
		}
		foodDB, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		foodQueue, err := simplequeue.NewSimpleQueue(ctx)
		if err != nil {
			return nil, err
		}
		foodService, err := trainticket.NewFoodServiceImpl(ctx, foodDB, foodQueue, trainFoodService, travelService, stationFoodService)
		if err != nil {
			return nil, err
		}

		userDB, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		userService, err := trainticket.NewUserServiceImpl(ctx, userDB)
		if err != nil {
			return nil, err
		}

		emailQueue, err := simplequeue.NewSimpleQueue(ctx)
		if err != nil {
			return nil, err
		}

		return trainticket.NewPreserveServiceImpl(
			ctx,
			assuranceService,
			basicService,
			consignService,
			contactsService,
			foodService,
			orderService,
			seatService,
			stationService,
			travelService,
			userService,
			emailQueue,
		)
	})
}

func TestPreserveServiceContactNotFound(t *testing.T) {
	ctx := context.Background()
	service, err := preserveServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	oti := trainticket.OrderTicketsInfo{
		AccountID:  "preserve_acc001",
		ContactsID: "nonexistent_contact",
		TripID:     "G_PRESERVE_001",
		SeatType:   2,
		Date:       "2026-05-01",
		From:       "shanghai",
		To:         "beijing",
	}
	_, err = service.Preserve(ctx, oti)
	assert.Error(t, err)
}
