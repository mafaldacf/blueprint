package digota

/* import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type Frontend interface {
	DeleteProduct(ctx context.Context, id string) error
}

type FrontendImpl struct {
	skuService     SkuService
	orderService   OrderService
	productService ProductService
}

func NewFrontendImpl(ctx context.Context, skuService SkuService, orderService OrderService, productService ProductService) (Frontend, error) {
	s := &FrontendImpl{skuService: skuService, orderService: orderService, productService: productService}
	return s, nil
}

func (s *FrontendImpl) DeleteProduct(ctx context.Context, id string) error {
	err := s.productService.Delete(ctx, id)
	if err != nil {
		return err
	}
	err = s.orderService.Delete(ctx, id)
	return err
} */
