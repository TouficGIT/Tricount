package handlers

import (
	"net/http"
	"encoding/json"
	"fmt"
)

func GetListUser2(w http.ResponseWriter, r *http.Request) {
	// Execute the query
	rows, err := db.Query("SELECT user_id,user_nom FROM Utilisateur;")
	checkErr(err)
	defer rows.Close()
	users := make([]*Utilisateur, 0)
	for rows.Next() {
		user := new(Utilisateur)
		err := rows.Scan(&user.user_id, &user.user_nom)
		checkErr(err)
		fmt.Println(user)
		users = append(users, user)
	}
	err = rows.Err()
	checkErr(err)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}

