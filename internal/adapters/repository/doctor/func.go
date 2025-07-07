package doctor

import "github.com/AlexanderMorozov1919/mobileapp/internal/domain/entities"

func (r *DoctorRepository) CreateDoctor(doctor *entities.Doctor) error {
	return r.db.Create(doctor).Error
}

func (r *DoctorRepository) UpdateDoctor(doctor *entities.Doctor) error {
	return r.db.Save(doctor).Error
}

func (r *DoctorRepository) DeleteDoctor(id uint) error {
	return r.db.Delete(&entities.Doctor{}, id).Error
}

func (r *DoctorRepository) GetDoctorByID(id uint) (*entities.Doctor, error) {
	var doctor entities.Doctor
	if err := r.db.First(&doctor, id).Error; err != nil {
		return nil, err
	}
	return &doctor, nil
}

func (r *DoctorRepository) GetDoctorByLogin(login string) (*entities.Doctor, error) {
	var doctor entities.Doctor
	if err := r.db.Where("login = ?", login).First(&doctor).Error; err != nil {
		return nil, err
	}
	return &doctor, nil
}

func (r *DoctorRepository) GetDoctorName(id uint) (string, error) {
	var doctor entities.Doctor
	if err := r.db.Select("full_name").First(&doctor, id).Error; err != nil {
		return "", err
	}
	return doctor.FullName, nil
}

func (r *DoctorRepository) GetDoctorSpecialization(id uint) (string, error) {
	var doctor entities.Doctor
	if err := r.db.Select("specialization").First(&doctor, id).Error; err != nil {
		return "", err
	}
	return doctor.Specialization, nil
}

func (r *DoctorRepository) GetDoctorPassHash(id uint) (string, error) {
	var doctor entities.Doctor
	if err := r.db.Select("password_hash").First(&doctor, id).Error; err != nil {
		return "", err
	}
	return doctor.PasswordHash, nil
}
