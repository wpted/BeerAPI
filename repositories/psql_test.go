package repositories

import (
	"database/sql"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"log"
	"testing"
)

func setupDB() (*sql.DB, sqlmock.Sqlmock) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("can't create sqlmock: %s", err)
	}
	return db, mock
}
func TestPostgreSQL_BeerExist(t *testing.T) {
	db, mock := setupDB()
	mockPSQL := PostgreSQL{DB: db}
	defer mockPSQL.Close()

	t.Run("Get an existing beer", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "name", "price", "company"}).AddRow(1, "testName", 20, "test_company")

		// mock.ExpectQuery takes doesn't perform simple string comparison
		// instead it compares it as RegexExp, so we have to escape the characters
		mock.ExpectQuery(`SELECT \* FROM beers WHERE name\=\$1`).WillReturnRows(rows)
		fmt.Println(rows)
		result := mockPSQL.BeerExist("testName")
		expected := true
		if result != expected {
			t.Errorf("expected %v, got %v", expected, result)
		}
	})

	t.Run("Get a non-existing beer", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"id", "name", "price", "company"}).AddRow(1, "testName", 20, "test_company")

		mock.ExpectQuery(`SELECT \* FROM beers WHERE name\=\$1`).WillReturnRows(rows)
		fmt.Println(rows)

		result := mockPSQL.BeerExist("nonExistingBeer")
		expected := false
		if result != expected {
			t.Errorf("expected %v, got %v", expected, result)
		}
	})
}
