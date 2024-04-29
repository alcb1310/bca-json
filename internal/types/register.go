package types

type RegisterInformation struct {
	RUC       string `json:"ruc"`
	Name      string `json:"name"`
	Employees *uint  `json:"employees"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	UserName  string `json:"username"`
}
