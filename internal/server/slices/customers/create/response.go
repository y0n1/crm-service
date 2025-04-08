package create

import (
	"github.com/google/uuid"
	"github.com/y0n1/crm-service/internal/models/aggregates"
)

type CreateCustomerResponse struct {
	ID uuid.UUID `json:"id"`
}

func newCreateCustomerResponseFromAggregate(aggregate *aggregates.CustomerAggregate) *CreateCustomerResponse {
	return &CreateCustomerResponse{
		ID: aggregate.Customer.ID,
	}
}
