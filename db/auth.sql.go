// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0
// source: auth.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const bDFLExists = `-- name: BDFLExists :one
select exists (
    select 1
    from users
    where bdfl = true
  )
`

func (q *Queries) BDFLExists(ctx context.Context) (bool, error) {
	row := q.db.QueryRow(ctx, bDFLExists)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const createBDFL = `-- name: CreateBDFL :one
insert into users(name, email, gid, bdfl)
values ($1, $2, $3, true)
returning id
`

type CreateBDFLParams struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Gid   string `json:"gid"`
}

func (q *Queries) CreateBDFL(ctx context.Context, arg CreateBDFLParams) (int32, error) {
	row := q.db.QueryRow(ctx, createBDFL, arg.Name, arg.Email, arg.Gid)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const createUser = `-- name: CreateUser :one
insert into users (email, name, gid)
values($1, $2, $3)
returning name
`

type CreateUserParams struct {
	Email string `json:"email"`
	Name  string `json:"name"`
	Gid   string `json:"gid"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (string, error) {
	row := q.db.QueryRow(ctx, createUser, arg.Email, arg.Name, arg.Gid)
	var name string
	err := row.Scan(&name)
	return name, err
}

const getBDFLId = `-- name: GetBDFLId :one
select id
from users
where bdfl = true
`

func (q *Queries) GetBDFLId(ctx context.Context) (int32, error) {
	row := q.db.QueryRow(ctx, getBDFLId)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const getGIdByEmail = `-- name: GetGIdByEmail :one
select gid
from users
where email = $1
`

func (q *Queries) GetGIdByEmail(ctx context.Context, email string) (string, error) {
	row := q.db.QueryRow(ctx, getGIdByEmail, email)
	var gid string
	err := row.Scan(&gid)
	return gid, err
}

const getIdByEmail = `-- name: GetIdByEmail :one
select id
from users
where email = $1
`

func (q *Queries) GetIdByEmail(ctx context.Context, email string) (int32, error) {
	row := q.db.QueryRow(ctx, getIdByEmail, email)
	var id int32
	err := row.Scan(&id)
	return id, err
}

const updateBDFL = `-- name: UpdateBDFL :exec
update users
set role = 'bdfl'
where bdfl = true
`

func (q *Queries) UpdateBDFL(ctx context.Context) error {
	_, err := q.db.Exec(ctx, updateBDFL)
	return err
}

const updateRole = `-- name: UpdateRole :exec
update users
set role = $1
where id = $2
`

type UpdateRoleParams struct {
	Role pgtype.Text `json:"role"`
	ID   int32       `json:"id"`
}

func (q *Queries) UpdateRole(ctx context.Context, arg UpdateRoleParams) error {
	_, err := q.db.Exec(ctx, updateRole, arg.Role, arg.ID)
	return err
}

const userExists = `-- name: UserExists :one
select exists(
    select 1
    from users
    where email = $1
  )
`

func (q *Queries) UserExists(ctx context.Context, email string) (bool, error) {
	row := q.db.QueryRow(ctx, userExists, email)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}

const userExistsByID = `-- name: UserExistsByID :one
select exists(
    select 1
    from users
    where id = $1
  )
`

func (q *Queries) UserExistsByID(ctx context.Context, id int32) (bool, error) {
	row := q.db.QueryRow(ctx, userExistsByID, id)
	var exists bool
	err := row.Scan(&exists)
	return exists, err
}
