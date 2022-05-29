package service

import (
	"github.com/kirktriplefive/labsMed"
	"github.com/kirktriplefive/labsMed/pkg/repository"
)

type PoliclinicService struct {
	repo repository.Polyclinic
}

func NewOrderService(repo repository.Repository) *PoliclinicService {
	return &PoliclinicService{
		repo: repo,
	}
}

func (s *PoliclinicService) CreatePatient(patient labsMed.Patient) (int, error) {
	e, err:= s.repo.CreatePatient(patient)
	return e,err
}

func (s *PoliclinicService) CreateDoctor(doctor labsMed.Doctor) (int, error) {
	e, err := s.repo.CreateDoctor(doctor)
	return e, err
}

func (s *PoliclinicService) CreatePatientRecord(record labsMed.Record) (int, error) {
	e, err:= s.repo.CreatePatientRecord(record)
	return e, err
}

func (s *PoliclinicService) GetDoctors() ([]labsMed.Doctor, error) {
	e, err:= s.repo.GetDoctors()
	return e, err
}

func (s *PoliclinicService) GetRecords() ([]labsMed.Record, error) {
	e, err:=s.repo.GetRecords()
	return e, err
}

func (s *PoliclinicService) GetRecordOfPatient(id int) ([]labsMed.Record, error) {
	e, err:=s.repo.GetRecordOfPatient(id)
	return e, err
}
