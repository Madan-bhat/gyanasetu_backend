package models

type CreateOrganizationDTO struct {
	Name        string   `json:"name" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Phno        []string `json:"phno" validate:"required"`
	Email       []string `json:"email" validate:"required"`
	Address     string   `json:"address" validate:"required"`
}

type JoinOrganizationDTO struct {
	Id int32 `json:"id" validate:"required"`
}
