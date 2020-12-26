package handlers

// AddList is a contract to add a list to database
type AddList struct {
	Title    string `json:title`
	Asignees []int  `json:asignees`
}

type AddUser struct {
	Email     string `json:"email"`
	Password  string `json:"password"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}
