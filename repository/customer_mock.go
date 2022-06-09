package repository

type customerRepositoryMock struct {
	customers []Customer
}

func NewCustomerRepositoryMock() CustomerRepository {
	customers := []Customer{
		{ID: 1000, Name: "Boat", City: "Hatyai", ZipCode: "90000"},
		{ID: 1001, Name: "Big", City: "Songkla", ZipCode: "90222"},
		{ID: 1002, Name: "Beam", City: "Sabayoi", ZipCode: "90210"},
	}
	return customerRepositoryMock{customers: customers}
}

func (r customerRepositoryMock) GetAll() ([]Customer, error) {
	return nil, nil
}

func (r customerRepositoryMock) GetById(id int) (*Customer, error) {
	return nil, nil
}

func (r customerRepositoryMock) CreateOne(name string, city string, zipcode string) (*Customer, error) {
	return nil, nil
}

func (r customerRepositoryMock) UpdateById(id int, name string, city string, zipcode string) error {
	return nil
}

func (r customerRepositoryMock) DeleteById(id int) error {
	return nil
}
