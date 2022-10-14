package filter

type Filter interface{
	Filtering(data interface{}) bool
}

