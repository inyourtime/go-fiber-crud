package service

import (
	"gobasic/repository"
	"log"
)

type customerService struct {
	custRepo repository.CustomerRepository
}

func NewCustomerService(custRepo repository.CustomerRepository) CustomerService {
	return customerService{custRepo: custRepo}
}

func (s customerService) GetCustomers() ([]CustomerResponse, error) {
	customers, err := s.custRepo.GetAll()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	custResponses := []CustomerResponse{}
	for _, customer := range customers {
		custResponse := CustomerResponse{
			ID:      customer.ID,
			Name:    customer.Name,
			City:    customer.City,
			ZipCode: customer.ZipCode,
		}
		custResponses = append(custResponses, custResponse)
	}

	return custResponses, nil
}

func (s customerService) GetCustomer(id int) (*CustomerResponse, error) {

	customer, err := s.custRepo.GetById(id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	custResponse := CustomerResponse{
		ID:      customer.ID,
		Name:    customer.Name,
		City:    customer.City,
		ZipCode: customer.ZipCode,
	}
	return &custResponse, nil
}

func (s customerService) CreateCustomer(name string, city string, zipcode string) (*CustomerResponse, error) {

	customer, err := s.custRepo.CreateOne(name, city, zipcode)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	custResponse := CustomerResponse{
		ID:      customer.ID,
		Name:    customer.Name,
		City:    customer.City,
		ZipCode: customer.ZipCode,
	}

	return &custResponse, nil
}

func (s customerService) UpdateCustomer(id int, name string, city string, zipcode string) error {

	err := s.custRepo.UpdateById(id, name, city, zipcode)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (s customerService) DeleteCustomer(id int) error {

	err := s.custRepo.DeleteById(id)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
