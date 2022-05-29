package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/kirktriplefive/labsMed"
)

type Polyclinic interface {
	CreatePatient(patient labsMed.Patient) (int, error)
	CreatePatientRecord(record labsMed.Record) (int, error)
	CreateDoctor(doctor labsMed.Doctor) (int, error)
	GetDoctors() ([]labsMed.Doctor, error)
	GetRecords() ([]labsMed.Record, error)
	GetRecordOfPatient(int) ([]labsMed.Record, error)
}

type Repository struct {
	Polyclinic
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Polyclinic: NewPoliclinicPostgres(db),
	}
}
