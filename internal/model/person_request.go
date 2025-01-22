package model

import (
	"golang-simple-api/internal/entity"
	"time"
)

type PersonRequest struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	BirthDate time.Time `json:"birth_date"`
}

func ToEntity(request PersonRequest) entity.Person {
	return entity.Person{
		FirstName: request.FirstName,
		LastName:  request.LastName,
		Email:     request.Email,
		BirthDate: request.BirthDate,
	}
}
