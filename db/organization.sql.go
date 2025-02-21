// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: organization.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const createApproval = `-- name: CreateApproval :exec
insert into approvals (user_id, organization_id)
values ($1, $2)
`

type CreateApprovalParams struct {
	UserID         int32 `json:"user_id"`
	OrganizationID int32 `json:"organization_id"`
}

func (q *Queries) CreateApproval(ctx context.Context, arg CreateApprovalParams) error {
	_, err := q.db.Exec(ctx, createApproval, arg.UserID, arg.OrganizationID)
	return err
}

const createOrganization = `-- name: CreateOrganization :one
INSERT into organizations (name, description, phno, email, address)
values ($1, $2, $3, $4, $5)
returning name
`

type CreateOrganizationParams struct {
	Name        string      `json:"name"`
	Description pgtype.Text `json:"description"`
	Phno        []string    `json:"phno"`
	Email       []string    `json:"email"`
	Address     string      `json:"address"`
}

func (q *Queries) CreateOrganization(ctx context.Context, arg CreateOrganizationParams) (string, error) {
	row := q.db.QueryRow(ctx, createOrganization,
		arg.Name,
		arg.Description,
		arg.Phno,
		arg.Email,
		arg.Address,
	)
	var name string
	err := row.Scan(&name)
	return name, err
}

const getAllOrganizations = `-- name: GetAllOrganizations :many
select id, name, description, phno, email, address
from organizations
`

func (q *Queries) GetAllOrganizations(ctx context.Context) ([]Organization, error) {
	rows, err := q.db.Query(ctx, getAllOrganizations)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Organization
	for rows.Next() {
		var i Organization
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Description,
			&i.Phno,
			&i.Email,
			&i.Address,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getOrgIDByName = `-- name: GetOrgIDByName :one
select id
from organizations
where name = $1
`

func (q *Queries) GetOrgIDByName(ctx context.Context, name string) (int32, error) {
	row := q.db.QueryRow(ctx, getOrgIDByName, name)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const isBDFL = `-- name: IsBDFL :one
select bdfl
from users
where id = $1
`

func (q *Queries) IsBDFL(ctx context.Context, id int32) (pgtype.Bool, error) {
	row := q.db.QueryRow(ctx, isBDFL, id)
	var bdfl pgtype.Bool
	err := row.Scan(&bdfl)
	return bdfl, err
}

const orgExistsByName = `-- name: OrgExistsByName :one
select exists (
    select 1
    from organizations
    where name ilike $1
  )
`

func (q *Queries) OrgExistsByName(ctx context.Context, name string) (bool, error) {
	row := q.db.QueryRow(ctx, orgExistsByName, name)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}
