package discount

type GetDiscountRequest struct {
	ProductName string
}

type CreateDiscountRequest struct {
	Coupon Coupon
}

type UpdateDiscountRequest struct {
	Coupon Coupon
}

type DeleteDiscountRequest struct {
	ProductName string
}

type DeleteDiscountResponse struct {
	Success bool
}
