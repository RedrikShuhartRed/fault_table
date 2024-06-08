package main

import (
	"fmt"
	"log"

	"github.com/RedrikShuhartRed/fault_table/controllers"
	"github.com/RedrikShuhartRed/fault_table/db"
	"github.com/RedrikShuhartRed/fault_table/text"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	db.ConnectDB()
	dbs := db.GetDB()

LOOP:
	enter, _ := text.EnterMain()

	switch enter {
	case "1":
		fault, _ := text.CreateFault()
		controllers.AddFault(dbs, fault)
	case "2":
		enter, _ := text.EnterGetAll()
		controllers.GetAll(enter, dbs)
	case "3":
		turbine, _ := text.EnterGetByTurbine()
		controllers.GetByTurbine(turbine, dbs)
	case "4":
		code, _ := text.EnterGetByFault()
		controllers.GetByFault(code, dbs)
	default:
		fmt.Println("Введите цифру от 1 до 4")
		goto LOOP
	}

	db.CloseDB(dbs)
}
