package coupons_app_sql

import (
	"context"
)

type Frontend interface {
	CreateStudent(ctx context.Context, studentID int, name string) error
	CreateCoupon(ctx context.Context, couponID int, category string, value int) error
	ClaimCoupon(ctx context.Context, couponID int, studentID int) error
}

type FrontendImpl struct {
	StudentService StudentService
	CouponService  CouponService
}

func NewFrontendImpl(ctx context.Context, StudentService StudentService, CouponService CouponService) (Frontend, error) {
	return &FrontendImpl{StudentService: StudentService, CouponService: CouponService}, nil
}

func (u *FrontendImpl) CreateStudent(ctx context.Context, studentID int, name string) error {
	err := u.StudentService.CreateStudent(ctx, studentID, name)
	return err
}

func (u *FrontendImpl) CreateCoupon(ctx context.Context, couponID int, category string, value int) error {
	err := u.CouponService.CreateCoupon(ctx, couponID, category, value)
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
