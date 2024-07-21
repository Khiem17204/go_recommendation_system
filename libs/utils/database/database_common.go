package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

type DatabaseManager struct {
	db       *sql.DB
	username string
	password string
	host     string
	dbname   string
}

func NewDatabaseManager() (*DatabaseManager, error) {
	dm := &DatabaseManager{
		username: os.Getenv("DB_USER"),
		password: os.Getenv("DB_PASSWORD"),
		host:     os.Getenv("DB_HOST"),
	}

	connStr := fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=disable",
		dm.username, dm.password, dm.host, dm.dbname)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	dm.db = db
	return dm, nil
}

func (dm *DatabaseManager) Close() {
	dm.db.Close()
}

func (dm *DatabaseManager) ExecuteSQL(query string, args ...interface{}) (sql.Result, error) {
	return dm.db.Exec(query, args...)
}

func (dm *DatabaseManager) CreateDatabase(dbname string) error {
	_, err := dm.ExecuteSQL(fmt.Sprintf("CREATE DATABASE %s", dbname))
	return err
}
