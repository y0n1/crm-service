package entities

import "github.com/google/uuid"

type Customer struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Role      string    `json:"role,omitempty"`
	Email     string    `json:"email,omitempty"`
	Phone     string    `json:"phone,omitempty"`
}

func NewCustomer(firstName, lastName, role, email, phone string) *Customer {

	return &Customer{
		ID:        uuid.Nil,
		FirstName: firstName,
		LastName:  lastName,
		Role:      role,
		Email:     email,
		Phone:     phone,
	}
}
