package entities

type Team struct {
	TeamName            string `json:"team_name"`
	TeamNationality     string `json:"team_nationality"`
	TeamPrincipal       string `json:"team_principal"`
	TeamEstablishedYear string `json:"team_established_year"`
	TeamYearsActive     string `json:"team_years_active"`
}
