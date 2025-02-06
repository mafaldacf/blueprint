package coupons_app

type Coupon struct {
	CouponID int
	Category string
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
