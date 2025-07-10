package models

type DoctorRegisterRequest struct {
	FullName        string `json:"full_name" binding:"required" example:"Иванов Иван Иванович" rus:"ФИО"`
	Login           string `json:"login" binding:"required" example:"doctor_ivanov" rus:"Логин"`
	Password        string `json:"password" binding:"required,min=8" example:"securepassword123" rus:"Пароль"`
	PasswordConfirm string `json:"password_confirm" binding:"required,eqfield=Password" example:"securepassword123" rus:"Подтверждение пароля"`
	Specialization  string `json:"specialization" binding:"required" example:"Терапевт" rus:"Специализация"`
}

type DoctorLoginRequest struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type DoctorResponse struct {
	ID             uint   `json:"id" example:"1" rus:"ID врача"`
	FullName       string `json:"full_name" example:"Иванов Иван Иванович" rus:"ФИО"`
	Specialization string `json:"specialization" example:"Терапевт" rus:"Специализация"`
}
