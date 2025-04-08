package update

import "github.com/y0n1/crm-service/internal/models/aggregates"

type UpdateCustomerResponse struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Role      string `json:"role,omitempty"`
	Email     string `json:"email" validate:"required"`
	Phone     string `json:"phone,omitempty"`
	Contacted bool   `json:"contacted"`
}

func newUpdateCustomerResponse(aggregate aggregates.CustomerAggregate) *UpdateCustomerResponse {
	return &UpdateCustomerResponse{
		FirstName: aggregate.Customer.FirstName,
		LastName:  aggregate.Customer.LastName,
		Role:      aggregate.Customer.Role,
		Email:     aggregate.Customer.Email,
		Phone:     aggregate.Customer.Phone,
		Contacted: aggregate.Contacted,
	}
}
