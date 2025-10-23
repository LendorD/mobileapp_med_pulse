package entities

type PatientData struct {
	PatientID   string
	FullName    string
	Age         int
	BirthDate   string
	MobilePhone string
	Policy      Policy
	Certificate Certificate
}
