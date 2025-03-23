package coupons_app_sql

import (
	"context"
)

type Frontend interface {
	CreateStudent(ctx context.Context, name string) error
	CreateCoupon(ctx context.Context, category string, value int) error
	ClaimCoupon(ctx context.Context, couponID int, studentID int) error
}

type FrontendImpl struct {
	StudentService StudentService
	CouponService  CouponService
}

func NewFrontendImpl(ctx context.Context, StudentService StudentService, CouponService CouponService) (Frontend, error) {
	return &FrontendImpl{StudentService: StudentService, CouponService: CouponService}, nil
}

func (u *FrontendImpl) CreateStudent(ctx context.Context, name string) error {
	err := u.StudentService.CreateStudent(ctx, name)
	return err
}

func (u *FrontendImpl) CreateCoupon(ctx context.Context, category string, value int) error {
	err := u.CouponService.CreateCoupon(ctx, category, value)
	return err
}

func (u *FrontendImpl) ClaimCoupon(ctx context.Context, couponID int, studentID int) error {
	value, err := u.CouponService.ClaimCoupon(ctx, couponID, studentID)
	if err != nil {
		return err
	}
	err = u.StudentService.AddToBalance(ctx, studentID, value)
	return err
}
