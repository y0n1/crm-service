package dtos

import (
	"time"

	"github.com/google/uuid"
	"github.com/y0n1/crm-service/internal/models/aggregates"
)

type CustomerDto struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"firstName" validate:"required"`
	LastName  string    `json:"lastName" validate:"required"`
	Role      string    `json:"role,omitempty"`
	Email     string    `json:"email" validate:"required"`
	Phone     string    `json:"phone,omitempty"`
	Contacted bool      `json:"contacted"`
	UpdatedAt time.Time `json:"updatedAt"`
	CreatedAt time.Time `json:"CreatedAt"`
}

func NewCustomerDtoFromAggregate(aggregate *aggregates.CustomerAggregate) *CustomerDto {
	return &CustomerDto{
		ID:        aggregate.Customer.ID,
		Contacted: aggregate.Contacted,
		FirstName: aggregate.Customer.FirstName,
		LastName:  aggregate.Customer.LastName,
		Role:      aggregate.Customer.Role,
		Email:     aggregate.Customer.Email,
		Phone:     aggregate.Customer.Phone,
		UpdatedAt: aggregate.Metadata.UpdatedAt,
		CreatedAt: aggregate.Metadata.CreatedAt,
	}
}