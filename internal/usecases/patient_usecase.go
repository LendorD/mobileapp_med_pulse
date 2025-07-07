package usecases

/*
type PatientUsecase struct {
	repo interfaces.PatientRepository
}

func NewPatientUsecase(repo interfaces.PatientRepository) interfaces.PatientUsecase {
	return &PatientUsecase{repo: repo}
}

func (u *PatientUsecase) Create(input models.CreatePatientRequest) (entities.Patient, *errors.AppError) {
	patient := entities.Patient{
		FullName:  input.FullName,
		BirthDate: input.BirthDate,
		IsMale:    input.IsMale,
	}

	createdPatient, err := u.repo.Create(&patient)
	if err != nil {
		return entities.Patient{}, errors.NewDBError("failed to create patient", err)
	}
	return *createdPatient, nil
}

func (u *PatientUsecase) GetByID(id uint) (entities.Patient, *errors.AppError) {
	patient, err := u.repo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return entities.Patient{}, errors.NewNotFoundError("patient not found")
		}
		return entities.Patient{}, errors.NewDBError("failed to get patient", err)
	}
	return *patient, nil
}

func (u *PatientUsecase) Update(input models.UpdatePatientRequest) (entities.Patient, *errors.AppError) {
	patient, err := u.repo.GetByID(input.ID)
	if err != nil {
		return entities.Patient{}, errors.NewDBError("failed to find patient", err)
	}

	if input.FullName != "" {
		patient.FullName = input.FullName
	}
	if !input.BirthDate.IsZero() {
		patient.BirthDate = input.BirthDate
	}

	updatedPatient, err := u.repo.Update(patient)
	if err != nil {
		return entities.Patient{}, errors.NewDBError("failed to update patient", err)
	}

	return *updatedPatient, nil
}

func (u *PatientUsecase) Delete(id uint) *errors.AppError {
	if err := u.repo.Delete(id); err != nil {
		return errors.NewDBError("failed to delete patient", err)
	}
	return nil
}

*/
