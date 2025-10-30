package converter

import (
	"time"

	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/entities"
	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/models"
)

func CallToReception(call models.Call) (entities.OneCReception, error) {
	now := time.Now()

	// Конвертируем Patients
	var patients []entities.PatientData
	for _, p := range call.Patients {
		patients = append(patients, entities.PatientData{
			PatientID:   p.PatientID,
			FullName:    p.FullName,
			Age:         p.Age,
			BirthDate:   p.BirthDate,
			MobilePhone: p.MobilePhone,
			Policy:      p.Policy,
			Certificate: p.Certificate,
		})
	}

	// Аналогично для других вложенных структур
	doctor := entities.DoctorData{
		Name:           call.Doctor.Name,
		Specialization: call.Doctor.Specialization,
	}

	var receptions []entities.Receptions
	for _, r := range call.Receptions {
		receptions = append(receptions, entities.Receptions{Data: r.Data})
	}

	flg := entities.Flg{
		CreatedAt:    call.Flg.CreatedAt,
		PatientID:    call.Flg.PatientID,
		Organization: call.Flg.Organization,
		Number:       call.Flg.Number,
		Result:       call.Flg.Result,
		Date:         call.Flg.Date,
		// FileID и File остаются пустыми (заполнятся позже при загрузке файла)
	}

	analysis := entities.Analysis{
		Code:  call.Analysis.Code,
		Title: call.Analysis.Title,
		Price: call.Analysis.Price,
	}

	return entities.OneCReception{
		CallID:      call.CallID,
		Address:     call.Address,
		Phone:       call.Phone,
		Status:      entities.CallStatus(call.Status),
		CreatedAt:   now,
		UpdatedAt:   now,
		MedServices: call.MedServices,
		Patients:    patients,
		Doctor:      doctor,
		Receptions:  receptions,
		Flg:         flg,
		Analysis:    analysis,
	}, nil
}

func ReceptionToCall(reception *entities.OneCReception) (models.Call, error) {
	// Конвертируем Patients
	var patients []entities.PatientData
	for _, p := range reception.Patients {
		patients = append(patients, entities.PatientData{
			PatientID:   p.PatientID,
			FullName:    p.FullName,
			Age:         p.Age,
			BirthDate:   p.BirthDate,
			MobilePhone: p.MobilePhone,
			Policy:      p.Policy,
			Certificate: p.Certificate,
		})
	}

	// Аналогично для других
	doctor := entities.DoctorData{
		Name:           reception.Doctor.Name,
		Specialization: reception.Doctor.Specialization,
	}

	var receptions []entities.Receptions
	for _, r := range reception.Receptions {
		receptions = append(receptions, entities.Receptions{Data: r.Data})
	}

	flg := entities.Flg{
		CreatedAt:    reception.Flg.CreatedAt,
		PatientID:    reception.Flg.PatientID,
		Organization: reception.Flg.Organization,
		Number:       reception.Flg.Number,
		Result:       reception.Flg.Result,
		Date:         reception.Flg.Date,
	}

	analysis := entities.Analysis{
		Code:  reception.Analysis.Code,
		Title: reception.Analysis.Title,
		Price: reception.Analysis.Price,
	}

	return models.Call{
		CallID:      reception.CallID,
		Address:     reception.Address,
		Phone:       reception.Phone,
		Status:      entities.CallStatus(reception.Status),
		MedServices: reception.MedServices,
		Patients:    patients,
		Doctor:      doctor,
		Receptions:  receptions,
		Flg:         flg,
		Analysis:    analysis,
	}, nil
}
