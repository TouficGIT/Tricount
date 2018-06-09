package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
)

func GetBalanceTricount(w http.ResponseWriter, r *http.Request) {
	// Execute the query
	params := mux.Vars(r)
	id := params["id"]
	rows, err := db.Query("SELECT user_nom,part_balance,tricount_titre,tricount_id FROM Depense INNER JOIN Comporte ON Depense.depense_id = Comporte.comp_depense_id INNER JOIN Tricount ON Tricount.tricount_id = Comporte.comp_tricount_id INNER JOIN Participe ON Participe.part_tricount_id = Tricount.tricount_id INNER JOIN Utilisateur ON Utilisateur.user_id = Participe.part_user_id AND tricount_id = $1;",id)
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
