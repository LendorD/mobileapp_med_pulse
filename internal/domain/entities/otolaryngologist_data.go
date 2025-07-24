package entities

type OtolaryngologistData struct {
	Complaints         []string `json:"complaints"`
	NoseExamination    string   `json:"nose_examination"`
	ThroatExamination  string   `json:"throat_examination"`
	EarExamination     string   `json:"ear_examination"`
	HearingTest        string   `json:"hearing_test"`
	Audiometry         string   `json:"audiometry"`
	VestibularFunction string   `json:"vestibular_function"`
	Endoscopy          string   `json:"endoscopy"`
	Diagnosis          string   `json:"diagnosis"`
	Recommendations    string   `json:"recommendations"`
}

// ToDocumentWithValues для OtolaryngologistData
func (o *OtolaryngologistData) ToDocumentWithValues() SpecializationDataDocument {
	fields := []CustomField{
		{
			Name:         "complaints",
			Type:         "array",
			Required:     false,
			Description:  "Жалобы пациента",
			MaxItems:     intPtr(50),
			Format:       "[]string",
			MaxLength:    intPtr(200),
			Example:      []string{"заложенность носа", "боль в ухе"},
			DefaultValue: []string{},
			Value:        o.Complaints,
		},
		{
			Name:         "nose_examination",
			Type:         "string",
			Required:     false,
			Description:  "Осмотр носа",
			MaxLength:    intPtr(1000),
			Example:      "Отек слизистой, выделения скудные",
			DefaultValue: "",
			Value:        o.NoseExamination,
		},
		{
			Name:         "throat_examination",
			Type:         "string",
			Required:     false,
			Description:  "Осмотр горла",
			MaxLength:    intPtr(1000),
			Example:      "Гиперемия задней стенки глотки",
			DefaultValue: "",
			Value:        o.ThroatExamination,
		},
		{
			Name:         "ear_examination",
			Type:         "string",
			Required:     false,
			Description:  "Осмотр уха",
			MaxLength:    intPtr(1000),
			Example:      "Барабанная перепонка интактна",
			DefaultValue: "",
			Value:        o.EarExamination,
		},
		{
			Name:         "hearing_test",
			Type:         "string",
			Required:     false,
			Description:  "Тест слуха",
			MaxLength:    intPtr(500),
			Example:      "Слух сохранен на оба уха",
			DefaultValue: "",
			Value:        o.HearingTest,
		},
		{
			Name:         "audiometry",
			Type:         "string",
			Required:     false,
			Description:  "Аудиометрия",
			MaxLength:    intPtr(1000),
			Example:      "Порог слышимости в норме",
			DefaultValue: "",
			Value:        o.Audiometry,
		},
		{
			Name:         "vestibular_function",
			Type:         "string",
			Required:     false,
			Description:  "Вестибулярная функция",
			MaxLength:    intPtr(500),
			Example:      "Нистагм отсутствует",
			DefaultValue: "",
			Value:        o.VestibularFunction,
		},
		{
			Name:         "endoscopy",
			Type:         "string",
			Required:     false,
			Description:  "Эндоскопия",
			MaxLength:    intPtr(1000),
			Example:      "Патологии не выявлено",
			DefaultValue: "",
			Value:        o.Endoscopy,
		},
		{
			Name:         "diagnosis",
			Type:         "string",
			Required:     false,
			Description:  "Отоларингологический диагноз",
			MaxLength:    intPtr(1000),
			Example:      "Хронический ринит",
			DefaultValue: "",
			Value:        o.Diagnosis,
		},
		{
			Name:         "recommendations",
			Type:         "string",
			Required:     false,
			Description:  "Рекомендации отоларинголога",
			MaxLength:    intPtr(2000),
			Example:      "Промывание носа, ингаляции",
			DefaultValue: "",
			Value:        o.Recommendations,
		},
	}

	return SpecializationDataDocument{
		DocumentType: "otolaryngologist_data",
		Fields:       fields,
	}
}
