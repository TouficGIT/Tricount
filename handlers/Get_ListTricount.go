package handlers

import (
	"net/http"
	"encoding/json"
	"fmt"
	"database/sql"
)

var db *sql.DB

type Tricount struct {
	tricount_id	string	`json:"id"`
	tricount_titre	string	`json:"titre"`
	tricount_desc	string	`json:"description"`
}

func GetListTricount(w http.ResponseWriter, r *http.Request) {
	// Execute the query
	rows, err := db.Query("SELECT * FROM Tricount;")
	checkErr(err)
	defer rows.Close()
	cols, _ := rows.Columns()
	colnames, _ := rows.Columns()
	vals := make([]interface{}, len(cols))

	for i, _ := range cols {
		vals[i] = &cols[i]
	}

	mymap := make(map[string]interface{})

	for i, val := range vals {
		mymap[colnames[i]] = val
	}

	for rows.Next() {
		err = rows.Scan(vals...)
		json, _ := json.Marshal(mymap)
		fmt.Fprintf(w,"%s\n",json)
	}
}