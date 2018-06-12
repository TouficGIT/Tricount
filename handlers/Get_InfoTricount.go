package handlers

import (
	"net/http"
	"encoding/json"
	"github.com/gorilla/mux"
)

func GetInfoTricount(w http.ResponseWriter, r *http.Request) {
	// Execute the query
	params := mux.Vars(r)
	titre := params["titre"]
	rows, err := db.Query("SELECT tricount_id, tricount_titre, user_nom FROM Tricount INNER JOIN Participe ON Tricount.tricount_id = Participe.part_tricount_id INNER JOIN Utilisateur ON Utilisateur.user_id = Participe.part_user_id AND tricount_titre = $1 ;",titre)
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
