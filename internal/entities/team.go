package entities

type Team struct {
	Name            string `json:"team_name"`
	Nationality     string `json:"team_nationality"`
	Principal       string `json:"team_principal"`
	EstablishedYear string `json:"team_established_year"`
	ID              string `json:"id"`
}
