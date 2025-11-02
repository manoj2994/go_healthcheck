package db

import (
	"database/sql"

	_ "github.com/lib/pq"
	"github.com/manoj2994/go_healthcheck/utils"
)

type DBResponse struct {
	ID        int64  `json:"id"`
	CheckedAt string `json:"checked_at"`
}

func createTable(db *sql.DB) {
	query := `
    CREATE TABLE IF NOT EXISTS health_checks (
        id SERIAL PRIMARY KEY,
        checked_at TIMESTAMP DEFAULT now()
    );`
	_, err := db.Exec(query)
	utils.CheckError(err, "Table ensured successfully;", "Error in creating table;")

}

func InsertData(db *sql.DB) (error, *DBResponse) {
	createTable(db)
	query := "INSERT INTO health_checks DEFAULT VALUES RETURNING id, checked_at"
	res := DBResponse{}
	err := db.QueryRow(query).Scan(&res.ID, &res.CheckedAt)
	utils.CheckError(err, "Data inserted successfully;", "Error in inserting data;")
	return nil, &res
}
