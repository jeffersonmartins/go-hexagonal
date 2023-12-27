package db_test

import (
	"database/sql"
	"log"

	"testing"

	"github.com/jeffersonmartins/go-hexagonal/adapters/db"
	"github.com/stretchr/testify/require"
)

var Db *sql.DB

func setUp() {
	Db, _ = sql.Open("sqlite3", ":memory:")
	createTable(Db)
	createProduct(Db)
}

func createTable(db *sql.DB) {
	table := `CREATE TABLE products (
		"id" string,
		"name" string,
		"price" float,
		"status" string
	)`

	_, err := db.Exec(table)

	if err != nil {
		panic(err)
	}

	db.Exec(table)
}

func createProduct(db *sql.DB) {
	insert := `INSERT INTO products (id, name, price, status) VALUES ("1", "Product A", 0.0, "disabled")`
	stmt, err := db.Prepare(insert)

	if err != nil {
		log.Fatal(err.Error())
	}

	stmt.Exec()

}

func TestProductDb_Get(t *testing.T) {
	setUp()
	defer Db.Close()
	db := db.NewProductDb(Db)
	product, err := db.Get("1")
	require.Nil(t, err)
	require.Equal(t, "Product A", product.GetName())
	require.Equal(t, 0.0, product.GetPrice())
	require.Equal(t, "disabled", product.GetStatus())

}
