package handlers


type Utilisateur struct {
	user_id	int
	user_nom	string
}

type Depense struct {
	depense_id	int
	depense_type	string
	depense_user	string
	depense_date	string
	depense_montant	string
}

type Tricount struct {
	tricount_id	int
	tricount_titre	string
}

type Participe struct {
	part_tricount_id int
	part_user_id string
	part_balance string
}

type Comporte struct {
	comp_tricount_id string
	comp_depense_id string
}