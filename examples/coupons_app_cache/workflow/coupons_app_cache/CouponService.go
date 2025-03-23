package coupons_app_cache

import (
	"context"
	"os"
	"strings"

	"github.com/blueprint-uservices/blueprint/runtime/core/backend"
)

type CouponService interface {
	CreateCoupon(ctx context.Context, couponID string, category string, value int) error
	ClaimCoupon(ctx context.Context, couponID string, userID string) (int, error)
}

type CouponServiceImpl struct {
	couponsDB backend.RelationalDB
}

func NewCouponServiceImpl(ctx context.Context, couponsDB backend.RelationalDB) (CouponService, error) {
	s := &CouponServiceImpl{couponsDB: couponsDB}
	s.createTables(ctx)
	return s, nil
}

func (s *CouponServiceImpl) CreateCoupon(ctx context.Context, couponID string, category string, value int) error {
	_, err := s.couponsDB.Exec(ctx, "INSERT INTO coupons(category, value) VALUES (?, ?);", category, value)
	return err
}

func (s *CouponServiceImpl) ClaimCoupon(ctx context.Context, couponID string, userID string) (int, error) {
	var coupon Coupon
	err := s.couponsDB.Select(ctx, &coupon, "SELECT * FROM claimed_coupons WHERE coupon_id = ?", couponID)
	if err != nil {
		return -1, err
	}

	_, err = s.couponsDB.Exec(ctx, "INSERT INTO claimed_coupons(coupon_id, user_id) VALUES (?, ?);", couponID, userID)
	return coupon.Value, err
}

func (c *CouponServiceImpl) createTables(ctx context.Context) error {
	sqlBytes, err := os.ReadFile("database/coupons.sql")
	if err != nil {
		return err
	}
	sqlStatements := strings.Split(string(sqlBytes), ";")
	for _, stmt := range sqlStatements {
		_, err := c.couponsDB.Exec(ctx, stmt)
		if err != nil {
			return err
		}
	}
	return nil
}
