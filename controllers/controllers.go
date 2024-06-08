package controllers

import (
	"database/sql"
	"log"

	"github.com/RedrikShuhartRed/fault_table/excel"
	"github.com/RedrikShuhartRed/fault_table/models"
	"github.com/RedrikShuhartRed/fault_table/query"
	"github.com/RedrikShuhartRed/fault_table/text"
)

func AddFault(db *sql.DB, fault models.Fault) error {

	_, err := db.Exec(query.QueryAddFault, fault.Turbine, fault.Date, fault.Code, fault.Description)
	if err != nil {
		log.Printf("Ошибка вставки: %s", err)
		return err
	}

	return nil
}

func GetAll(enter string, db *sql.DB) error {
	var fault []models.Fault
	var err error
	var rows *sql.Rows

	switch enter {
	case "1":
		rows, err = db.Query(query.QueryGetAllByDateASC)
	case "2":
		rows, err = db.Query(query.QueryGetAllByDateDESC)
	case "3":
		begin, end, _ := text.GetBetweenDate()
		rows, err = db.Query(query.QueryGetAllBetweenDate, begin, end)
	case "4":
		rows, err = db.Query(query.QueryGetAllByTurbine)
	case "5":
		rows, err = db.Query(query.QueryGetAllByCode)
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

	rows, err := db.Query(query.QueryGetByTurbine, turbine)
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

	rows, err := db.Query(query.QueryGetByFault, code)
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
