package crawler

type Crawler interface {
	Ready()
	SetData(data interface{})
	GetResult() interface{}
	GoroutineGetResult(done <-chan interface{}) chan interface{}
}
