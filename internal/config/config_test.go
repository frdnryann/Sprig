package config_test

import (
	"context"
	"fmt"
	"go-financial/internal/config"
	"testing"
)

func TestSQLConn(t *testing.T) {
	cfg := config.LoadConfig()

	// fmt.Println(cfg)
	db, err := config.OpenConnection(cfg)
	if err != nil {
		panic(err)
	}

	defer db.Close()
}

func TestSQLQuery(t *testing.T) {
	cfg := config.LoadConfig()

	// fmt.Println(cfg)
	db, _ := config.OpenConnection(cfg)
	defer db.Close()

	ctx := context.Background()

	query := "INSERT INTO users(name, email, password) VALUES('ryanbajindul', 'ryan24@gmail.com', 'ryann24432')"
	_, err := db.ExecContext(ctx, query)
	if err != nil {
		panic(err)
	}

	fmt.Println("Berhasil menambahkan data pengguna!")
}
