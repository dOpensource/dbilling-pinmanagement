package dbilling_pinmanagement

import (
	"database/sql"
	"strings"
)

func HandleError(e error) {
	if e != nil {
		panic(e)
	}
}

func GetPinInfo(db *sql.DB, pin ...string) []Pin {
	var res *sql.Rows
	var err error
	if len(pin) == 0 {
		res, err = db.Query("SELECT * FROM pin;")
	} else if len(pin) == 1 {
		res, err = db.Query("SELECT * FROM pin WHERE pin=?;", pin[0])
	} else {
		query_arg := ""
		for i, end := 0, len(pin); i < end; i++ {
			query_arg += "'" + strings.ReplaceAll(pin[i], "'", "") + "',"
		}
		query_arg = query_arg[:len(query_arg)-1]
		res, err = db.Query("SELECT * FROM pin WHERE pin IN(?);", query_arg)
	}
	HandleError(err)
	defer res.Close()

	var rows []Pin
	var row Pin
	for res.Next() {
		err = res.Scan(&row.Id, &row.Pin, &row.MinutesPerMonth, &row.UserId, &row.Status)
		HandleError(err)
		rows = append(rows, row)
	}
	return rows
}
