package handlers


type Utilisateur struct {
	user_id	string	`json:"id"`
	user_nom	string	`json:"nom"`
}

type Depense struct {
	depense_id	string	`json:"id"`
	depense_type	string	`json:"type"`
	depense_user	string	`json:"user"`
	depense_date	string	`json:"date"`
	depense_montant	string	`json:"montant"`
}

type Tricount struct {
	tricount_id	string	`json:"id"`
	tricount_titre	string	`json:"titre"`
	tricount_desc	string	`json:"description"`
}

type Participe struct {
	part_tricount_id string `json:"id"`
	part_user_id string `json:"id"`
	part_balance string `json:"balance"`
}