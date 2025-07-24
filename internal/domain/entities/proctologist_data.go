package entities

type ProctologistData struct {
	Complaints         []string `json:"complaints"`
	DigitalExamination string   `json:"digital_examination"`
	Rectoscopy         string   `json:"rectoscopy"`
	Colonoscopy        string   `json:"colonoscopy"`
	Hemorrhoids        bool     `json:"hemorrhoids"`
	AnalFissure        bool     `json:"anal_fissure"`
	Paraproctitis      bool     `json:"paraproctitis"`
	Tumor              bool     `json:"tumor"`
	Diagnosis          string   `json:"diagnosis"`
	Recommendations    string   `json:"recommendations"`
}

// ToDocumentWithValues для ProctologistData
func (p *ProctologistData) ToDocumentWithValues() SpecializationDataDocument {
	fields := []CustomField{
		{
			Name:         "complaints",
			Type:         "array",
			Required:     false,
			Description:  "Жалобы пациента",
			MaxItems:     intPtr(50),
			Format:       "[]string",
			MaxLength:    intPtr(200),
			Example:      []string{"боль в области ануса"},
			DefaultValue: []string{},
			Value:        p.Complaints,
		},
		{
			Name:         "digital_examination",
			Type:         "string",
			Required:     false,
			Description:  "Пальцевое ректальное исследование",
			MaxLength:    intPtr(1000),
			Example:      "болезненность при пальпации",
			DefaultValue: "",
			Value:        p.DigitalExamination,
		},
		{
			Name:         "rectoscopy",
			Type:         "string",
			Required:     false,
			Description:  "Ректороманоскопия",
			MaxLength:    intPtr(1500),
			Example:      "геморроидальные узлы видны",
			DefaultValue: "",
			Value:        p.Rectoscopy,
		},
		{
			Name:         "colonoscopy",
			Type:         "string",
			Required:     false,
			Description:  "Колоноскопия",
			MaxLength:    intPtr(2000),
			Example:      "патологии не выявлено",
			DefaultValue: "",
			Value:        p.Colonoscopy,
		},
		{
			Name:         "hemorrhoids",
			Type:         "boolean",
			Required:     false,
			Description:  "Геморрой",
			Example:      true,
			DefaultValue: false,
			Value:        p.Hemorrhoids,
		},
		{
			Name:         "anal_fissure",
			Type:         "boolean",
			Required:     false,
			Description:  "Анальная трещина",
			Example:      false,
			DefaultValue: false,
			Value:        p.AnalFissure,
		},
		{
			Name:         "paraproctitis",
			Type:         "boolean",
			Required:     false,
			Description:  "Парарактит",
			Example:      false,
			DefaultValue: false,
			Value:        p.Paraproctitis,
		},
		{
			Name:         "tumor",
			Type:         "boolean",
			Required:     false,
			Description:  "Опухоль",
			Example:      false,
			DefaultValue: false,
			Value:        p.Tumor,
		},
		{
			Name:         "diagnosis",
			Type:         "string",
			Required:     false,
			Description:  "Проктологический диагноз",
			MaxLength:    intPtr(1000),
			Example:      "Хронический геморрой",
			DefaultValue: "",
			Value:        p.Diagnosis,
		},
		{
			Name:         "recommendations",
			Type:         "string",
			Required:     false,
			Description:  "Рекомендации проктолога",
			MaxLength:    intPtr(2000),
			Example:      "Диета, местная терапия",
			DefaultValue: "",
			Value:        p.Recommendations,
		},
	}

	return SpecializationDataDocument{
		DocumentType: "proctologist_data",
		Fields:       fields,
	}
}
