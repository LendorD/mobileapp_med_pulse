package converter

import (
	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/entities"
	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/models"
)

func CallToReception(call models.Call) (entities.OneCReception, error) {
	// patientsData, err := json.Marshal(call.Patients)
	// if err != nil {
	// 	return nil, fmt.Errorf("marshal patients: %w", err)
	// }

	return entities.OneCReception{
		// CallID:           call.CallID,
		// Address:          call.Address,
		// Phone:            call.Phone,
		// PatientCount:     call.PatientCount,
		// DoctorID:         call.DoctorID,
		// CallStatus:       string(call.Status),
		// ProcessingStatus: "received",
		// Patients:         patientsData,
	}, nil
}

func ReceptionToCall(reception *entities.OneCReception) (models.Call, error) {
	// var patients []models.Patient
	// if err := json.Unmarshal(reception.Patients, &patients); err != nil {
	// 	return nil, fmt.Errorf("unmarshal patients: %w", err)
	// }

	return models.Call{
		// CallID:       reception.CallID,
		// Address:      reception.Address,
		// Phone:        reception.Phone,
		// PatientCount: reception.PatientCount,
		// Status:       models.CallStatus(reception.CallStatus),
		// Patients:     patients,
		// DoctorID:     reception.DoctorID,
	}, nil
}
