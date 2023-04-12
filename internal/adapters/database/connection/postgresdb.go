package connection

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"strings"

	_ "github.com/lib/pq"
	"go.uber.org/zap"
)

func DatabaseConnection(connection Connection, log *zap.SugaredLogger) (*sql.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		connection.Host,
		connection.Port,
		connection.DbUser,
		connection.DbName,
		connection.Password,
	)

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	err = migration(db)
	if err != nil {
		log.Fatal("Error migrating the database...", err)
	}

	log.Info("Database connection successful...")
	return db, nil
}

func migration(db *sql.DB) error {
	sqlFile, err := ioutil.ReadFile("../../scripts/migration.sql")
	if err != nil {
		panic(err)
	}

	sqlStatements := strings.Split(string(sqlFile), ";")

	for _, stmt := range sqlStatements {
		if strings.TrimSpace(stmt) == "" {
			continue
		}

		sqlStmt, err := db.Prepare(stmt)
		if err != nil {
			panic(err)
		}
		defer sqlStmt.Close()

		_, err = sqlStmt.Exec()
		if err != nil {
			panic(err)
		}
	}

	return nil
}

type Connection struct {
	Host     string
	Port     string
	DbName   string
	DbUser   string
	Password string
}
