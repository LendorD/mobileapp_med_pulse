package models

type MedServicesResponse struct {
	ID    uint   `json:"id"`
	Name  string `json:"name" example:"EKG"`
	Price uint   `json:"price" example:"100"`
}
