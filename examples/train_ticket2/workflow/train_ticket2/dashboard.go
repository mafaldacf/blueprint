package train_ticket2

import (
	"context"
)

type Dashboard interface {
	AddAssurance(ctx context.Context, typeIndex int, orderID string) (Assurance, error)
	AddContact(ctx context.Context, contact Contact) error
	GetTravel(ctx context.Context, info Travel) (TravelResult, error)
	GetAllConfigs(ctx context.Context) ([]Config, error)
	GetAllContacts(ctx context.Context) ([]Contact, error)
	GetAllOrders(ctx context.Context) ([]Order, error)
	Pay(ctx context.Context, payment Payment) error
	PreserveTicketConfirm(ctx context.Context, oti OrderTicketsInfo) (Order, error)
	TicketCollect(ctx context.Context, orderID string) error
	TicketExecute(ctx context.Context, orderID string) error

	PayDifference(ctx context.Context, info RebookInfo) error
	Rebook(ctx context.Context, info RebookInfo) error

	QueryOrderWithAllInfo(ctx context.Context, orderID string) (Order, FoodOrder, Assurance, ConsignRecord, Delivery, error)
}

type DashboardImpl struct {
	basicService    BasicService
	configService   ConfigService
	contactsService ContactsService
	paymentService  PaymentService
	preserveService PreserveService
	executeService  ExecuteService
	cancelService   CancelService
	rebookService   RebookService

	assuranceService AssuranceService
	orderService     OrderService
	foodService      FoodService
	consignService   ConsignService
	deliveryService  DeliveryService
}

func NewDashboardImpl(ctx context.Context,
	basicService BasicService,
	configService ConfigService,
	contactsService ContactsService,
	paymentService PaymentService,
	preserveService PreserveService,
	executeService ExecuteService,
	cancelService CancelService,
	rebookService RebookService,

	assuranceService AssuranceService,
	orderService OrderService,
	foodService FoodService,
	consignService ConsignService,
	deliveryService DeliveryService,
) (Dashboard, error) {
	return &DashboardImpl{
		basicService:    basicService,
		configService:   configService,
		contactsService: contactsService,
		paymentService:  paymentService,
		preserveService: preserveService,
		executeService:  executeService,
		cancelService:   cancelService,
		rebookService:   rebookService,

		assuranceService: assuranceService,
		orderService:     orderService,
		foodService:      foodService,
		consignService:   consignService,
		deliveryService:  deliveryService,
	}, nil
}

func (d *DashboardImpl) AddAssurance(ctx context.Context, typeIndex int, orderID string) (Assurance, error) {
	return d.assuranceService.Create(ctx, typeIndex, orderID)
}

func (d *DashboardImpl) AddContact(ctx context.Context, contact Contact) error {
	return d.contactsService.CreateContacts(ctx, contact)
}

func (d *DashboardImpl) GetTravel(ctx context.Context, info Travel) (TravelResult, error) {
	return d.basicService.QueryForTravel(ctx, info)
}

func (d *DashboardImpl) GetAllConfigs(ctx context.Context) ([]Config, error) {
	return d.configService.FindAll(ctx)
}

func (d *DashboardImpl) GetAllContacts(ctx context.Context) ([]Contact, error) {
	return d.contactsService.GetAllContacts(ctx)
}

func (d *DashboardImpl) GetAllOrders(ctx context.Context) ([]Order, error) {
	return d.orderService.FindAllOrder(ctx)
}

func (d *DashboardImpl) Pay(ctx context.Context, payment Payment) error {
	return d.paymentService.Pay(ctx, payment)
}

func (d *DashboardImpl) PreserveTicketConfirm(ctx context.Context, oti OrderTicketsInfo) (Order, error) {
	return d.preserveService.Preserve(ctx, oti)
}

func (d *DashboardImpl) TicketCollect(ctx context.Context, orderID string) error {
	return d.TicketCollect(ctx, orderID)
}

func (d *DashboardImpl) TicketExecute(ctx context.Context, orderID string) error {
	return d.TicketCollect(ctx, orderID)
}

func (d *DashboardImpl) CalculateRefund(ctx context.Context, orderID string) (string, error) {
	return d.cancelService.CalculateRefund(ctx, orderID)
}

func (d *DashboardImpl) CancelOrder(ctx context.Context, orderID string, loginID string) error {
	return d.cancelService.CancelOrder(ctx, orderID, loginID)
}

func (d *DashboardImpl) PayDifference(ctx context.Context, info RebookInfo) error {
	return d.rebookService.PayDifference(ctx, info)
}
func (d *DashboardImpl) Rebook(ctx context.Context, info RebookInfo) error {
	return d.rebookService.Rebook(ctx, info)
}

func (d *DashboardImpl) QueryOrderWithAllInfo(ctx context.Context, orderID string) (Order, FoodOrder, Assurance, ConsignRecord, Delivery, error) {
	order, err := d.orderService.GetOrderById(ctx, orderID)
	if err != nil {
		return Order{}, FoodOrder{}, Assurance{}, ConsignRecord{}, Delivery{}, nil
	}
	foodOrder, err := d.foodService.FindFoodOrderByOrderId(ctx, order.ID)
	if err != nil {
		return Order{}, FoodOrder{}, Assurance{}, ConsignRecord{}, Delivery{}, nil
	}

	assurance, err := d.assuranceService.FindAssuranceByOrderId(ctx, order.ID)
	if err != nil {
		return Order{}, FoodOrder{}, Assurance{}, ConsignRecord{}, Delivery{}, nil
	}

	consign, err := d.consignService.FindByOrderId(ctx, order.ID)
	if err != nil {
		return Order{}, FoodOrder{}, Assurance{}, ConsignRecord{}, Delivery{}, nil
	}

	delivery, err := d.deliveryService.FindDelivery(ctx, order.ID)
	if err != nil {
		return Order{}, FoodOrder{}, Assurance{}, ConsignRecord{}, Delivery{}, nil
	}

	return order, foodOrder, assurance, consign, delivery, nil
}
