package repositories

import (
	"WallesterAssigment/dto"
	"WallesterAssigment/interfaces"
	"WallesterAssigment/models"
	"context"
	"gorm.io/gorm"
)

type CustomerRepository struct {
	Db *gorm.DB
}

func (c CustomerRepository) Find(id int) (models.Customer, error) {
	customer := models.Customer{}
	result := c.Db.First(&customer, id)

	if result.Error != nil {
		return models.Customer{}, result.Error
	}

	return customer, nil
}

func (c CustomerRepository) Create(ctx context.Context, customer models.Customer) (models.Customer, error) {
	result := c.Db.WithContext(ctx).Create(&customer)
	return customer, result.Error
}

func (c CustomerRepository) Update(ctx context.Context, id int, customer models.Customer) (models.Customer, error) {
	result := c.Db.WithContext(ctx).Model(&customer).Where("id = ?", id).Updates(&customer)
	return customer, result.Error
}

func (c CustomerRepository) Search(ctx context.Context, search dto.SearchRequest) ([]models.Customer, interfaces.Pagination, error) {
	var customers []models.Customer
	var totalCount int64
	var filteredCount int64

	result := c.Db.
		WithContext(ctx).
		Model(&customers).
		Count(&totalCount).
		Where("firstname like ?", "%"+search.Search+"%").
		Or("lastname like ?", "%"+search.Search+"%").
		Count(&filteredCount).
		Offset(search.Start).
		Limit(search.Length).
		Order(search.GetOrderQueryPart()).
		Find(&customers)

	return customers, interfaces.Pagination{Total: int(totalCount), Filtered: int(filteredCount)}, result.Error
}

func (c CustomerRepository) Delete(customerId int) error {
	var customer models.Customer
	result := c.Db.Where("id = ?", customerId).Delete(&customer)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
