package coupons_app

import (
	"context"
	"fmt"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
	"go.mongodb.org/mongo-driver/bson"
)

type CouponService interface {
	CreateCoupon(ctx context.Context, couponID int, category string) (Coupon, error)
	ClaimCoupon(ctx context.Context, couponID int, category string, userID int) error
}

type CouponServiceImpl struct {
	couponsDB backend.NoSQLDatabase
}

func NewCouponServiceImpl(ctx context.Context, couponsDB backend.NoSQLDatabase) (CouponService, error) {
	s := &CouponServiceImpl{couponsDB: couponsDB}
	return s, nil
}

func (s *CouponServiceImpl) CreateCoupon(ctx context.Context, couponID int, category string) (Coupon, error) {
	coupon := Coupon{
		CouponID: couponID,
		Category: category,
	}
	collection, err := s.couponsDB.GetCollection(ctx, "coupons", "coupons")
	if err != nil {
		return coupon, err
	}
	err = collection.InsertOne(ctx, coupon)
	return coupon, err
}

func (s *CouponServiceImpl) ClaimCoupon(ctx context.Context, couponID int, category string, userID int) error {
	var claimedCoupon ClaimedCoupon

	collection, err := s.couponsDB.GetCollection(ctx, "coupons", "claimed_coupons")
	if err != nil {
		return err
	}

	query := bson.D{{Key: "category", Value: category}, {Key: "userID", Value: userID}}
	result, err := collection.FindOne(ctx, query)
	if err != nil {
		return err
	}
	found, err := result.One(ctx, claimedCoupon)
	if err != nil {
		return err
	}
	if found {
		return fmt.Errorf("user %d already claimed one coupon (%d) within the '%s' category", claimedCoupon.UserID, claimedCoupon.CouponID, category)
	}

	claimedCoupon = ClaimedCoupon{
		CouponID: couponID,
		UserID:   userID,
	}
	err = collection.InsertOne(ctx, claimedCoupon)
	return err
}
