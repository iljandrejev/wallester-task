package models

import (
	"crypto/sha256"
	"fmt"
	"gorm.io/gorm"
	"time"
)

type Customer struct {
	Id        int       `gorm:"column:id"`
	Hash      string    `gorm:"column:hash"`
	Firstname string    `gorm:"column:firstname"`
	Lastname  string    `gorm:"column:lastname"`
	Birthdate time.Time `gorm:"column:birthday"`
	Gender    string    `gorm:"column:gender"`
	Email     string    `gorm:"column:email"`
	Address   string    `gorm:"column:address"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

func (Customer) TableName() string {
	return "customers"
}

func (c *Customer) BeforeUpdate(tx *gorm.DB) (err error) {
	c.UpdatedAt = time.Now()
	c.Hash = AsSha256(c)
	return
}

func (c *Customer) BeforeCreate(tx *gorm.DB) (err error) {
	c.UpdatedAt = time.Now()
	c.CreatedAt = time.Now()
	c.Hash = AsSha256(c)
	return
}

func AsSha256(o interface{}) string {
	h := sha256.New()
	h.Write([]byte(fmt.Sprintf("%v", o)))

	return fmt.Sprintf("%x", h.Sum(nil))
}
