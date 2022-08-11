package main

import (
	"WallesterAssigment/controllers"
	"WallesterAssigment/models"
	"WallesterAssigment/repositories"
	"WallesterAssigment/services"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"sync"
)

type IServiceContainer interface {
	InjectCustomerController() controllers.CustomerController
}

type kernel struct{}

func (k *kernel) InjectCustomerController() controllers.CustomerController {
	con := NewDBConn()
	repository := repositories.CustomerRepository{}
	repository.Db = con
	service := services.CustomerService{ICustomerRepository: repository}
	return controllers.CustomerController{ICustomerService: service}
}

func NewDBConn() (db *gorm.DB) {
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		"disable")

	db, err := gorm.Open(
		postgres.New(postgres.Config{
			DSN:                  dsn,
			PreferSimpleProtocol: true,
		}),
		&gorm.Config{})

	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(&models.Customer{})
	if err != nil {
		panic(err)
	}
	return db
}

var (
	k             *kernel
	containerOnce sync.Once
)

func ServiceContainer() IServiceContainer {
	if k == nil {
		containerOnce.Do(func() {
			k = &kernel{}
		})
	}
	return k
}
