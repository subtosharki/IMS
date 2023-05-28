package structs

type LoginRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type UpdateNotesRequest struct {
	Notes string `json:"notes" validate:"required"`
}

type CreateUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Admin    bool   `json:"admin" validate:"required"`
}

type APIKeyRequest struct {
	APIKey string `reqHeader:"apikey"`
}

type SetStatusRequest struct {
	Status string `json:"status" validate:"required"`
}

type CreateCloneRequest struct {
	Name     string  `json:"name" validate:"required"`
	Quantity float64 `json:"quantity" validate:"required"`
	Date     string  `json:"date" validate:"required"`
}

type EditNameRequest struct {
	Name string `json:"name" validate:"required"`
}

type CreateCustomerRequest struct {
	Name           string `json:"name" validate:"required"`
	Email          string `json:"email" validate:"required,email"`
	SubscriptionId string `json:"subscriptionId" validate:"required"`
	Phone          string `json:"phone" validate:"required"`
	Address        string `json:"address" validate:"required"`
	City           string `json:"city" validate:"required"`
	State          string `json:"state" validate:"required"`
	Zip            string `json:"zip" validate:"required"`
}
