package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq" // for loading postgres driver
)

const dbDriver = "postgres"

var dbSource = "postgresql://root:password@localhost:5432/simple_bank?sslmode=disable"
var testQueries *Queries

func TestMain(m *testing.M) {

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	testQueries = New(conn)

	os.Exit(m.Run())
}
