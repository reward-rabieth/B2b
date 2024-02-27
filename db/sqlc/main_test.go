package users

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"os"
	"testing"
)

const (
	driverName = "postgres"
	dbSource   = "postgresql://root:root@localhost:5432/B2B"
)

var (
	testRepo Store
)

func TestMain(m *testing.M) {
	connPool, err := pgxpool.New(context.Background(), dbSource)
	if err != nil {
		log.Fatal("cannot connect to the database", err)
	}

	testRepo = NewStore(connPool)
	os.Exit(m.Run())

}
