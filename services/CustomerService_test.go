package services

import (
	"WallesterAssigment/dto"
	"WallesterAssigment/interfaces/mocks"
	"WallesterAssigment/models"
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCustomerService_Create_Validate_Request_Failed(t *testing.T) {
	ctx := context.TODO()
	request := dto.CreateCustomerRequest{}
	service := CustomerService{mocks.MockCustomerRepository{}}
	customer, err := service.Create(ctx, request)
	assert.Equal(t, models.Customer{}, customer)
	assert.Equal(t, "Birthday: cannot be blank; Email: cannot be blank; Firstname: cannot be blank; Gender: cannot be blank; Lastname: cannot be blank.", err.Error())
}

func TestCustomerService_Create_Validate_Request_Age_Not_Allowed_Failed(t *testing.T) {
	ctx := context.TODO()
	request := dto.CreateCustomerRequest{
		Firstname: "Peeter",
		Lastname:  "Pakkiraam",
		Gender:    "Male",
		Email:     "peeter.pakkiraam@raam.ee",
		Birthday:  time.Now(),
	}
	service := CustomerService{mocks.MockCustomerRepository{}}
	customer, err := service.Create(ctx, request)
	assert.Equal(t, models.Customer{}, customer)
	assert.Equal(t, "Birthday: age should be at least 18 or less than 61 years.", err.Error())
}

func TestCustomerService_Create_Success(t *testing.T) {
	ctx := context.TODO()
	request := dto.CreateCustomerRequest{
		Firstname: "Peeter",
		Lastname:  "Pakkiraam",
		Gender:    "Male",
		Email:     "peeter.pakkiraam@raam.ee",
		Birthday:  time.Now().Add(-time.Hour * 24 * 365 * 20),
	}
	service := CustomerService{mocks.MockCustomerRepository{}}
	customer, err := service.Create(ctx, request)
	assert.Equal(t, nil, err)
	assert.Equal(t, request.Firstname, customer.Firstname)
	assert.Equal(t, request.Lastname, customer.Lastname)
	assert.Equal(t, request.Gender, customer.Gender)
	assert.Equal(t, request.Birthday, customer.Birthdate)
	assert.Equal(t, request.Address, customer.Address)
	assert.NotEmpty(t, customer.Hash)

}

func TestCustomerService_Update_Validate_Request_Failed(t *testing.T) {
	ctx := context.TODO()
	request := dto.CreateCustomerRequest{}
	service := CustomerService{mocks.MockCustomerRepository{}}
	customer, err := service.Update(ctx, 1, request)
	assert.Equal(t, models.Customer{}, customer)
	assert.Equal(t, "Birthday: cannot be blank; Email: cannot be blank; Firstname: cannot be blank; Gender: cannot be blank; Lastname: cannot be blank.", err.Error())
}

func TestCustomerService_Update_Validate_Request_Age_Not_Allowed_Failed(t *testing.T) {
	ctx := context.TODO()
	request := dto.CreateCustomerRequest{
		Firstname: "Peeter",
		Lastname:  "Pakkiraam",
		Gender:    "Male",
		Email:     "peeter.pakkiraam@raam.ee",
		Birthday:  time.Now(),
	}
	service := CustomerService{mocks.MockCustomerRepository{}}
	customer, err := service.Update(ctx, 1, request)
	assert.Equal(t, models.Customer{}, customer)
	assert.Equal(t, "Birthday: age should be at least 18 or less than 61 years.", err.Error())
}

func TestCustomerService_Update_Hash_Validation_Failed(t *testing.T) {
	ctx := context.TODO()
	request := dto.CreateCustomerRequest{
		Hash:      "Some other hash",
		Firstname: "Peeter",
		Lastname:  "Pakkiraam",
		Gender:    "Male",
		Email:     "peeter.pakkiraam@raam.ee",
		Birthday:  time.Now().Add(-time.Hour * 24 * 365 * 20),
	}
	service := CustomerService{mocks.MockCustomerRepository{}}
	customer, err := service.Update(ctx, 1, request)
	assert.Equal(t, models.Customer{}, customer)
	assert.Equal(t, "Conflict. Customer data was changed before update", err.Error())
}

func TestCustomerService_Update_Success(t *testing.T) {
	var customer models.Customer
	ctx := context.TODO()
	service := CustomerService{mocks.MockCustomerRepository{}}

	for _, cus := range mocks.GetCustomers() {
		if cus.Id == 2 {
			customer = cus
		}
	}
	request := dto.CreateCustomerRequest{
		Hash:      customer.Hash,
		Firstname: customer.Firstname,
		Lastname:  customer.Lastname,
		Gender:    customer.Gender,
		Email:     customer.Email,
		Birthday:  customer.Birthdate,
		Address:   "Tallinn, Viru street",
	}

	updatedCustomer, err := service.Update(ctx, 2, request)
	assert.Equal(t, nil, err)
	assert.Equal(t, request.Address, updatedCustomer.Address)
	assert.NotEqual(t, customer.Hash, updatedCustomer.Hash)
	assert.NotEqual(t, updatedCustomer.CreatedAt, updatedCustomer.UpdatedAt)
	assert.NotEqual(t, customer.UpdatedAt, updatedCustomer.UpdatedAt)
}

func TestCustomerService_Search(t *testing.T) {
	ctx := context.TODO()
	service := CustomerService{mocks.MockCustomerRepository{}}

	customers, pagination, err := service.Search(ctx, dto.SearchRequest{Search: ""})
	assert.Equal(t, nil, err)
	assert.Equal(t, 2, pagination.Filtered)
	assert.Equal(t, 2, pagination.Total)
	assert.Equal(t, 2, len(customers))

}

func TestCustomerService_GetById(t *testing.T) {
	service := CustomerService{mocks.MockCustomerRepository{}}

	customer, err := service.GetById(1)
	assert.Equal(t, nil, err)
	assert.Equal(t, 1, customer.Id)
}
