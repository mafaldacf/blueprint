package catalog

type ProductPriceChangedEvent struct {
	CatalogItemID int
	NewPrice float64
	OldPrice float64
}
