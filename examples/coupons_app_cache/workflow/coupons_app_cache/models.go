package coupons_app_cache

type Coupon struct {
	CouponID string
	Category string
	Value    int
}

type ClaimedCoupon struct {
	CouponID string
	UserID   string
}

type Student struct {
	StudentID string
	Name      string
	Balance   int
}
