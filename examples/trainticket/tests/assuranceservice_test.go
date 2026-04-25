package tests

import (
	"context"
	"testing"

	"github.com/blueprint-uservices/blueprint/examples/trainticket/workflow/trainticket"
	"github.com/blueprint-uservices/blueprint/runtime/core/registry"
	"github.com/blueprint-uservices/blueprint/runtime/plugins/simplenosqldb"
	"github.com/stretchr/testify/assert"
)

var assuranceServiceRegistry = registry.NewServiceRegistry[trainticket.AssuranceService]("assurance_service")

func init() {
	assuranceServiceRegistry.Register("local", func(ctx context.Context) (trainticket.AssuranceService, error) {
		db, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		return trainticket.NewAssuranceServiceImpl(ctx, db)
	})
}

func TestAssuranceServiceCreate(t *testing.T) {
	ctx := context.Background()
	service, err := assuranceServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	assurance, err := service.Create(ctx, 1, "order001")
	assert.NoError(t, err)
	assert.Equal(t, "order001", assurance.OrderID)
	assert.NotEmpty(t, assurance.ID)
}

func TestAssuranceServiceFindById(t *testing.T) {
	ctx := context.Background()
	service, err := assuranceServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	assurance, err := service.Create(ctx, 1, "order002")
	assert.NoError(t, err)

	found, err := service.FindAssuranceById(ctx, assurance.ID)
	assert.NoError(t, err)
	assert.Equal(t, "order002", found.OrderID)
}

func TestAssuranceServiceFindByOrderId(t *testing.T) {
	ctx := context.Background()
	service, err := assuranceServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	_, err = service.Create(ctx, 1, "order003")
	assert.NoError(t, err)

	found, err := service.FindAssuranceByOrderId(ctx, "order003")
	assert.NoError(t, err)
	assert.Equal(t, "order003", found.OrderID)
}

func TestAssuranceServiceGetAllAssurances(t *testing.T) {
	ctx := context.Background()
	service, err := assuranceServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	_, err = service.Create(ctx, 1, "order004")
	assert.NoError(t, err)
	_, err = service.Create(ctx, 1, "order005")
	assert.NoError(t, err)

	all, err := service.GetAllAssurances(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, all)
}

func TestAssuranceServiceGetAllAssuranceTypes(t *testing.T) {
	ctx := context.Background()
	service, err := assuranceServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	types, err := service.GetAllAssuranceTypes(ctx)
	assert.NoError(t, err)
	assert.NotEmpty(t, types)
}
