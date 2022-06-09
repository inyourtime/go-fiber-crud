package repository

import (
	"gorm.io/gorm"
)

type customerRepositoryDB struct {
	db *gorm.DB
}

func NewCustomerRepositoryDB(db *gorm.DB) CustomerRepository {
	return customerRepositoryDB{db: db}
}

func (r customerRepositoryDB) GetAll() ([]Customer, error) {

	customers := []Customer{}

	tx := r.db.Find(&customers)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return customers, nil
}

func (r customerRepositoryDB) GetById(id int) (*Customer, error) {

	customer := Customer{}

	tx := r.db.First(&customer, id)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &customer, nil
}

func (r customerRepositoryDB) CreateOne(name string, city string, zipcode string) (*Customer, error) {

	customer := Customer{Name: name, City: city, ZipCode: zipcode}

	tx := r.db.Create(&customer)
	if tx.Error != nil {
		return nil, tx.Error
	}

	return &customer, nil
}

func (r customerRepositoryDB) UpdateById(id int, name string, city string, zipcode string) error {

	customer := Customer{Name: name, City: city, ZipCode: zipcode}

	tx := r.db.Model(&Customer{}).Where("id=?", id).Updates(customer)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (r customerRepositoryDB) DeleteById(id int) error {

	tx := r.db.Delete(&Customer{}, id)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
