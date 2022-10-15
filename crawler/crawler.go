package crawler

type Crawler interface {
	Ready()
	SetData(data interface{})
	GetResult()
}
