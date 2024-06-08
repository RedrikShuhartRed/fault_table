package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var dbs *sql.DB

func ConnectDB() error {
	connStr := "user=postgres password=root host=127.0.0.1 port= 5432 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Printf("Ошибка подключения к БД: %s", err)
		return err
	}

	rows, err := db.Query("SELECT 1 FROM pg_database WHERE datname = 'faultdb'")
	if err != nil {
		log.Printf("Ошибка проверки существования БД: %s", err)
		return err
	}

	var result int
	for rows.Next() {

		err := rows.Scan(&result)
		if err != nil {
			log.Printf("Ошибка чтения результата запроса в pg_database: %s", err)
			return err
		}

	}

	if result != 1 {
		_, err = db.Exec("CREATE DATABASE faultdb")
		if err != nil {
			log.Printf("Ошибка создания БД:%s", err)
			return err
		}

	}
	connStr = "user=postgres password=root host=127.0.0.1 port=5432 dbname=faultdb sslmode=disable"
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Printf("Ошибка подключения к БД faultDB: %s", err)
		return err
	}

	_, err = db.Exec(
		`CREATE TABLE IF NOT EXISTS fault (
		id BIGSERIAL NOT NULL PRIMARY KEY,
		turbine VARCHAR(30) NOT NULL,
		date DATE NOT NULL,
		code VARCHAR(10) NOT NULL,
		description VARCHAR(300) NOT NULL)`,
	)

	if err != nil {
		log.Printf("Ошибка создания таблицы fault:%s", err)
		return nil
	}

	dbs = db
	return nil
}

func GetDB() *sql.DB {
	return dbs
}

func CloseDB(db *sql.DB) error {
	err := db.Close()
	if err != nil {
		log.Printf("Ошибка закрытия подключения в БД: %s", err)
		return err
	}
	return nil
}
