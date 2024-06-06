package main

import (
	"log"

	"github.com/RedrikShuhartRed/fault_table/controllers"
	"github.com/RedrikShuhartRed/fault_table/db"
)

func main() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	db.ConnectDB()

	dbs := db.GetDB()

	//controllers.AddFault(dbs)
	//controllers.GetAll(dbs)
	controllers.GetByTurbine(dbs)
}
