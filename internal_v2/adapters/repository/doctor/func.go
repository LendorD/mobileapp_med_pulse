package Allergy

import "github.com/AlexanderMorozov1919/mobileapp/internal_v2/domain/entities"

func (r *DoctorRepositor) Create(doctor *entities.Doctor) error {
	return r.db.Create(doctor).Error
}

func (r *DoctorRepositor) Update(doctor *entities.Doctor) error {
	return r.db.Save(doctor).Error
}

func (r *DoctorRepositor) Delete(id uint) error {
	return r.db.Delete(&entities.Doctor{}, id).Error
}

func (r *DoctorRepositor) GetByID(id uint) (*entities.Doctor, error) {
	var doctor entities.Doctor
	if err := r.db.First(&doctor, id).Error; err != nil {
		return nil, err
	}
	return &doctor, nil
}

func (r *DoctorRepositor) GetByLogin(login string) (*entities.Doctor, error) {
	var doctor entities.Doctor
	if err := r.db.Where("login = ?", login).First(&doctor).Error; err != nil {
		return nil, err
	}
	return &doctor, nil
}

func (r *DoctorRepositor) GetName(id uint) (string, error) {
	var doctor entities.Doctor
	if err := r.db.Select("full_name").First(&doctor, id).Error; err != nil {
		return "", err
	}
	return doctor.FullName, nil
}

func (r *DoctorRepositor) GetSpecialization(id uint) (string, error) {
	var doctor entities.Doctor
	if err := r.db.Select("specialization").First(&doctor, id).Error; err != nil {
		return "", err
	}
	return doctor.Specialization, nil
}

func (r *DoctorRepositor) GetPassHash(id uint) (string, error) {
	var doctor entities.Doctor
	if err := r.db.Select("password_hash").First(&doctor, id).Error; err != nil {
		return "", err
	}
	return doctor.PasswordHash, nil
}
