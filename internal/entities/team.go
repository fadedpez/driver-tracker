package entities

type Team struct {
	Name            string `dynamodbav:"team_name"`
	Nationality     string `dynamodbav:"team_nationality"`
	Principal       string `dynamodbav:"team_principal"`
	EstablishedYear string `dynamodbav:"team_established_year"`
	ID              string `dynamodbav:"id"`
}
