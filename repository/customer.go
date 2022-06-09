package repository

type Customer struct {
	ID      int
	Name    string
	City    string
	ZipCode string
}

type CustomerRepository interface {
	GetAll() ([]Customer, error)
	GetById(id int) (*Customer, error)
	CreateOne(name string, city string, zipcode string) (*Customer, error)
	UpdateById(id int, name string, city string, zipcode string) error
	DeleteById(id int) error
}
