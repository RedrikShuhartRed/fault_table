package query

const (
	QueryAddFault          = `INSERT INTO fault (turbine, date, code, description) VALUES ($1,$2,$3,$4)`
	QueryGetAllByDateASC   = "SELECT turbine, to_char(date,'YYYY-MM-DD'), code, description from fault ORDER BY date ASC"
	QueryGetAllByDateDESC  = "SELECT turbine, to_char(date,'YYYY-MM-DD'), code, description from fault ORDER BY date DESC"
	QueryGetAllBetweenDate = `SELECT turbine, to_char(date,'YYYY-MM-DD'), code, description from fault
	WHERE date BETWEEN $1 AND $2 ORDER BY date DESC`
	QueryGetAllByTurbine = "SELECT turbine, to_char(date,'YYYY-MM-DD'), code, description from fault ORDER BY turbine ASC, date DESC"
	QueryGetAllByCode    = "SELECT turbine, to_char(date,'YYYY-MM-DD'), code, description from fault ORDER BY code, date DESC"
	QueryGetByTurbine    = "SELECT turbine, to_char(date,'YYYY-MM-DD'), code, description from fault WHERE turbine=$1 ORDER BY date DESC"
	QueryGetByFault      = "SELECT turbine, to_char(date,'YYYY-MM-DD'), code, description from fault WHERE code=$1 ORDER BY date DESC"
)
