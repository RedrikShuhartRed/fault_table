package controller_test

import (
	"testing"

	"github.com/RedrikShuhartRed/fault_table/controllers"
	"github.com/RedrikShuhartRed/fault_table/db"
	"github.com/RedrikShuhartRed/fault_table/models"
	"github.com/stretchr/testify/require"
)

func getTestFault() models.Fault {
	return models.Fault{
		Turbine:     "555",
		Date:        "2000.01.01",
		Code:        "555",
		Description: "Описание ошибки",
	}
}

func TestAddFault(t *testing.T) {
	err := db.ConnectDB()
	require.NoError(t, err)
	dbs := db.GetDB()

	fault := getTestFault()

	err = controllers.AddFault(dbs, fault)
	require.NoError(t, err)

}
