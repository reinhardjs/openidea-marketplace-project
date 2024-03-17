package config

import (
	"database/sql"
	"fmt"
	"log"
	"net/url"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"

	"github.com/spf13/viper"
)

func NewDatabase(viper *viper.Viper) *sql.DB {
	//prepare database
	dbHost := viper.GetString("database.host")
	dbPort := viper.GetString("database.port")
	dbUser := viper.GetString("database.username")
	dbPass := viper.GetString("database.password")
	dbName := viper.GetString("database.name")
	connection := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", dbUser, dbPass, dbHost, dbPort, dbName)

	val := url.Values{}
	env := os.Getenv("ENV")

	if env == "" || env == "dev" {
		env = "dev"
		val.Add("sslmode", "disable")
	}

	if env == "prod" {
		val.Add("sslmode", "verify-full")
		val.Add("sslrootcert", "ap-southeast-1-bundle.pem")
	}

	dsn := fmt.Sprintf("%s?%s", connection, val.Encode())
	dbConn, err := sql.Open(`pgx`, dsn)

	if err != nil {
		log.Fatal("failed to open connection to database", err)
	}

	err = dbConn.Ping()
	if err != nil {
		log.Fatal("failed to ping database ", err)
	}

	return dbConn
}
