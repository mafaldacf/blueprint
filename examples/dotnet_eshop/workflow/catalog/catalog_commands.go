package catalog

type CreateItemRequest struct {
	ID                int
	Name              string
	Description       string
	Price             float64
	PriceFileName     string
	CatalogTypeID     int
	CatalogType       CatalogType
	CatalogBrandID    int
	CatalogBrand      CatalogBrand
	AvailableStock    int
	RestockThreshold  int
	MaxStockThreshold int
}

type CreateItemResponse struct {
	Item CatalogItem
}

type UpdateItemResponse struct {
	ID int
}

type DeleteItemRequest struct {
	ID string
}

type GetItemByIDRequest struct {
	ID int
}

type GetItemsByIDsRequest struct {
	IDs []int
}

type GetItemByIDResponse struct {
	Item CatalogItem
}

type GetItemsByIDsResponse struct {
	Item []CatalogItem
}


type GetItemByNameRequest struct {
	Name string
}

type GetItemByNameResponse struct {
	Item CatalogItem
}

type GetAllItemsResponse struct {
	Items []CatalogItem
}
