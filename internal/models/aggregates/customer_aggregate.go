package aggregates

import (
	"fmt"
	"time"

	"github.com/y0n1/crm-service/internal/models/entities"
	"github.com/y0n1/crm-service/internal/models/valueobjects"
)

type CustomerAggregate struct {
	Customer  *entities.Customer
	Contacted bool
	Metadata  *valueobjects.Metadata
}

func NewCustomerAggregate(customer *entities.Customer) (*CustomerAggregate, error) {
	if err := validateRequiredFields(customer.FirstName, customer.LastName, customer.Email); err != nil {
		return nil, err
	}

	var zeroTime time.Time

	return &CustomerAggregate{
		Customer:  customer,
		Contacted: false,
		Metadata:  valueobjects.NewMetadata(time.Now(), zeroTime),
	}, nil
}

func validateRequiredFields(firstName, lastName, email string) error {
	var zeroString string

	if firstName == zeroString {
		return fmt.Errorf("cannot create customer: missing first name")
	}

	if lastName == zeroString {
		return fmt.Errorf("cannot create customer: missing last name")
	}

	if email == zeroString {
		return fmt.Errorf("cannot create customer: missing email address")
	}

	return nil
}

func (c *CustomerAggregate) Update(firstName, lastName, role, email, phone string, contacted bool) error {
	var zeroString string

	if firstName != zeroString {
		c.Customer.FirstName = firstName
	}

	if lastName != zeroString {
		c.Customer.LastName = lastName
	}
	if role != zeroString {
		c.Customer.Role = role
	}
	if email != zeroString {
		c.Customer.Email = email
	}
	if phone != zeroString {
		c.Customer.Phone = phone
	}

	c.Contacted = contacted
	c.Metadata.UpdatedAt = time.Now()
	return nil
}
