package handlers

import (
	"context"
	"fmt"
	"gyanasetu/backend/db"
	"gyanasetu/backend/models"
	"net/http"

	"github.com/jackc/pgx/v5/pgtype"
)

func (s *Handlers) ListOrganizations(w http.ResponseWriter, r *http.Request) {
	orgs, err := s.Services.Db.GetAllOrganizations(s.Services.Ctx)
	if s.Services.ISEOnError(w, err) {
		return
	}
	if len(orgs) == 0 {
		s.Services.RespondJson(w, "No organizations registered", http.StatusNoContent)
		return
	}
	s.Services.WriteJson(w, map[string]interface{}{
		"data": orgs,
	}, http.StatusOK)
}
func (s *Handlers) CreateOrganization(w http.ResponseWriter, r *http.Request) {
	var params models.CreateOrganizationDTO
	if !s.Services.DecodeAndValidateRequest(w, r, &params, s.Validator) {
		return
	}
	var hasDescription bool = false
	if len(params.Description) > 0 {
		hasDescription = true
	}
	exists, err := s.Services.Db.OrgExistsByName(context.Background(), params.Name)
	if s.Services.ISEOnError(w, err) {
		return
	}
	if exists {
		s.Services.HttpError(w, "Organization already exists", http.StatusConflict)
		return
	}
	name, err := s.Services.Db.CreateOrganization(s.Services.Ctx, db.CreateOrganizationParams{
		Name: params.Name,
		Description: pgtype.Text{
			String: params.Description,
			Valid:  hasDescription,
		},
		Phno:    params.Phno,
		Email:   params.Email,
		Address: params.Address,
	})
	if s.Services.ISEOnError(w, err) {
		return
	}
	s.Services.RespondJson(w, fmt.Sprintf("Successfully created organization: %s", name), http.StatusCreated)
}

func (s *Handlers) JoinOrganization(w http.ResponseWriter, r *http.Request) {
	var params models.JoinOrganizationDTO
	userID := r.Context().Value("user_id").(int32)
	if !s.Services.DecodeAndValidateRequest(w, r, &params, s.Validator) {
		return
	}
	err := s.Services.Db.CreateApproval(s.Services.Ctx, db.CreateApprovalParams{
		UserID:         userID,
		OrganizationID: params.Id,
	})
	if err != nil {
		s.Services.HttpError(w, "Cannot create approval", http.StatusInternalServerError)
		return
	} else {
		s.Services.RespondJson(w, "Successfully sent approval", http.StatusOK)
		return
	}
}
