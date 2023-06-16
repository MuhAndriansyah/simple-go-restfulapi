package config

import (
	"database/sql"
	"go_restfulapi/helpers"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectToDB() *sql.DB {
	db, err := sql.Open("mysql", os.Getenv("DB_CONN_URL"))
	helpers.PanicIfError(err)

	db.SetMaxOpenConns(100)                // Set maksimum koneksi terbuka
	db.SetMaxIdleConns(10)                 // Set maksimum koneksi idle
	db.SetConnMaxIdleTime(5 * time.Minute) // Set durasi maksimum koneksi idle
	db.SetConnMaxLifetime(1 * time.Hour)   // Set durasi maksimum koneksi aktif

	return db
}
