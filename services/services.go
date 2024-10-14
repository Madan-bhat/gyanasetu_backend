package services

import (
	"context"
	"gyanasetu/backend/db"

	"github.com/jackc/pgx/v5"
)

type Services struct {
	Db     *db.Queries
	DbSQL  *pgx.Conn
	Ctx    context.Context
	Secret []byte
	BDFLId int32
}
