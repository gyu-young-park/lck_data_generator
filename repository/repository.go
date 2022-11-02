package repository

type Repository interface {
	Store(string, interface{}) error
	Get(string) (string, error)
}
