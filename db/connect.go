package db

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/manoj2994/go_healthcheck/utils"
)

func createdb(db_host, db_port, db_user, db_pwd, db_name string) error {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=postgres sslmode=disable", db_host, db_port, db_user, db_pwd)
	dbc, err := sql.Open("postgres", connStr)
	utils.CheckError(err, "Error in opening connection;")

	defer dbc.Close()
	_, err = dbc.Exec("CREATE DATABASE " + db_name)
	fmt.Printf("Creating Db %s", db_name)
	utils.CheckError(err, "Error in creating DB;")

	//fmt.Println("DB Created succesfully")
	return nil
}

func Connect() (*sql.DB, error) {
	err := godotenv.Load()
	utils.CheckError(err, "Error in loading the db connection variable")

	db_host := os.Getenv("DB_HOST")
	db_port := os.Getenv("DB_PORT")
	db_user := os.Getenv("DB_USER")
	db_pwd := os.Getenv("DB_PASSWORD")
	db_name := os.Getenv("DB_NAME")
	createdb(db_host, db_port, db_user, db_pwd, db_name)
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", db_host, db_port, db_user, db_pwd, db_name)

	db, err := sql.Open("postgres", connStr)
	utils.CheckError(err, "Error in establishing connection;")

	return db, nil

}
