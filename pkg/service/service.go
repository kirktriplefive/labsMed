package service

import (
	"github.com/kirktriplefive/labsMed"
	"github.com/kirktriplefive/labsMed/pkg/repository"
)

type Polyclinic interface {
	CreatePatient(patient labsMed.Patient) (int, error)
	CreatePatientRecord(record labsMed.Record) (int, error)
	CreateDoctor(doctor labsMed.Doctor) (int, error)
	GetDoctors() ([]labsMed.Doctor, error)
	GetRecords() ([]labsMed.Record, error)
	GetRecordOfPatient(int) ([]labsMed.Record, error)
}

type Service struct {
	Polyclinic
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Polyclinic: &PoliclinicService{*repos},
	}
}
