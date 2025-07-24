package entities

type UrologistData struct {
	Complaints []string `json:"complaints"`

	Urinalysis struct {
		Color        string `json:"color"`
		Transparency string `json:"transparency"`
		Protein      string `json:"protein"`
		Glucose      string `json:"glucose"`
		Leukocytes   string `json:"leukocytes"`
		Erythrocytes string `json:"erythrocytes"`
	} `json:"urinalysis"`

	Ultrasound          string `json:"ultrasound"`
	ProstateExamination string `json:"prostate_examination"`
	Diagnosis           string `json:"diagnosis"`
	Treatment           string `json:"treatment"`
}

// ToDocumentWithValues для UrologistData
func (u *UrologistData) ToDocumentWithValues() SpecializationDataDocument {
	fields := []CustomField{
		{
			Name:         "complaints",
			Type:         "array",
			Required:     false,
			Description:  "Жалобы пациента",
			MaxItems:     intPtr(50),
			Format:       "[]string",
			MaxLength:    intPtr(200),
			Example:      []string{"ургентность", "никтурия"},
			DefaultValue: []string{},
			Value:        u.Complaints,
		},
		{
			Name:         "urinalysis",
			Type:         "object",
			Required:     false,
			Description:  "Анализ мочи",
			Example:      map[string]string{"color": "желтый", "protein": "отсутствует"},
			DefaultValue: map[string]string{},
			Value:        u.Urinalysis,
		},
		{
			Name:         "ultrasound",
			Type:         "string",
			Required:     false,
			Description:  "УЗИ органов малого таза",
			MaxLength:    intPtr(1500),
			Example:      "Предстательная железа увеличена",
			DefaultValue: "",
			Value:        u.Ultrasound,
		},
		{
			Name:         "prostate_examination",
			Type:         "string",
			Required:     false,
			Description:  "Пальцевое ректальное исследование предстательной железы",
			MaxLength:    intPtr(1000),
			Example:      "Предстательная железа увеличена, плотная",
			DefaultValue: "",
			Value:        u.ProstateExamination,
		},
		{
			Name:         "diagnosis",
			Type:         "string",
			Required:     false,
			Description:  "Урологический диагноз",
			MaxLength:    intPtr(1000),
			Example:      "Доброкачественная гиперплазия предстательной железы",
			DefaultValue: "",
			Value:        u.Diagnosis,
		},
		{
			Name:         "treatment",
			Type:         "string",
			Required:     false,
			Description:  "Лечение",
			MaxLength:    intPtr(2000),
			Example:      "Альфа-блокаторы, 5-ARI",
			DefaultValue: "",
			Value:        u.Treatment,
		},
	}

	return SpecializationDataDocument{
		DocumentType: "urologist_data",
		Fields:       fields,
	}
}
