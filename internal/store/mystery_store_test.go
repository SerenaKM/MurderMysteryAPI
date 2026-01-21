package store

import (
	"database/sql"
	"testing"

	_ "github.com/jackc/pgx/v4/stdlib"
)

// helper function to connect to test database
func setupTestDB(t *testing.T) *sql.DB {
	db, err := sql.Open("pgx", "host=localhost user=postgres password=postgres dbname=postgres port=5433 sslmode=disable")
	if err != nil {
		t.Fatalf("Opening test db: %v", err)
	}

	// run migrations for test db
	err = Migrate(db, "../../migrations/")
	if err != nil {
		t.Fatalf("Migrating test db error: %v", err)
	}

	_, err = db.Exec(`TRUNCATE cases, case_suspects, crime_scenes, suspects CASCADE`) // clear database tables before each test
	if err != nil {
		t.Fatalf("Truncating tables error: %v", err)
	}
	return db
}

func TestCreateWorkout(t *testing.T) {
	db := setupTestDB(t)
	defer db.Close()
}

// convert integer to pointer to integer
func IntPtr(i int) *int {
	return &i
}

// convert float to pointer to float
func FloatPtr(i float64) *float64 {
	return &i
}