package handler

import (
	"net/http"

	"github.com/kirktriplefive/labsMed/pkg/service"
)

type HandlerInterface interface {
	CreatePatient(w http.ResponseWriter, r *http.Request)
	CreatePatientRecord(w http.ResponseWriter, r *http.Request)
	CreateDoctor(w http.ResponseWriter, r *http.Request)
	GetDoctors(w http.ResponseWriter, r *http.Request)
	GetRecords(w http.ResponseWriter, r *http.Request)
	GetRecordOfPatient(w http.ResponseWriter, r *http.Request)
}

type Handler struct {
	HandlerInterface
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{&PoliclinicHandler{*service}}
}
