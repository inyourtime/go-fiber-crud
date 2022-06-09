package service

type CustomerResponse struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	City    string `json:"city"`
	ZipCode string `json:"zipcode"`
}

type CustomerService interface {
	GetCustomers() ([]CustomerResponse, error)
	GetCustomer(id int) (*CustomerResponse, error)
	CreateCustomer(name string, city string, zipcode string) (*CustomerResponse, error)
	UpdateCustomer(id int, name string, city string, zipcode string) error
	DeleteCustomer(id int) error
}
