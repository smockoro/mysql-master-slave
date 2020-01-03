package repository_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/jmoiron/sqlx"
)

func IntegrationTestWithMasterSlave(t *testing.T) {
	shutdown := setupData(t)
	defer shutdown()

	teardownData(t)
}

func setupData(t *testing.T) func() {
	t.Helper()
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s",
		"root",
		os.Getenv("MYSQL_ROOT_PASSWORD"),
		os.Getenv("SETUP_HOST"),
		os.Getenv("DB_SCHEMA"))
	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		t.Fatal(err)
	}
	// drop table

	// create table

	// insert data

	return func() { defer db.Close() }
}

func teardownData(t *testing.T) {
	t.Helper()
	// delete data

	// drop table
}
