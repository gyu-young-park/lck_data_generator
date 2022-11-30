package crawler

type QueryOption struct {
	FandomQueryOption
	InvenQueryOption
	Url  string
	Date string
}

type FandomQueryOption struct {
	Season string
}

type InvenQueryOption struct {
	ShipGroup string
}
