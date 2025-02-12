package coupons_app_sql

type Coupon struct {
	CouponID int
	Category string
	Value    int
}

type ClaimedCoupon struct {
	CouponID int
	UserID   int
}

type Student struct {
	StudentID int
	Name      string
	Balance   int
}
