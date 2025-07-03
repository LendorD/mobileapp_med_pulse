package handlers

type PatientHandler struct {
	//patientService services.PatientService
}

func NewPatientHandler(patientService services.PatientService) *PatientHandler {
	return &PatientHandler{patientService: patientService}
}
