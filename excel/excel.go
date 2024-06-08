package excel

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/RedrikShuhartRed/fault_table/models"
	"github.com/xuri/excelize/v2"
)

func CreateExcel(fault []models.Fault) error {
	f := excelize.NewFile()

	defer f.Close()

	index, err := f.NewSheet("Sheet1")
	if err != nil {
		log.Printf("Ошибка создания листа excel: %s", err)
		return err
	}

	i := 1
	for _, v := range fault {
		turbine, err := strconv.Atoi(v.Turbine)
		if err != nil {
			log.Printf("Ошибка преобразования номера турбины: %s", err)
			return err
		}
		code, err := strconv.Atoi(v.Code)
		if err != nil {
			log.Printf("Ошибка преобразования кода аварии турбины: %s", err)
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
