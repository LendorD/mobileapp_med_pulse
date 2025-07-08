package models

type DoctorResponse struct {
	ID             uint   `json:"id"`
	FullName       string `json:"full_name"`
	Login          string `json:"login"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	Specialization string `json:"specialization"`
}

type CreateDoctorRequest struct {
	FullName       string `json:"full_name"`
	Login          string `json:"login"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	Specialization string `json:"specialization"`
}

type UpdateDoctorRequest struct {
	ID             uint   `json:"id"`
	FullName       string `json:"full_name"`
	Login          string `json:"login"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	Specialization string `json:"specialization"`
}
