package controllers

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/RedrikShuhartRed/fault_table/models"
	"github.com/xuri/excelize/v2"
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
	rows, err := db.Query("SELECT turbine, to_char(date,'YYYY-MM-DD'), code, description from fault")
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

	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	index, err := f.NewSheet("Sheet1")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	i := 1
	for _, v := range fault {
		turbine, err := strconv.Atoi(v.Turbine)
		if err != nil {
			log.Printf("Ошибка преоразования номера турбины: %s", err)
			return err
		}
		code, err := strconv.Atoi(v.Code)
		if err != nil {
			log.Printf("Ошибка преоразования кода аварии турбины: %s", err)
			return err
		}

		date, err := time.Parse("2006-01-02", v.Date)
		if err != nil {
			log.Printf("Ошибка форматирования даты аварии: %s", err)
			return err
		}

		f.SetCellValue("Sheet1", fmt.Sprintf("A%v", i), turbine)
		f.SetCellValue("Sheet1", fmt.Sprintf("B%v", i), date.Format("2006-01-02"))
		f.SetCellValue("Sheet1", fmt.Sprintf("C%v", i), code)
		f.SetCellValue("Sheet1", fmt.Sprintf("D%v", i), v.Description)
		i++
	}

	f.SetActiveSheet(index)

	if err := f.SaveAs("Общий список аварий.xlsx"); err != nil {
		log.Printf("Ошибка сохранения файла \"Общий список аварий.xlsx\": %s", err)
		return err
	}
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

	rows, err := db.Query("SELECT turbine, to_char(date,'YYYY-MM-DD'), code, description from fault WHERE turbine=$1", turbine)
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

	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	index, err := f.NewSheet("Sheet1")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	i := 1
	for _, v := range fault {
		turbine, err := strconv.Atoi(v.Turbine)
		if err != nil {
			log.Printf("Ошибка преоразования номера турбины: %s", err)
			return err
		}
		code, err := strconv.Atoi(v.Code)
		if err != nil {
			log.Printf("Ошибка преоразования кода аварии турбины: %s", err)
			return err
		}

		date, err := time.Parse("2006-01-02", v.Date)
		if err != nil {
			log.Printf("Ошибка форматирования даты аварии: %s", err)
			return err
		}

		f.SetCellValue("Sheet1", fmt.Sprintf("A%v", i), turbine)
		f.SetCellValue("Sheet1", fmt.Sprintf("B%v", i), date.Format("2006-01-02"))
		f.SetCellValue("Sheet1", fmt.Sprintf("C%v", i), code)
		f.SetCellValue("Sheet1", fmt.Sprintf("D%v", i), v.Description)
		i++
	}

	f.SetActiveSheet(index)

	if err := f.SaveAs("Список аварий.xlsx"); err != nil {
		log.Printf("Ошибка сохранения файла \"Список аварий.xlsx\": %s", err)
		return err
	}
	return nil
}
