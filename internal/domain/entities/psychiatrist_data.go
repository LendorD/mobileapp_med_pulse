package entities

type PsychiatristData struct {
	MentalStatus   string `json:"mental_status"`
	Mood           string `json:"mood"`
	Affect         string `json:"affect"`
	ThoughtProcess string `json:"thought_process"`
	ThoughtContent string `json:"thought_content"`
	Perception     string `json:"perception"`
	Cognition      string `json:"cognition"`
	Insight        string `json:"insight"`
	Judgment       string `json:"judgment"`

	RiskAssessment struct {
		Suicide  bool `json:"suicide"`
		SelfHarm bool `json:"self_harm"`
		Violence bool `json:"violence"`
	} `json:"risk_assessment"`

	DiagnosisICD string `json:"diagnosis_icd"`
	TherapyPlan  string `json:"therapy_plan"`
}

// ToDocumentWithValues для PsychiatristData
func (p *PsychiatristData) ToDocumentWithValues() SpecializationDataDocument {
	fields := []CustomField{
		{
			Name:         "mental_status",
			Type:         "string",
			Required:     false,
			Description:  "Психический статус",
			MaxLength:    intPtr(500),
			Example:      "Ясен, ориентирован в месте и времени",
			DefaultValue: "",
			Value:        p.MentalStatus,
		},
		{
			Name:         "mood",
			Type:         "string",
			Required:     false,
			Description:  "Настроение",
			MaxLength:    intPtr(200),
			Example:      "депрессивное",
			DefaultValue: "",
			Value:        p.Mood,
		},
		{
			Name:         "affect",
			Type:         "string",
			Required:     false,
			Description:  "Аффект",
			MaxLength:    intPtr(200),
			Example:      "сужен",
			DefaultValue: "",
			Value:        p.Affect,
		},
		{
			Name:         "thought_process",
			Type:         "string",
			Required:     false,
			Description:  "Мышление (процесс)",
			MaxLength:    intPtr(500),
			Example:      "логично, последовательно",
			DefaultValue: "",
			Value:        p.ThoughtProcess,
		},
		{
			Name:         "thought_content",
			Type:         "string",
			Required:     false,
			Description:  "Мышление (содержание)",
			MaxLength:    intPtr(1000),
			Example:      "бредовые идеи преследования",
			DefaultValue: "",
			Value:        p.ThoughtContent,
		},
		{
			Name:         "perception",
			Type:         "string",
			Required:     false,
			Description:  "Восприятие",
			MaxLength:    intPtr(500),
			Example:      "галлюцинации отрицает",
			DefaultValue: "",
			Value:        p.Perception,
		},
		{
			Name:         "cognition",
			Type:         "string",
			Required:     false,
			Description:  "Познавательные функции",
			MaxLength:    intPtr(500),
			Example:      "память сохр, внимание снижено",
			DefaultValue: "",
			Value:        p.Cognition,
		},
		{
			Name:         "insight",
			Type:         "string",
			Required:     false,
			Description:  "Критика болезни",
			MaxLength:    intPtr(300),
			Example:      "частично сохранена",
			DefaultValue: "",
			Value:        p.Insight,
		},
		{
			Name:         "judgment",
			Type:         "string",
			Required:     false,
			Description:  "Критика",
			MaxLength:    intPtr(300),
			Example:      "сохранена",
			DefaultValue: "",
			Value:        p.Judgment,
		},
		{
			Name:         "risk_assessment",
			Type:         "object",
			Required:     false,
			Description:  "Оценка рисков",
			Example:      map[string]bool{"suicide": false, "self_harm": false, "violence": false},
			DefaultValue: map[string]bool{"suicide": false, "self_harm": false, "violence": false},
			Value:        p.RiskAssessment,
		},
		{
			Name:         "diagnosis_icd",
			Type:         "string",
			Required:     false,
			Description:  "Диагноз по МКБ",
			MaxLength:    intPtr(200),
			Example:      "F32.0 - депрессивный эпизод легкой степени",
			DefaultValue: "",
			Value:        p.DiagnosisICD,
		},
		{
			Name:         "therapy_plan",
			Type:         "string",
			Required:     false,
			Description:  "План терапии",
			MaxLength:    intPtr(2000),
			Example:      "Фармакотерапия, психотерапия",
			DefaultValue: "",
			Value:        p.TherapyPlan,
		},
	}

	return SpecializationDataDocument{
		DocumentType: "psychiatrist_data",
		Fields:       fields,
	}
}
