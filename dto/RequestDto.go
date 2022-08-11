package dto

import (
	"WallesterAssigment/models"
	"errors"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	_ "github.com/monoculum/formam"
	"time"
)

type CreateCustomerRequest struct {
	Firstname string    `formam:"firstname"`
	Lastname  string    `formam:"lastname"`
	Address   string    `formam:"address"`
	Gender    string    `formam:"gender"`
	Email     string    `formam:"email"`
	Birthday  time.Time `formam:"birthday"`
	Hash      string    `formam:"hash"`
}

func (r CreateCustomerRequest) Validate() error {
	return validation.ValidateStruct(&r,
		validation.Field(&r.Firstname, validation.Required, is.Alpha, validation.Length(3, 100)),
		validation.Field(&r.Lastname, validation.Required, is.Alpha, validation.Length(3, 100)),
		validation.Field(&r.Email, validation.Required, is.Email),
		validation.Field(&r.Gender, validation.Required, validation.In("Male", "Female")),
		validation.Field(&r.Birthday, validation.Required, validation.By(isAdult)),
	)
}

func (r CreateCustomerRequest) ToEntity() models.Customer {
	return models.Customer{
		Firstname: r.Firstname,
		Lastname:  r.Lastname,
		Address:   r.Address,
		Gender:    r.Gender,
		Email:     r.Email,
		Birthdate: r.Birthday,
	}
}

func isAdult(value interface{}) error {
	s, _ := value.(time.Time)
	now := time.Now()
	difference := now.Sub(s)

	fullYears := int64(difference.Hours() / 24 / 365)

	if fullYears < 18 || fullYears > 60 {
		return errors.New("age should be at least 18 or less than 61 years")
	}
	return nil
}

type SearchRequest struct {
	Draw           int    `schema:"draw"`
	Search         string `schema:"search[value]"`
	Column         string
	OrderColumn    int    `schema:"order[0][column]"`
	OrderDirection string `schema:"order[0][dir]"`
	Start          int    `schema:"start"`
	Length         int    `schema:"length"`
}

func (s SearchRequest) SetOrderDirection() SearchRequest {
	switch s.OrderDirection {
	default:
	case "asc":
	case "ASC":
		s.OrderDirection = "ASC"
		break
	case "desc":
	case "DESC":
		s.OrderDirection = "DESC"
		break
	}
	return s
}

func (s SearchRequest) SetOrderColumn(columns map[int]string) SearchRequest {
	s.Column = columns[s.OrderColumn]
	return s
}

func (s SearchRequest) GetOrderQueryPart() string {
	if s.OrderDirection == "asc" || s.OrderDirection == "ASC" {
		return s.Column + " ASC"
	}
	return s.Column + " DESC"
}
