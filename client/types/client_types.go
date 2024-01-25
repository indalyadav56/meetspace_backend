package types

import (
	"meetspace_backend/client/models"

	"github.com/google/uuid"
)

type ClientCreateData struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	CompanyName string `json:"company_name"`
	Email string `json:"email"`
	Password string `json:"password"`
}

type ClientAddUser struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Email string `json:"email"`
	Password string `json:"password"`
	ClientID uuid.UUID `json:"client_id"`
	CreatedBy *models.Client `json:"created_by"`
    UpdatedBy  *models.Client `json:"updated_by"`
}
