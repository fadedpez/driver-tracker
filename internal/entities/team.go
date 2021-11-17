package entities

type Team struct {
	TeamName            string `json:"team_name"`
	TeamNationality     string `json:"team_nationality"`
	TeamPrincipal       string `json:"team_principal"`
	TeamEstablishedYear string `json:"team_established_year"`
	ID                  string `json:"id"`
}
