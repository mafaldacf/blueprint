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

	// TODO:
	// executeTicket @ /execute/collected
}

type DashboardImpl struct {
	assuranceService AssuranceService
	basicService     BasicService
	configService    ConfigService
	contactsService  ContactsService
	paymentService   PaymentService
	preserveService  PreserveService
	orderService     OrderService
}

func NewDashboardImpl(ctx context.Context,
	assuranceService AssuranceService,
	basicService BasicService,
	configService ConfigService,
	contactsService ContactsService,
	paymentService PaymentService,
	preserveService PreserveService,
	orderService OrderService,
) (Dashboard, error) {
	return &DashboardImpl{
		assuranceService: assuranceService,
		basicService:     basicService,
		configService:    configService,
		contactsService:  contactsService,
		paymentService:   paymentService,
		preserveService:  preserveService,
		orderService:     orderService,
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
