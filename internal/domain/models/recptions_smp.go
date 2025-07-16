package models

type ReceptionSMPShortResponse struct {
	Id          uint   `json:"id"`
	PatientName string `json:"patient_name" example:"Иванов Иван"` // ФИО пациента
	Diagnosis   string `json:"diagnosis" example:"ОРВИ"`
}

type ReceptionSMPResponse struct {
	Id              uint   `json:"id"`
	PatientName     string `json:"patient_name" example:"Иванов Иван"` // ФИО пациента
	Diagnosis       string `json:"diagnosis" example:"ОРВИ" rus:"Диагноз"`
	Recommendations string `json:"recommendations" example:"Постельный режим" rus:"Рекомендации"`
	MedServices []MedServicesResponse  `json:"med_services" rus:"Медицинские услуги"`
}
