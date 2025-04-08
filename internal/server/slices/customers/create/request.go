package create

type CreateCustomerRequest struct {
	FirstName string `json:"firstName" validate:"required"`
	LastName  string `json:"lastName" validate:"required"`
	Role      string `json:"role,omitempty"`
	Email     string `json:"email" validate:"required"`
	Phone     string `json:"phone,omitempty"`
}
