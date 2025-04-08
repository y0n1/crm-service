package store

import "errors"

var (
	ErrCustomerNotFound           = errors.New("customer not found")
	ErrCustomerAlreadyExists      = errors.New("customer already exists")
	ErrCustomerEmailAlreadyExists = errors.New("email address is not available")
)
