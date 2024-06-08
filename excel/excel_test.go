package excel

import (
	"testing"

	"github.com/RedrikShuhartRed/fault_table/models"
	"github.com/stretchr/testify/require"
)

func GetTestFault() []models.Fault {
	return []models.Fault{
		{Turbine: "555", Date: "1990-01-01", Code: "555", Description: "Description"},
		{Turbine: "666", Date: "2001-01-01", Code: "666", Description: "Description"},
	}
}
func TestCreateExcel(t *testing.T) {
	fault := GetTestFault()
	err := CreateExcel(fault)
	require.NoError(t, err)
}

func TestCreateExcelEmptyValue(t *testing.T) {
	fault := []models.Fault{
		{Turbine: "", Date: "1990-01-01", Code: "555", Description: "Description"},
		{Turbine: "", Date: "2001-01-01", Code: "666", Description: "Description"},
	}
	err := CreateExcel(fault)
	require.Error(t, err)
	fault = []models.Fault{
		{Turbine: "555", Date: "1990-01-01", Code: "", Description: "Description"},
		{Turbine: "666", Date: "2001-01-01", Code: "", Description: "Description"},
	}
	err = CreateExcel(fault)
	require.Error(t, err)

	fault = []models.Fault{
		{Turbine: "555", Date: "", Code: "555", Description: "Description"},
		{Turbine: "666", Date: "", Code: "666", Description: "Description"},
	}
	err = CreateExcel(fault)
	require.Error(t, err)

}
