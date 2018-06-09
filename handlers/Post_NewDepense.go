package handlers

import (
	"net/http"
	"github.com/gorilla/mux"
	"fmt"
	"time"
	"strconv"
)

func NewDepense(w http.ResponseWriter, r *http.Request){
	currentTime := time.Now()
	date := currentTime.Format("2006-01-02")
	params := mux.Vars(r)
	nom := params["nom"]
	montant := params["montant"]
	dep_type := params["type"]
	tricount := params["id"]
	var depense_id int
	var user_id int
	var depense_montant float64
	depense_montant, err := strconv.ParseFloat(params["montant"],64)
	checkErr(err)

	rows, err := db.Query("SELECT COUNT(tricount_id) AS count FROM Tricount WHERE tricount_id = $1;", tricount)
	checkErr(err)
	count := checkCount(rows)

	if count != 1 {
		fmt.Fprintln(w, "Tricount inexistant :", tricount)
		return
	}

	rows2, err := db.Query("SELECT COUNT(user_nom) AS count FROM Utilisateur u INNER JOIN Participe p ON u.user_id = p.part_user_id INNER JOIN Tricount t ON p.part_tricount_id = t.tricount_id WHERE t.tricount_id = $1 AND u.user_nom = $2;", tricount,nom)
	checkErr(err)
	count2 := checkCount(rows2)

	if count2 != 1 {
		fmt.Fprintln(w, "L'utilisateur n'appartient pas à ce Tricount")
		return
	}

	insert := db.QueryRow("INSERT INTO Depense (depense_type, depense_user, depense_date, depense_montant) VALUES ($1,$2,$3,$4) RETURNING depense_id;",dep_type,nom,date,montant).Scan(&depense_id)
	checkErr(insert)
	insert2 := db.QueryRow("INSERT INTO Comporte (comp_tricount_id, comp_depense_id) VALUES ($1,$2) RETURNING comp_tricount_id;",tricount,depense_id).Scan(&tricount)
	checkErr(insert2)

	rows3, err := db.Query("SELECT COUNT(part_user_id) FROM Participe WHERE part_tricount_id = $1;", tricount)
	checkErr(err)
	var nbUser float64
	nbUser = checkCount(rows3)

	if nbUser > 1 {
		balance := depense_montant / nbUser

		row4 := db.QueryRow("SELECT user_id FROM Utilisateur WHERE user_nom = $1;", nom).Scan(&user_id)
		checkErr(row4)
		var part_balance float64
		row5 := db.QueryRow("SELECT part_balance FROM Participe WHERE part_user_id = $1;", user_id).Scan(&part_balance)
		checkErr(row5)

		part_balance += balance

		rows6, err := db.Query("UPDATE Participe SET part_balance = $1 WHERE part_user_id = $2;", part_balance, user_id)
		checkErr(err)
		checkCount(rows6)
		/*
	rows6, err := db.Query("UPDATE Participe SET part_balance = $1 WHERE part_user_id != $2 and part_tricount_id = $3;",balance,user_id,tricount)
	checkErr(err)
	checkCount(rows6)
	*/

		rows7, err := db.Query(" SELECT part_balance,part_user_id FROM Participe WHERE part_user_id != $1 and part_tricount_id = $2;", user_id, tricount)
		defer rows7.Close()
		for rows7.Next() {
			var part_user_id int
			var part_balance2 float64
			err = rows7.Scan(&part_balance2, &part_user_id)
			part_balance2 -= balance
			rows8, err := db.Query("UPDATE Participe SET part_balance = $1 WHERE part_user_id = $2 and part_tricount_id = $3;", part_balance2, part_user_id, tricount)
			checkErr(err)
			checkCount(rows8)
		}
		checkCount(rows7)
		fmt.Fprintln(w, "Mise à jour du balance des utilisateurs : ", balance)
	}
	fmt.Fprintln(w, "Ajout du montant : ", montant, " pour l'utilisateur : ", nom," dans le Tricount : ", tricount)
}
