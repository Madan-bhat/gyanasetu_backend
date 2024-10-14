package services

import (
	"encoding/json"
	"fmt"
	"gyanasetu/backend/models"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	v "github.com/go-playground/validator/v10"
)

func (s *Services) HttpError(w http.ResponseWriter, msg any, status int) {
	errMessage, err := json.Marshal(models.BasicHttpResponse{
		Message: msg,
	})
	if err != nil {
		errMessage, _ := json.Marshal(models.BasicHttpResponse{
			Message: fmt.Sprintf("Internal Server Error: %v", err),
		})
		http.Error(w, string(errMessage), http.StatusInternalServerError)
		return
	}
	http.Error(w, string(errMessage), status)
}

func (s *Services) ISEOnError(w http.ResponseWriter, err error) bool {

	if err != nil {
		s.HttpError(w, fmt.Sprintf("Internal Server Error: %v", err), http.StatusInternalServerError)
		return true
	}
	return false
}
func (s *Services) WriteJson(w http.ResponseWriter, body any, status int) {
	w.WriteHeader(status)
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(body)
}

func (s *Services) RespondJson(w http.ResponseWriter, body any, status int) {
	s.WriteJson(w, models.BasicHttpResponse{
		Message: body,
	}, status)
}

func (s *Services) DecodeAndValidateRequest(w http.ResponseWriter, r *http.Request, decodeStruct any, validator *validator.Validate) bool {
	var err error

	if err = json.NewDecoder(r.Body).Decode(decodeStruct); err != nil {
		s.HttpError(w, fmt.Sprintf("Invalid request body: %v", err), http.StatusBadRequest)
		return false
	}

	if err = validator.Struct(decodeStruct); err != nil {

		validationErrors := err.(v.ValidationErrors)
		var validationErrMsgs []string
		for _, ve := range validationErrors {
			validationErrMsgs = append(validationErrMsgs, fmt.Sprintf("Field %s: %s", ve.Field(), ve.ActualTag()))
		}
		s.HttpError(w, fmt.Sprintf("Validation error: %s", strings.Join(validationErrMsgs, ", ")), http.StatusBadRequest)
		return false
	}

	return true
}
