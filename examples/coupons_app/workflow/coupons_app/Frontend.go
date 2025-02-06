package coupons_app

import (
	"context"
)

type Frontend interface {
	CreateStudent(ctx context.Context, studentID int, name string) (Student, error)
	CreateCoupon(ctx context.Context, couponID int, category string) (Coupon, error)
	ClaimCoupon(ctx context.Context, couponID int, category string, studentID int, value int) (Student, error)
}

type FrontendImpl struct {
	StudentService StudentService
	CouponService  CouponService
}

func NewFrontendImpl(ctx context.Context, StudentService StudentService, CouponService CouponService) (Frontend, error) {
	return &FrontendImpl{StudentService: StudentService, CouponService: CouponService}, nil
}

func (u *FrontendImpl) CreateStudent(ctx context.Context, studentID int, name string) (Student, error) {
	student, err := u.StudentService.CreateStudent(ctx, studentID, name)
	return student, err
}

func (u *FrontendImpl) CreateCoupon(ctx context.Context, couponID int, category string) (Coupon, error) {
	coupon, err := u.CouponService.CreateCoupon(ctx, couponID, category)
	return coupon, err
}

func (u *FrontendImpl) ClaimCoupon(ctx context.Context, couponID int, category string, studentID int, value int) (Student, error) {
	err := u.CouponService.ClaimCoupon(ctx, couponID, category, studentID)
	if err != nil {
		return Student{}, err
	}
	student, err := u.StudentService.AddToBalance(ctx, studentID, value)
	return student, err
}
