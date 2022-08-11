package services

import (
	"WallesterAssigment/dto"
	"WallesterAssigment/interfaces"
	"WallesterAssigment/models"
	"context"
	"errors"
	"fmt"
	"time"
)

type CustomerService struct {
	interfaces.ICustomerRepository
}

func (c CustomerService) GetById(id int) (models.Customer, error) {
	return c.ICustomerRepository.Find(id)
}

func (c CustomerService) Create(ctx context.Context, req dto.CreateCustomerRequest) (models.Customer, error) {
	if err := req.Validate(); err != nil {
		return models.Customer{}, err
	}
	now := time.Now()
	customer := req.ToEntity()
	customer.CreatedAt = now
	customer.UpdatedAt = now
	result, err := c.ICustomerRepository.Create(ctx, customer)
	if err != nil {
		return models.Customer{}, err
	}
	return result, nil
}

func (c CustomerService) Update(ctx context.Context, customerId int, req dto.CreateCustomerRequest) (models.Customer, error) {
	if err := req.Validate(); err != nil {
		return models.Customer{}, err
	}
	customer, err := c.ICustomerRepository.Find(customerId)
	if err != nil {
		return models.Customer{}, err
	}
	if customer.Hash != req.Hash {
		return models.Customer{}, errors.New("Conflict. Customer data was changed before update")
	}
	customer.Firstname = req.Firstname
	customer.Lastname = req.Lastname
	customer.Birthdate = req.Birthday
	customer.Gender = req.Gender
	customer.Address = req.Address
	customer.Email = req.Email

	result, err := c.ICustomerRepository.Update(ctx, customerId, customer)
	if err != nil {
		return models.Customer{}, err
	}
	return result, nil
}

func (c CustomerService) Search(ctx context.Context, search dto.SearchRequest) ([]models.Customer, interfaces.Pagination, error) {
	return c.ICustomerRepository.Search(ctx, search)
}

func (c CustomerService) Delete(ctx context.Context, id int, hash string) error {
	customer, err := c.ICustomerRepository.Find(id)
	if err != nil {
		return err
	}
	if customer.Hash != hash {
		return errors.New("Conflict. Customer data was changed before delete")
	}
	err = c.ICustomerRepository.Delete(id)
	if err != nil {
		fmt.Println(err)
		return errors.New("Something went wrong")
	}
	return nil
}
