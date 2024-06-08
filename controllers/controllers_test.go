package controllers

import (
	"testing"

	"github.com/RedrikShuhartRed/fault_table/db"
	"github.com/RedrikShuhartRed/fault_table/models"
	"github.com/RedrikShuhartRed/fault_table/query"
	"github.com/stretchr/testify/require"
)

func getTestFault() models.Fault {
	return models.Fault{
		Turbine:     "555",
		Date:        "2000-01-01",
		Code:        "555",
		Description: "Description",
	}
}

func getTestDate() (string, string) {
	return "1990.01.01", "2099.01.01"
}

func TestAddFault(t *testing.T) {
	err := db.ConnectDB()
	require.NoError(t, err)
	dbs := db.GetDB()

	fault := getTestFault()

	err = AddFault(dbs, fault)
	require.NoError(t, err)

	rows, err := dbs.Query("SELECT turbine, to_char(date,'YYYY-MM-DD'), code, description FROM fault ORDER BY id DESC LIMIT 1")
	require.NoError(t, err)
	defer rows.Close()

	var alarm models.Fault
	for rows.Next() {
		err := rows.Scan(&alarm.Turbine, &alarm.Date, &alarm.Code, &alarm.Description)
		require.NoError(t, err)
	}
	require.Equal(t, fault, alarm)

	err = db.CloseDB(dbs)
	require.NoError(t, err)

}

func TestGetAll(t *testing.T) {
	err := db.ConnectDB()
	require.NoError(t, err)
	dbs := db.GetDB()
	begin, end := getTestDate()

	fault, err := GetAll("1", begin, end, dbs)
	require.NoError(t, err)
	row, err := dbs.Query(query.QueryGetAllByDateASC)
	require.NoError(t, err)
	count := 0
	for row.Next() {
		count++
	}
	require.Equal(t, count, len(fault))

	fault, err = GetAll("2", begin, end, dbs)
	require.NoError(t, err)
	row, err = dbs.Query(query.QueryGetAllByDateDESC)
	require.NoError(t, err)
	count = 0
	for row.Next() {
		count++
	}
	require.Equal(t, count, len(fault))

	fault, err = GetAll("3", begin, end, dbs)
	require.NoError(t, err)
	row, err = dbs.Query(query.QueryGetAllBetweenDate, begin, end)
	require.NoError(t, err)
	count = 0
	for row.Next() {
		count++
	}
	require.Equal(t, count, len(fault))

	fault, err = GetAll("4", begin, end, dbs)
	require.NoError(t, err)
	row, err = dbs.Query(query.QueryGetAllByTurbine)
	require.NoError(t, err)
	count = 0
	for row.Next() {
		count++
	}
	require.Equal(t, count, len(fault))

	fault, err = GetAll("5", begin, end, dbs)
	require.NoError(t, err)
	row, err = dbs.Query(query.QueryGetAllByCode)
	require.NoError(t, err)
	count = 0
	for row.Next() {
		count++
	}
	require.Equal(t, count, len(fault))

	err = db.CloseDB(dbs)
	require.NoError(t, err)

}

func TestGetByTurbine(t *testing.T) {
	err := db.ConnectDB()
	require.NoError(t, err)
	dbs := db.GetDB()

	fault := getTestFault()

	err = GetByTurbine(fault.Turbine, dbs)
	require.NoError(t, err)
	rows, err := dbs.Query(query.QueryGetByTurbine, fault.Turbine)
	require.NoError(t, err)

	var faultList []models.Fault

	for rows.Next() {
		alarm := models.Fault{}

		err := rows.Scan(&alarm.Turbine, &alarm.Date, &alarm.Code, &alarm.Description)
		require.NoError(t, err)

		faultList = append(faultList, alarm)
	}
	for _, v := range faultList {
		require.Equal(t, v.Turbine, fault.Turbine)
	}
}
func TestGetByFault(t *testing.T) {
	err := db.ConnectDB()
	require.NoError(t, err)
	dbs := db.GetDB()

	fault := getTestFault()

	err = GetByTurbine(fault.Turbine, dbs)
	require.NoError(t, err)
	rows, err := dbs.Query(query.QueryGetByFault, fault.Code)
	require.NoError(t, err)

	var faultList []models.Fault

	for rows.Next() {
		alarm := models.Fault{}

		err := rows.Scan(&alarm.Turbine, &alarm.Date, &alarm.Code, &alarm.Description)
		require.NoError(t, err)

		faultList = append(faultList, alarm)
	}
	for _, v := range faultList {
		require.Equal(t, v.Code, fault.Code)
	}

}
