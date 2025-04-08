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
	if err := validateRequiredFields(customer.FirstName, customer.LastName, customer.Email, customer.Phone); err != nil {
		return nil, err
	}

	var zeroTime time.Time

	return &CustomerAggregate{
		Customer:  customer,
		Contacted: false,
		Metadata:  valueobjects.NewMetadata(time.Now(), zeroTime),
	}, nil
}

func validateRequiredFields(firstName, lastName, email, phone string) error {
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

	if phone == zeroString {
		return fmt.Errorf("cannot create customer: missing phone number")
	}

	return nil
}

func (c *CustomerAggregate) Update(firstName, lastName, role, email, phone string, contacted bool) error {
	if err := validateRequiredFields(firstName, lastName, email, phone); err != nil {
		return err
	}

	c.Customer.FirstName = firstName
	c.Customer.LastName = lastName
	c.Customer.Role = role
	c.Customer.Email = email
	c.Customer.Phone = phone
	c.Contacted = contacted

	c.Metadata.UpdatedAt = time.Now()
	return nil
}
