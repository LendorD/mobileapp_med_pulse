package entities

// Статусы для приёма стационара
type HospitalReceptionStatus string

const (
	HospitalReceptionStatusScheduled HospitalReceptionStatus = "scheduled"
	HospitalReceptionStatusCompleted HospitalReceptionStatus = "completed"
	HospitalReceptionStatusCancelled HospitalReceptionStatus = "cancelled"
	HospitalReceptionStatusNoShow    HospitalReceptionStatus = "no_show"
)

var allowedHospitalReceptionStatuses = map[HospitalReceptionStatus]struct{}{
	HospitalReceptionStatusScheduled: {},
	HospitalReceptionStatusCompleted: {},
	HospitalReceptionStatusCancelled: {},
	HospitalReceptionStatusNoShow:    {},
}

func IsValidHospitalReceptionStatus(status string) bool {
	_, ok := allowedHospitalReceptionStatuses[HospitalReceptionStatus(status)]
	return ok
}

type StatusUpdateInput struct {
	Status string `json:"status" binding:"required"`
}
