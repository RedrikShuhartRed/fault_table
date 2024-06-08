package controllers

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/RedrikShuhartRed/fault_table/excel"
	"github.com/RedrikShuhartRed/fault_table/models"
	"github.com/RedrikShuhartRed/fault_table/text"
)

func AddFault(db *sql.DB, fault models.Fault) error {

	_, err := db.Exec(`INSERT INTO fault (turbine, date, code, description) VALUES ($1,$2,$3,$4)`, fault.Turbine, fault.Date, fault.Code, fault.Description)
	if err != nil {
		log.Printf("Ошибка вставки: %s", err)
		return err
	}

	return nil
}

func GetAll(db *sql.DB) error {
	var fault []models.Fault

	var rows *sql.Rows
LOOP:
	query, err := text.EnterGetAll()

	switch query {
	case "1":
		rows, err = db.Query("SELECT turbine, to_char(date,'YYYY-MM-DD'), code, description from fault ORDER BY date ASC")
	case "2":
		rows, err = db.Query("SELECT turbine, to_char(date,'YYYY-MM-DD'), code, description from fault ORDER BY date DESC")
	case "3":
		reader := bufio.NewReader(os.Stdin)

		fmt.Println("Введите начало периода в формате гггг.мм.дд")

		begin, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Ошибка ввода начала периода: %s", err)
			return err
		}
		begin = strings.TrimRight(begin, "\r\n")

		fmt.Println("Введите окончание периода периода в формате гггг.мм.дд")
		end, err := reader.ReadString('\n')
		if err != nil {
			log.Printf("Ошибка ввода окончания периода: %s", err)
			return err
		}
		end = strings.TrimRight(end, "\r\n")

		rows, err = db.Query(`SELECT turbine, to_char(date,'YYYY-MM-DD'), code, description from fault
	WHERE date BETWEEN $1 AND $2 ORDER BY date DESC`, begin, end)
	case "4":
		rows, err = db.Query("SELECT turbine, to_char(date,'YYYY-MM-DD'), code, description from fault ORDER BY turbine ASC, date DESC")
	case "5":
		rows, err = db.Query("SELECT turbine, to_char(date,'YYYY-MM-DD'), code, description from fault ORDER BY code, date DESC")
	default:
		fmt.Println("Введите цифру от 1 до 5")
		goto LOOP

	}

	if err != nil {
		log.Printf("Ошибка считывания данных, %s:", err)
		return err
	}
	defer rows.Close()
	for rows.Next() {
		alarm := models.Fault{}

		err := rows.Scan(&alarm.Turbine, &alarm.Date, &alarm.Code, &alarm.Description)
		if err != nil {
			log.Printf("Ошибка считывания данных, %s:", err)
			return err
		}

		fault = append(fault, alarm)
	}
	excel.CreateExcel(fault)

	return nil
}

func GetByTurbine(turbine string, db *sql.DB) error {

	rows, err := db.Query("SELECT turbine, to_char(date,'YYYY-MM-DD'), code, description from fault WHERE turbine=$1 ORDER BY date DESC", turbine)
	if err != nil {
		log.Printf("Ошибка считывания данных, %s:", err)
		return err
	}
	defer rows.Close()

	var fault []models.Fault

	for rows.Next() {
		alarm := models.Fault{}

		err := rows.Scan(&alarm.Turbine, &alarm.Date, &alarm.Code, &alarm.Description)
		if err != nil {
			log.Printf("Ошибка считывания данных, %s:", err)
			return err
		}

		fault = append(fault, alarm)
	}
	excel.CreateExcel(fault)

	return nil
}

func GetByFault(code string, db *sql.DB) error {

	rows, err := db.Query("SELECT turbine, to_char(date,'YYYY-MM-DD'), code, description from fault WHERE code=$1 ORDER BY date DESC", code)
	if err != nil {
		log.Printf("Ошибка считывания данных, %s:", err)
		return err
	}
	defer rows.Close()

	var fault []models.Fault

	for rows.Next() {
		alarm := models.Fault{}

		err := rows.Scan(&alarm.Turbine, &alarm.Date, &alarm.Code, &alarm.Description)
		if err != nil {
			log.Printf("Ошибка считывания данных, %s:", err)
			return err
		}

		fault = append(fault, alarm)
	}
	excel.CreateExcel(fault)

	return nil
}
