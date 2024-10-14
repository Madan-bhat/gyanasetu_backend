package handlers

import (
	"fmt"
	"gyanasetu/backend/db"
	"gyanasetu/backend/models"
	"gyanasetu/backend/services"
	"net/http"
	"strings"

	"github.com/jackc/pgx/v5/pgtype"
)

func (s *Handlers) Register(w http.ResponseWriter, r *http.Request) {
	var params models.RegisterDTO

	if !s.Services.DecodeAndValidateRequest(w, r, &params, s.Validator) {
		return
	}

	exists, err := s.Services.Db.UserExists(s.Services.Ctx, params.Email)
	if s.Services.ISEOnError(w, err) {
		return
	}

	if exists {
		sGID, err := s.Services.Db.GetGIdByEmail(s.Services.Ctx, params.Email)
		if s.Services.ISEOnError(w, err) {
			return
		}
		valid := services.CompareGID(sGID, params.GID)
		if valid {
			token, err := s.Services.CreateToken(params.Email)
			if s.Services.ISEOnError(w, err) {
				return
			}
			s.Services.WriteJson(w, map[string]string{
				"token": token,
			}, http.StatusOK)
			return
		} else {
			s.Services.HttpError(w, "Invalid GID", http.StatusUnauthorized)
			return
		}
	} else {
		hGID := services.HashGID(params.GID)
		name, err := s.Services.Db.CreateUser(s.Services.Ctx, db.CreateUserParams{
			Email: params.Email,
			Name:  params.Name,
			Gid:   hGID,
		})

		if s.Services.ISEOnError(w, err) {
			return
		}
		token, err := s.Services.CreateToken(params.Email)
		if s.Services.ISEOnError(w, err) {
			return
		}
		s.Services.WriteJson(w, map[string]string{
			"message": fmt.Sprintf("Successfully created user %s", name),
			"token":   token,
		}, http.StatusCreated)
		return
	}
}

func (s *Handlers) SelectRole(w http.ResponseWriter, r *http.Request) {
	id, ok := r.Context().Value("user_id").(int32)
	if !ok {
		s.Services.RespondJson(w, "no uder_id", http.StatusBadRequest)
		return
	}
	var selectRoleDTO models.SelectRoleDTO
	if !s.Services.DecodeAndValidateRequest(w, r, &selectRoleDTO, s.Validator) {
		return
	}
	rawRole := strings.ToLower(selectRoleDTO.Role)

	var role string
	switch rawRole {
	case "teacher":
		role = "teacher"
	case "student":
		role = "student"
	default:
		s.Services.HttpError(w, "Invalid role", http.StatusBadRequest)
		return
	}
	err := s.Services.Db.UpdateRole(s.Services.Ctx, db.UpdateRoleParams{
		Role: pgtype.Text{
			String: role,
			Valid:  true,
		},
		ID: id,
	})
	if s.Services.ISEOnError(w, err) {
		return
	}
	// tx, err := s.Services.DbSQL.Begin()
	// if s.Services.ISEOnError(w, err) {
	// 	return
	// }
	// qtx := s.Services.Db.WithTx(tx)
	// err = qtx.UpdateRole(s.Services.Ctx, db.UpdateRoleParams{
	// 	Role: sql.NullString{
	// 		String: ,
	// 	},
	// })
	// if s.Services.ISEOnError(w, err) {
	// 	return
	// }
	// tx.Commit()
	s.Services.RespondJson(w, "Successfully set role", http.StatusOK)
}
