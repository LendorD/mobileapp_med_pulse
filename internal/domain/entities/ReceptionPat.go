package entities

type ReceptionPat struct {
	Reception
	PatientName string `json:"patient_name"`
}
