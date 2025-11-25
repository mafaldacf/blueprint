package eshopmicroservices

import (
	"context"
	"fmt"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
	"go.mongodb.org/mongo-driver/bson"
)

type DiscountService interface {
	GetDiscount(ctx context.Context, request GetDiscountRequest) (Coupon, error)
	CreateDiscount(ctx context.Context, request CreateDiscountRequest) (Coupon, error)
	UpdateDiscount(ctx context.Context, request UpdateDiscountRequest) (Coupon, error)
	DeleteDiscount(ctx context.Context, request DeleteDiscountRequest) (DeleteDiscountResponse, error)
}

type DiscountServiceImpl struct {
	database backend.NoSQLDatabase
}

func NewDiscountServiceImpl(ctx context.Context, database backend.NoSQLDatabase) (DiscountService, error) {
	s := &DiscountServiceImpl{
		database: database,
	}
	return s, nil
}

func (s *DiscountServiceImpl) GetDiscount(ctx context.Context, request GetDiscountRequest) (Coupon, error) {
	return s.get(ctx, request.ProductName)
}

func (s *DiscountServiceImpl) CreateDiscount(ctx context.Context, request CreateDiscountRequest) (Coupon, error) {
	err := s.add(ctx, request.Coupon)
	if err != nil {
		return Coupon{}, err
	}
	return request.Coupon, nil
}

func (s *DiscountServiceImpl) UpdateDiscount(ctx context.Context, request UpdateDiscountRequest) (Coupon, error) {
	err := s.update(ctx, request.Coupon)
	if err != nil {
		return Coupon{}, err
	}
	return request.Coupon, nil
}

func (s *DiscountServiceImpl) DeleteDiscount(ctx context.Context, request DeleteDiscountRequest) (DeleteDiscountResponse, error) {
	err := s.remove(ctx, request.ProductName)
	if err != nil {
		return DeleteDiscountResponse{Success: false}, err
	}
	return DeleteDiscountResponse{Success: true}, nil
}

func (s *DiscountServiceImpl) get(ctx context.Context, productName string) (Coupon, error) {
	collection, err := s.database.GetCollection(ctx, "discount_db", "coupon")
	if err != nil {
		return Coupon{}, err
	}
	filter := bson.D{{Key: "ProductName", Value: productName}}
	cursor, err := collection.FindOne(ctx, filter)
	if err != nil {
		return Coupon{}, err
	}
	var coupon Coupon
	ok, err := cursor.One(ctx, &coupon)
	if err != nil {
		return Coupon{}, err
	}
	if !ok {
		return Coupon{}, fmt.Errorf("coupon not found for product name (%s)", productName)
	}
	return coupon, nil
}

func (s *DiscountServiceImpl) add(ctx context.Context, coupon Coupon) error {
	collection, err := s.database.GetCollection(ctx, "discount_db", "coupon")
	if err != nil {
		return err
	}
	return collection.InsertOne(ctx, coupon)
}

func (s *DiscountServiceImpl) update(ctx context.Context, coupon Coupon) error {
	collection, err := s.database.GetCollection(ctx, "discount_db", "coupon")
	if err != nil {
		return err
	}
	filter := bson.D{{Key: "Id", Value: coupon.Id}}
	updated, err := collection.ReplaceOne(ctx, filter, coupon)
	if err != nil {
		return err
	}
	if updated == 0 {
		return fmt.Errorf("coupon not found for id (%d)", coupon.Id)
	}
	return nil
}

func (s *DiscountServiceImpl) remove(ctx context.Context, productName string) error {
	collection, err := s.database.GetCollection(ctx, "discount_db", "coupon")
	if err != nil {
		return err
	}
	filter := bson.D{{Key: "ProductName", Value: productName}}
	return collection.DeleteOne(ctx, filter)
}
