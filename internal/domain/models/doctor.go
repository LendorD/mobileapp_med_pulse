package models

// DoctorResponse - полная информация о враче
// @Description Содержит все данные о враче включая идентификационные и контактные данные
type DoctorResponse struct {
	ID             uint   `json:"id" example:"1"`                           // Уникальный идентификатор врача
	FullName       string `json:"full_name" example:"Иванов Иван Иванович"` // Полное имя врача
	Login          string `json:"login" example:"+79123456789"`             // Логин для входа (обычно телефон)
	Password       string `json:"password" example:"qwerty123"`             // Хэш пароля (не должен возвращаться в API)
	Specialization string `json:"specialization" example:"Терапевт"`        // Медицинская специализация
}

type DoctorInfoResponse struct {
	FullName       string `json:"full_name" example:"Иванов Иван Иванович"` // Полное имя врача
	Specialization string `json:"specialization" example:"Терапевт"`        // Медицинская специализация
}

// CreateDoctorRequest - запрос на создание врача
// @Description Используется для регистрации нового врача в системе
type CreateDoctorRequest struct {
	FullName       string `json:"full_name" binding:"required" example:"Иванов Иван Иванович"` // ФИО врача (обязательное)
	Login          string `json:"login" binding:"required" example:"+79123456789"`             // Логин (обязательное)
	Password       string `json:"password" binding:"required" example:"qwerty123"`             // Пароль (обязательное)
	Specialization string `json:"specialization" binding:"required" example:"Терапевт"`        // Специализация (обязательное)
}

// UpdateDoctorRequest - запрос на обновление данных врача
// @Description Используется для изменения информации о враче
type UpdateDoctorRequest struct {
	ID             uint   `json:"id" example:"1"`                           // ID врача для обновления
	FullName       string `json:"full_name" example:"Иванов Иван Иванович"` // Новое ФИО
	Login          string `json:"login" example:"+79123456789"`             // Новый логин
	Password       string `json:"password" example:"newpassword123"`        // Новый пароль
	Specialization string `json:"specialization" example:"Хирург"`          // Новая специализация
}
