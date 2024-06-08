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
)

func AddFault(db *sql.DB) error {
	var fault models.Fault

	reader := bufio.NewReader(os.Stdin)
	var err error

	fmt.Println("Введите номер турбины")
	fault.Turbine, err = reader.ReadString('\n')
	if err != nil {
		log.Printf("Ошибка ввода номера турбины: %s", err)
		return err
	}

	fmt.Println("Введите дату ошибки в формате гггг.мм.дд")
	fault.Date, err = reader.ReadString('\n')
	if err != nil {
		log.Printf("Ошибка ввода даты: %s", err)
		return err
	}

	fmt.Println("Введите код ошибки")
	fault.Code, err = reader.ReadString('\n')
	if err != nil {
		log.Printf("Ошибка ввода кода аварии: %s", err)
		return err
	}

	fmt.Println("Введите описание ошибки")
	fault.Description, err = reader.ReadString('\n')
	if err != nil {
		log.Printf("Ошибка ввода описания аварии: %s", err)
	}

	fault.Turbine = strings.TrimRight(fault.Turbine, "\r\n")
	fault.Date = strings.TrimRight(fault.Date, "\r\n")
	fault.Code = strings.TrimRight(fault.Code, "\r\n")
	fault.Description = strings.TrimRight(fault.Description, "\r\n")

	_, err = db.Exec(`INSERT INTO fault (turbine, date, code, description) VALUES ($1,$2,$3,$4)`, fault.Turbine, fault.Date, fault.Code, fault.Description)
	if err != nil {
		log.Printf("Ошибка вставки: %s", err)
		return err
	}

	return nil
}

func GetAll(db *sql.DB) error {
	var fault []models.Fault
	var query string
	var rows *sql.Rows
	fmt.Println("1 - Выгрузить по возрастанию даты авари")
	fmt.Println("2 - Выгрузить по убыванию даты авари")
	fmt.Println("3 - Выгрузить за определенный период")
	fmt.Println("4 - Сгруппировать выгрузку по номеру турбины")
	fmt.Println("5 - Сгрупиировать выгрузку по коду аварии")
	_, err := fmt.Scanln(&query)
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
		rows, err = db.Query("SELECT turbine, to_char(date,'YYYY-MM-DD'), code, description from fault ORDER BY turbine")
	case "5":
		rows, err = db.Query("SELECT turbine, to_char(date,'YYYY-MM-DD'), code, description from fault ORDER BY code")
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

func GetByTurbine(db *sql.DB) error {

	fmt.Println("Введите номер турбины")
	reader := bufio.NewReader(os.Stdin)
	turbine, err := reader.ReadString('\n')
	if err != nil {
		log.Printf("Ошибка ввода номера турбины: %s", err)
		return err
	}
	turbine = strings.TrimRight(turbine, "\r\n")

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
