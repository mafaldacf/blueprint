package tests

import (
	"context"
	"testing"

	"github.com/blueprint-uservices/blueprint/examples/dsb_mediamicroservices/workflow/mediamicroservices"
	"github.com/blueprint-uservices/blueprint/runtime/core/registry"
	"github.com/blueprint-uservices/blueprint/runtime/plugins/simplecache"
	"github.com/blueprint-uservices/blueprint/runtime/plugins/simplenosqldb"
	"github.com/stretchr/testify/assert"
)

var castInfoServiceRegistry = registry.NewServiceRegistry[mediamicroservices.CastInfoService]("castinfo_service")

func init() {
	castInfoServiceRegistry.Register("local", func(ctx context.Context) (mediamicroservices.CastInfoService, error) {
		db, err := simplenosqldb.NewSimpleNoSQLDB(ctx)
		if err != nil {
			return nil, err
		}
		cache, err := simplecache.NewSimpleCache(ctx)
		if err != nil {
			return nil, err
		}
		return mediamicroservices.NewCastInfoServiceImpl(ctx, db, cache)
	})
}

func TestCastInfoServiceWrite(t *testing.T) {
	ctx := context.Background()
	service, err := castInfoServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	castInfo, err := service.WriteCastInfo(ctx, 0, "cast001", "John Actor", "Male", "A renowned actor known for dramatic roles.")
	assert.NoError(t, err)
	assert.Equal(t, "cast001", castInfo.CastInfoID)
	assert.Equal(t, "John Actor", castInfo.Name)
	assert.Equal(t, "Male", castInfo.Gender)
	assert.Equal(t, "A renowned actor known for dramatic roles.", castInfo.Intro)
}

func TestCastInfoServiceWriteMultiple(t *testing.T) {
	ctx := context.Background()
	service, err := castInfoServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	cast1, err := service.WriteCastInfo(ctx, 0, "cast002", "Jane Actress", "Female", "Award-winning actress.")
	assert.NoError(t, err)
	assert.Equal(t, "cast002", cast1.CastInfoID)
	assert.Equal(t, "Jane Actress", cast1.Name)

	cast2, err := service.WriteCastInfo(ctx, 0, "cast003", "Bob Director", "Male", "Acclaimed director and actor.")
	assert.NoError(t, err)
	assert.Equal(t, "cast003", cast2.CastInfoID)
	assert.Equal(t, "Bob Director", cast2.Name)
}

func TestCastInfoServiceReadEmpty(t *testing.T) {
	ctx := context.Background()
	service, err := castInfoServiceRegistry.Get(ctx)
	assert.NoError(t, err)

	castInfos, err := service.ReadCastInfos(ctx, 0, []string{})
	assert.NoError(t, err)
	assert.Nil(t, castInfos)
}
