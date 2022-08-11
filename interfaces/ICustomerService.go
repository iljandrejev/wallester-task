package interfaces

import (
	"WallesterAssigment/dto"
	"WallesterAssigment/models"
	"context"
)

type ICustomerService interface {
	Create(ctx context.Context, req dto.CreateCustomerRequest) (models.Customer, error)
	Update(ctx context.Context, customerId int, req dto.CreateCustomerRequest) (models.Customer, error)
	GetById(id int) (models.Customer, error)
	Search(ctx context.Context, search dto.SearchRequest) ([]models.Customer, Pagination, error)
	Delete(ctx context.Context, id int, hash string) error
}
