package crawler

type Crawler interface {
	Ready()
	SetQueryOption(queryOption *QueryOption)
	GetResult() interface{}
}
