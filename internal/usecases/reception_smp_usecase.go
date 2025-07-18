package usecases

import (
	"math"
	"time"

	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/entities"
	"github.com/AlexanderMorozov1919/mobileapp/internal/domain/models"
	"github.com/AlexanderMorozov1919/mobileapp/internal/interfaces"
	"github.com/AlexanderMorozov1919/mobileapp/pkg/errors"
)

type ReceptionSmpUsecase struct {
	recepSmpRepo interfaces.ReceptionSmpRepository
	patientRepo  interfaces.PatientRepository
}

func NewReceptionSmpUsecase(recepRepo interfaces.ReceptionSmpRepository, patientRepo interfaces.PatientRepository) interfaces.ReceptionSmpUsecase {
	return &ReceptionSmpUsecase{
		recepSmpRepo: recepRepo,
		patientRepo:  patientRepo,
	}
}

func (u *ReceptionSmpUsecase) CreateReceptionSMP(input *models.CreateEmergencyRequest) (entities.ReceptionSMP, *errors.AppError) {
	var patient entities.Patient
	var err error
	//Если передан id пациента
	if input.PatientID != nil {
		patient, err = u.patientRepo.GetPatientByID(*input.PatientID)
		if err != nil {
			return entities.ReceptionSMP{}, errors.NewAppError(
				errors.NotFoundErrorCode,
				"Patient not found",
				err,
				true,
			)
		}
	} else if input.Patient != nil {
		// Создаем нового пациента, если id не передан
		parsedTime, parseErr := time.Parse("2006-01-02", input.Patient.BirthDate)
		if parseErr != nil {
			return entities.ReceptionSMP{}, errors.NewAppError(
				errors.InternalServerErrorCode,
				"Invalid birth date format",
				parseErr,
				true,
			)
		}

		newPatient := entities.Patient{
			FullName:  input.Patient.FullName,
			BirthDate: parsedTime,
			IsMale:    input.Patient.IsMale,
		}

		patientID, createErr := u.patientRepo.CreatePatient(newPatient)
		if createErr != nil {
			return entities.ReceptionSMP{}, errors.NewAppError(
				errors.InternalServerErrorCode,
				"Failed to create patient",
				createErr,
				false,
			)
		}

		patient, err = u.patientRepo.GetPatientByID(patientID)
		if err != nil {
			return entities.ReceptionSMP{}, errors.NewAppError(
				errors.InternalServerErrorCode,
				"Failed to get created patient",
				err,
				false,
			)
		}
	} else {
		// Не передан ни ID пациента, ни данные пациента
		return entities.ReceptionSMP{}, errors.NewAppError(
			errors.InvalidDataCode,
			"Either patient_id or patient data must be provided",
			nil,
			true,
		)
	}

	reception := entities.ReceptionSMP{
		EmergencyCallID: input.EmergencyCallID,
		DoctorID:        input.DoctorID,
		PatientID:       patient.ID,
		Diagnosis:       "", // Можно оставить пустым или установить дефолтное значение
		Recommendations: "", // Будет заполнено позже
	}

	createdReceptionID, createErr := u.recepSmpRepo.CreateReceptionSmp(reception)
	if createErr != nil {
		return entities.ReceptionSMP{}, errors.NewAppError(
			errors.InternalServerErrorCode,
			"Failed to create emergency reception",
			createErr,
			false,
		)
	}

	fullReception, err := u.recepSmpRepo.GetReceptionSmpByID(createdReceptionID)
	if err != nil {
		return entities.ReceptionSMP{}, errors.NewAppError(
			errors.InternalServerErrorCode,
			"Failed to get created reception",
			err,
			false,
		)
	}

	return fullReception, nil
}

func (u *ReceptionSmpUsecase) GetReceptionsSMPByEmergencyCall(
	emergencyCallID uint,
	page int,
	perPage int,
) (*models.FilterResponse[[]models.ReceptionSMPShortResponse], error) {
	// Валидация параметров
	if emergencyCallID == 0 {
		return nil, errors.NewAppError(
			errors.InternalServerErrorCode,
			"failed to get call",
			errors.ErrEmptyData,
			true,
		)
	}

	if page < 1 {
		return nil, errors.NewAppError(
			errors.InternalServerErrorCode,
			"page number must be greater than 0",
			errors.ErrDataNotFound,
			true,
		)
	}

	if perPage < 5 {
		return nil, errors.NewAppError(
			errors.InternalServerErrorCode,
			"Perpage number must be greater than 5",
			errors.ErrDataNotFound,
			true,
		)
	}

	// Получение данных из репозитория
	receptions, total, err := u.recepSmpRepo.GetWithPatientsByEmergencyCallID(emergencyCallID, page, perPage)
	if err != nil {
		return nil, errors.NewAppError(
			errors.InternalServerErrorCode,
			"GetWithPatientsByEmergencyCallID failed to get from repo",
			errors.ErrDataNotFound,
			true,
		)
	}

	// Преобразование в DTO
	result := make([]models.ReceptionSMPShortResponse, len(receptions))
	for i, reception := range receptions {
		result[i] = models.ReceptionSMPShortResponse{
			Id:          reception.ID,
			PatientName: reception.Patient.FullName, // Предполагается что Patient предзагружен
			Diagnosis:   reception.Diagnosis,
		}
	}

	// Расчет пагинации
	totalPages := int(math.Ceil(float64(total) / float64(perPage)))

	return &models.FilterResponse[[]models.ReceptionSMPShortResponse]{
		Hits:        result,
		CurrentPage: page,
		TotalPages:  totalPages,
		TotalHits:   int(total),
		HitsPerPage: perPage,
	}, nil
}

func (u *ReceptionSmpUsecase) GetReceptionWithMedServicesByID(id uint) (*models.ReceptionSMPResponse, error) {
	// Валидация
	if id == 0 {
		return nil, errors.NewAppError(
			errors.InternalServerErrorCode,
			"failed to get SMP",
			errors.ErrEmptyData,
			true,
		)
	}

	// Получение данных
	reception, err := u.recepSmpRepo.GetReceptionWithMedServicesByID(id)
	if err != nil {
		if errors.Is(err, errors.ErrDataNotFound) {
			return nil, errors.NewAppError(
				errors.InternalServerErrorCode,
				"GetReceptionWithMedServicesByID failed to get SMP from repo",
				errors.ErrDataNotFound,
				true,
			)
		}
		return nil, errors.NewAppError(
			errors.InternalServerErrorCode,
			"GetReceptionWithMedServicesByID failed to get Med_Services from repo",
			errors.ErrDataNotFound,
			true,
		)
	}

	// Преобразование в DTO
	return &models.ReceptionSMPResponse{
		Id:              reception.ID,
		PatientName:     reception.Patient.FullName,
		Diagnosis:       reception.Diagnosis,
		Recommendations: reception.Recommendations,
		MedServices:     convertMedServicesToResponse(reception.MedServices),
	}, nil
}

func convertMedServicesToResponse(services []entities.MedService) []models.MedServicesResponse {
	result := make([]models.MedServicesResponse, len(services))
	for i, svc := range services {
		result[i] = models.MedServicesResponse{
			Name:  svc.Name,
			Price: svc.Price,
		}
	}
	return result
}
