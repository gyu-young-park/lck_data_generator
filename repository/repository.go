package repository

type Repository interface{
	Store(interface{}) error
}
