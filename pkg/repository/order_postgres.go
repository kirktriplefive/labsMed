package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/kirktriplefive/labsMed"
)

type PoliclinicPostgres struct {
	db *sqlx.DB
}

func NewPoliclinicPostgres(db *sqlx.DB) *PoliclinicPostgres {
	return &PoliclinicPostgres{db: db}
}

func (r *PoliclinicPostgres) CreatePatient(patient labsMed.Patient) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	var patientId int
	createPatientQuery := fmt.Sprintf("INSERT INTO %s (name, second_name) VALUES ($1, $2) RETURNING p_id", pacientTable)
	row := tx.QueryRow(createPatientQuery, patient.Name, patient.SecondName)
	if err := row.Scan(&patientId); err != nil {
		tx.Rollback()
		return 0, err
	}
	return patientId, tx.Commit()

}

func (r *PoliclinicPostgres) CreatePatientRecord(record labsMed.Record) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	var recordId int
	createPatientQuery := fmt.Sprintf("INSERT INTO %s (date, pacient_id, doctor_id, diagnosis) VALUES ($1, $2, $3, $4) RETURNING r_id", recordsTable)
	row := tx.QueryRow(createPatientQuery, record.Date, record.PacientId, record.DoctorId, record.Diagnosis)
	if err := row.Scan(&recordId); err != nil {
		tx.Rollback()
		return 0, err
	}
	return recordId, tx.Commit()
}

func (r *PoliclinicPostgres) CreateDoctor(doctor labsMed.Doctor) (int, error) {
	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}
	var doctorId int
	createPatientQuery := fmt.Sprintf("INSERT INTO %s (name, second_name, middle_name, specialization) VALUES ($1, $2, $3, $4) RETURNING d_id", doctorsTable)
	row := tx.QueryRow(createPatientQuery, doctor.Name, doctor.SecondName, doctor.MiddleName, doctor.Specialization)
	if err := row.Scan(&doctorId); err != nil {
		tx.Rollback()
		return 0, err
	}
	return doctorId, tx.Commit()
}

func (r *PoliclinicPostgres) GetDoctors() ([]labsMed.Doctor, error) {
	var doctor []labsMed.Doctor
	doctorQuery := fmt.Sprintf("SELECT d_id, name, second_name, middle_name, specialization FROM %s",
		doctorsTable)
	if err := r.db.Select(&doctor, doctorQuery); err != nil {
		return doctor, err
	}

	return doctor, nil
}

func (r *PoliclinicPostgres) GetRecords() ([]labsMed.Record, error) {
	var record []labsMed.Record
	recordQuery := fmt.Sprintf("SELECT r_id, date, doctor_id, pacient_id, diagnosis FROM %s",
		recordsTable)
	if err := r.db.Select(&record, recordQuery); err != nil {
		return record, err
	}

	return record, nil
}

func (r *PoliclinicPostgres) GetRecordOfPatient(id int) ([]labsMed.Record, error) {
	var record []labsMed.Record
	recordQuery := fmt.Sprintf("SELECT td.r_id, td.date, td.doctor_id, td.pacient_id, td.diagnosis FROM %s td INNER JOIN %s pt ON pt.p_id=td.pacient_id WHERE pt.p_id = $1",
		recordsTable, pacientTable)
	if err := r.db.Select(&record, recordQuery, id); err != nil {
		return record, err
	}

	return record, nil
}