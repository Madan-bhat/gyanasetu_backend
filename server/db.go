package server

import (
	"context"
	"fmt"
	"os"
	"sync"

	"gyanasetu/backend/db"
	"gyanasetu/backend/services"

	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/joho/godotenv"
)

var (
	dbOnce = &sync.Once{}
	conn   *pgx.Conn
	dbInst *db.Queries
)

func RegisterDB(ctx context.Context) (*db.Queries, *pgx.Conn, int32) {
	err := godotenv.Load()
	if err != nil {
		panic(fmt.Errorf("[Environment Error]: Cannot load .env file - %v", err))
	}
	// 0	pgUser := os.Getenv("POSTGRES_USER")
	// 	pgPass := os.Getenv("POSTGRES_PASSWORD")
	// 	pgUrl := os.Getenv("POSTGRES_URL")
	// 	pgDB := os.Getenv("POSTGRES_DB")

	// 	dbUrl := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", pgUser, pgPass, pgUrl, pgDB)

	dbOnce.Do(func() {

		var err error
		conn, err = pgx.Connect(ctx, "postgresql://neondb_owner:b3rzQfpKw6EG@ep-little-moon-a5nyw9tp.us-east-2.aws.neon.tech/neondb?sslmode=require")
		if err != nil {
			panic(fmt.Errorf("[Database Error]: Initialization failed - %v", err))
		}

		if err := conn.Ping(ctx); err != nil {
			panic(fmt.Errorf("[Database Error]: Cannot connect to database - %v", err))
		}
	})

	dbInst = db.New(conn)
	bID := GenerateBDFL(dbInst)
	return dbInst, conn, bID
}
func GenerateBDFL(queries *db.Queries) int32 {
	godotenv.Load()
	ctx := context.Background()

	exists, err := queries.BDFLExists(ctx)
	if err != nil {
		panic(fmt.Errorf("Cannot check BDFL existance: %v", err))
	}
	if exists {
		id, err := queries.GetBDFLId(ctx)
		if err != nil {
			panic(fmt.Errorf("Cannot get BDFL id: %v", err))
		}
		return id
	}
	bdflName := os.Getenv("BDFL_NAME")
	bdflEmail := os.Getenv("BDFL_EMAIL")
	bdflSig := os.Getenv("BDFL_SIGNATURE")
	if bdflName == "" || bdflEmail == "" || bdflSig == "" {
		panic(fmt.Errorf("BDFL_NAME or BDFL_EMAIL or BDFL_SIGNATURE is missing in environment variables"))
	}
	id, err := queries.CreateBDFL(ctx, db.CreateBDFLParams{
		Name:  bdflName,
		Email: bdflEmail,
		Gid:   services.HashGID(bdflSig),
	})
	if err != nil {
		panic(fmt.Errorf("Cannot create BDFL: %v", err))
	}
	err = queries.UpdateBDFL(context.Background())
	if err != nil {
		panic(fmt.Errorf("Cannot update BDFL: %v", err))
	}
	return id
}
