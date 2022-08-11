package mocks

import (
	"WallesterAssigment/dto"
	"WallesterAssigment/interfaces"
	"WallesterAssigment/models"
	"context"
	"errors"
	"time"
)

type MockCustomerRepository struct{}

func (m MockCustomerRepository) Create(ctx context.Context, customer models.Customer) (models.Customer, error) {
	now := time.Now()
	customer.CreatedAt = now
	customer.UpdatedAt = now
	customer.Hash = models.AsSha256(customer)
	customer.Id = 1
	return customer, nil
}

func (m MockCustomerRepository) Find(id int) (models.Customer, error) {
	for _, customer := range GetCustomers() {
		if customer.Id == id {
			return customer, nil
		}
	}
	return models.Customer{}, errors.New("No customer")
}

func (m MockCustomerRepository) Update(ctx context.Context, id int, customer models.Customer) (models.Customer, error) {
	customer.UpdatedAt = time.Now()
	customer.Hash = models.AsSha256(customer)
	return customer, nil
}

func (m MockCustomerRepository) Search(ctx context.Context, search dto.SearchRequest) ([]models.Customer, interfaces.Pagination, error) {
	return GetCustomers(), interfaces.Pagination{Total: 2, Filtered: 2}, nil
}

func (m MockCustomerRepository) Delete(customerId int) error {
	return nil
}

func GetCustomers() []models.Customer {
	var customers []models.Customer
	customer1 := models.Customer{
		Firstname: "Peeter",
		Lastname:  "Pakkiraam",
		Gender:    "Male",
		Email:     "peeter.pakkiraam@raam.ee",
		Birthdate: time.Date(1990, 05, 04, 00, 0, 0, 0, time.UTC),
		CreatedAt: time.Date(2022, 8, 10, 00, 0, 0, 0, time.UTC),
		UpdatedAt: time.Date(2022, 8, 10, 25, 0, 0, 0, time.UTC),
	}
	customer1.Hash = models.AsSha256(customer1)
	customer1.Id = 1
	customers = append(customers, customer1)
	customer2 := models.Customer{
		Firstname: "Liisa",
		Lastname:  "Pakkiraam",
		Gender:    "Female",
		Email:     "liisa.pakkiraam@raam.ee",
		Birthdate: time.Date(1991, 8, 20, 00, 0, 0, 0, time.UTC),
		CreatedAt: time.Date(2022, 8, 10, 00, 0, 0, 0, time.UTC),
		UpdatedAt: time.Date(2022, 8, 10, 00, 0, 0, 0, time.UTC),
	}
	customer2.Hash = models.AsSha256(customer2)
	customer2.Id = 2
	customers = append(customers, customer2)

	return customers
}
