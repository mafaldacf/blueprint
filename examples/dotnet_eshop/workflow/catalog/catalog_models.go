package catalog

type CatalogBrand struct {
	ID    int
	Brand string
}

type CatalogItem struct {
	ID                int
	Name              string
	Description       string
	Price             float64
	PictureFileName   string
	CatalogTypeID     int
	CatalogType       CatalogType
	CatalogBrandID    int
	CatalogBrand      CatalogBrand
	AvailableStock    int
	RestockThreshold  int
	MaxStockThreshold int
}

type CatalogType struct {
	ID   int
	Type string
}

type PaginatedItems struct {
	PageIndex int
	PageSize  int
	Count     int64
}
