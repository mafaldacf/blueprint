package tests

import (
	"context"
	"testing"

	"github.com/blueprint-uservices/blueprint/examples/trainticket/workflow/trainticket"
	"github.com/blueprint-uservices/blueprint/runtime/core/registry"
	"github.com/blueprint-uservices/blueprint/runtime/plugins/simplenosqldb"
	"github.com/stretchr/testify/assert"
)

var contactsServiceRegistry = registry.NewServiceRegistry[trainticket.ContactsService]("contacts_service")

func init() {
	contactsServiceRegistry.Register("local", func(ctx context.Context) (trainticket.ContactsService, error) {
		db, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		return trainticket.NewContactsServiceImpl(ctx, db)
	})
}

func TestContactsServiceCreateAndFind(t *testing.T) {
	ctx := context.Background()
	service, err := contactsServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	contact := trainticket.Contact{
		ID:             "contact001",
		AccountID:      "account001",
		Name:           "John Doe",
		DocumentType:   int(trainticket.ID_CARD),
		DocumentNumber: "123456789",
		PhoneNumber:    "555-0100",
	}
	err = service.CreateContacts(ctx, contact)
	assert.NoError(t, err)

	found, err := service.FindContactsById(ctx, "contact001")
	assert.NoError(t, err)
	assert.Equal(t, "John Doe", found.Name)
	assert.Equal(t, "account001", found.AccountID)
	assert.Equal(t, "555-0100", found.PhoneNumber)
}

func TestContactsServiceFindByAccountId(t *testing.T) {
	ctx := context.Background()
	service, err := contactsServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	contact1 := trainticket.Contact{
		ID:             "contact002",
		AccountID:      "account002",
		Name:           "Alice",
		DocumentNumber: "111111111",
	}
	contact2 := trainticket.Contact{
		ID:             "contact003",
		AccountID:      "account002",
		Name:           "Bob",
		DocumentNumber: "222222222",
	}
	err = service.CreateContacts(ctx, contact1)
	assert.NoError(t, err)
	err = service.CreateContacts(ctx, contact2)
	assert.NoError(t, err)

	contacts, err := service.FindContactsByAccountId(ctx, "account002")
	assert.NoError(t, err)
	assert.Len(t, contacts, 2)
}

func TestContactsServiceGetAllContacts(t *testing.T) {
	ctx := context.Background()
	service, err := contactsServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	contact := trainticket.Contact{
		ID:        "contact004",
		AccountID: "account003",
		Name:      "Carol",
	}
	err = service.CreateContacts(ctx, contact)
	assert.NoError(t, err)

	all, err := service.GetAllContacts(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, all)
}

func TestContactsServiceModify(t *testing.T) {
	ctx := context.Background()
	service, err := contactsServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	contact := trainticket.Contact{
		ID:          "contact005",
		AccountID:   "account004",
		Name:        "Dave",
		PhoneNumber: "555-0200",
	}
	err = service.CreateContacts(ctx, contact)
	assert.NoError(t, err)

	contact.PhoneNumber = "555-0999"
	ok, err := service.Modify(ctx, contact)
	assert.NoError(t, err)
	assert.True(t, ok)
}
