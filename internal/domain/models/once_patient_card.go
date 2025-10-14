package models

// PatientCard — полная карточка пациента
type PatientCard struct {
	DisplayName         string      `json:"display_name"`                   // Представление (ФИО или иное отображаемое имя)
	Age                 string      `json:"age"`                            // Возраст (например, "35 лет")
	BirthDate           string      `json:"birth_date"`                     // Дата рождения (в формате, как из 1С, например "14.10.1990")
	MobilePhone         string      `json:"mobile_phone"`                   // Телефон (мобильный)
	AdditionalPhone     string      `json:"additional_phone"`               // Сотовый телефон (доп. телефон)
	Address             string      `json:"address"`                        // Адрес (фактический)
	Email               string      `json:"email"`                          // Адрес электронной почты
	LegalRepresentative *ClientRef  `json:"legal_representative,omitempty"` // Законный представитель (ссылка на справочник клиента)
	Relative            *Relative   `json:"relative,omitempty"`             // Родственник
	Workplace           string      `json:"workplace"`                      // Место работы
	AttendingDoctor     Doctor      `json:"attending_doctor"`               // Лечащий врач
	Snils               string      `json:"snils"`                          // СНИЛС
	Policy              Policy      `json:"policy"`                         // Полис
	Certificate         Certificate `json:"certificate"`                    // Сертификат
}

type UpdateMedicalCardRequest struct {
	PatientID string `json:"patient_id" binding:"required"`
}

// ClientRef — ссылка на клиента из справочника (законный представитель)
type ClientRef struct {
	ID   string `json:"id"`   // Уникальный идентификатор клиента в справочнике
	Name string `json:"name"` // Наименование / ФИО
}

// Relative — информация о родственнике
type Relative struct {
	Status string `json:"status"` // Статус родственника (например, "Отец", "Мать", "Супруг")
	Name   string `json:"name"`   // Наименование / ФИО родственника
}

// Doctor — лечащий врач
type Doctor struct {
	FullName           string `json:"full_name"`             // ФИО врача
	PolicyOrCertNumber string `json:"policy_or_cert_number"` // Номер полиса/сертификата (уточни, что именно)
	AttachmentStart    string `json:"attachment_start"`      // Начало прикрепления (дата как строка)
	AttachmentEnd      string `json:"attachment_end"`        // Конец прикрепления (дата как строка)
	Clinic             string `json:"clinic"`                // Клиника
}
