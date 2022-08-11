package interfaces

import (
	"WallesterAssigment/dto"
	"WallesterAssigment/models"
	"context"
)

type ICustomerRepository interface {
	Create(ctx context.Context, customer models.Customer) (models.Customer, error)
	Find(id int) (models.Customer, error)
	Update(ctx context.Context, id int, customer models.Customer) (models.Customer, error)
	Search(ctx context.Context, search dto.SearchRequest) ([]models.Customer, Pagination, error)
	Delete(id int) error
}

type Pagination struct {
	Total    int
	Filtered int
}
