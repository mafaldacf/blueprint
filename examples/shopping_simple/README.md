# Shopping Simple

## Compiling

If errors are encountered:

```zsh
go clean -cache -modcache
export GOFLAGS=-mod=mod
export GOWORK=off
go mod tidy
```

Build application
```zsh
go run wiring/main.go -h
rm -rf build
go run wiring/main.go -w docker -o build
```

## Running

```zsh
nano build/docker/.env
```

```zsh
CART_DB_BIND_ADDR=0.0.0.0:9000
CART_SERVICE_THRIFT_BIND_ADDR=0.0.0.0:9001
FRONTEND_HTTP_BIND_ADDR=0.0.0.0:9002
PRODUCT_DB_BIND_ADDR=0.0.0.0:9003
PRODUCT_QUEUE_BIND_ADDR=0.0.0.0:9004
PRODUCT_SERVICE_THRIFT_BIND_ADDR=0.0.0.0:9005
```

```zsh
docker compose -f build/docker/docker-compose.yml up
```

## Testing

```go
type Frontend interface {
	GetAllProducts(ctx context.Context) ([]Product, error)
	GetProduct(ctx context.Context, productID string) (Product, error)
	CreateProduct(ctx context.Context, productID string, description string, pricePerUnit int, category string) (Product, error)
	DeleteProduct(ctx context.Context, productID string) (bool, error)
	GetCart(ctx context.Context, cartID string) (Cart, error)
	AddProductToCart(ctx context.Context, cartID string, productID string) error
}
```

```zsh
# ----- PRODUCTS -----
# GET
curl http://localhost:9002/GetAllProducts
curl http://localhost:9002/GetProduct?productID=zero
# CREATE
curl http://localhost:9002/CreateProduct?productID=zero\&description=product0\&pricePerUnit=5\&category=snacks
curl http://localhost:9002/CreateProduct?productID=one\&description=product1\&pricePerUnit=10\&category=drinks
curl http://localhost:9002/CreateProduct?productID=two\&description=product2\&pricePerUnit=15\&category=cakes
# DELETE
curl http://localhost:9002/DeleteProduct?productID=zero

# ----- PRODUCTS -----
# GET
curl http://localhost:9002/GetCart?cartID=zero
# ADD
curl http://localhost:9002/AddProductToCart?cartID=zero\&productID=zero
```
