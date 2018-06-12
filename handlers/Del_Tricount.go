package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
)

func DelTricount(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	tricount := params["nom"]
	var tricount_id int

	rows, err := db.Query("SELECT COUNT(tricount_id) AS count FROM Tricount WHERE tricount_titre = $1;", tricount)
	checkErr(err)
	count := checkCount(rows)

	if count != 1 {
		fmt.Fprintln(w, "Tricount inexistant :", tricount)
		return
	}

	row2 :=  db.QueryRow("SELECT tricount_id FROM Tricount WHERE tricount_titre = $1;",tricount).Scan(&tricount_id)
	checkErr(row2)

	del, err := db.Exec("DELETE FROM Comporte WHERE comp_tricount_id = $1;",tricount_id)
	checkErr(err)
	test, err := del.RowsAffected()
	checkErr(err)
	fmt.Println(test)

	del2, err := db.Exec("DELETE FROM Participe WHERE part_tricount_id = $1;",tricount_id)
	checkErr(err)
	test2, err := del2.RowsAffected()
	checkErr(err)
	fmt.Println(test2)

	del3, err := db.Exec("DELETE FROM Tricount WHERE tricount_id = $1;",tricount_id)
	checkErr(err)
	test3, err := del3.RowsAffected()
	checkErr(err)
	fmt.Println(test3)

	fmt.Fprintln(w, "Suppression du tricount : ", tricount)
}