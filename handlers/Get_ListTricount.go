package handlers

import (
	"net/http"
	"encoding/json"
)



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
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(mymap)
	}
}