package handlers

import (
	"gyanasetu/backend/services"

	"github.com/go-playground/validator/v10"
)

type Handlers struct {
	Services  services.Services
	Validator *validator.Validate
}
