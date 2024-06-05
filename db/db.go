package db

import (
	"database/sql"
	"encoding/json"
	"log"
	"os"
	"time"

	"github.com/RedrikShuhartRed/fault_table/models"
	_ "github.com/lib/pq"
)

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
		turbine VARCHAR(4) NOT NULL,
		date VARCHAR(15) NOT NULL,
		code NUMERIC(5) NOT NULL,
		description VARCHAR(50) NOT NULL)`,
	)

	if err != nil {
		log.Printf("Ошибка создания таблицы fault:%s", err)
		return nil
	}

	data, err := os.ReadFile("D:/project/alarmtable/db/DATA.json")
	if err != nil {
		log.Printf("Ошибка чтения файла с данными: %s", err)
		return err
	}

	var fault []models.Fault
	err = json.Unmarshal(data, &fault)
	if err != nil {
		log.Printf("Ошибка Unmarshal исходных данных: %s", err)
		return err
	}
	for _, f := range fault {
		date, _ := time.Parse("02.01.2006", f.Date)
		dateStr := date.Format("02.01.2006")
		_, err := db.Exec(`INSERT INTO fault (turbine, date, code, description) VALUES ($1,$2,$3,$4)`, f.Turbine, dateStr, f.Code, f.Description)
		if err != nil {
			log.Printf("Ошибка вставки данных из файла в БД: %s", err)
			return err
		}
	}
	return nil
}
