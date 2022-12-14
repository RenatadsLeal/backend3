package handler

import (
	"checkpoint2/internal/domain"
	"checkpoint2/internal/patient"
	"checkpoint2/pkg/web"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type patientHandler struct {
	s patient.Service
}

func NewPatientHandler(s patient.Service) *patientHandler {
	return &patientHandler{
		s: s,
	}
}

func validateEmptysPatient(patient *domain.Patient) (bool, error) {
	switch {
	case patient.Surname == "" || patient.Name == "" || patient.RG == "" || patient.RegistrationDate == "":
		return false, errors.New("fields can't be empty")
	}
	return true, nil
}

func (h *patientHandler) ReadById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, errors.New("invalid id"))
			return
		}
		patient, err := h.s.ReadById(id)
		if err != nil {
			web.Failure(ctx, http.StatusNotFound, err)
			return
		}
		web.Success(ctx, http.StatusOK, patient)
	}
}

func (h *patientHandler) ReadByRg() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		patient, err := h.s.ReadByRg(ctx.Param("rg"))
		if err != nil {
			web.Failure(ctx, http.StatusNotFound, err)
			return
		}
		web.Success(ctx, http.StatusOK, patient)
	}
}

func (h *patientHandler) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var patient domain.Patient
		err := ctx.ShouldBindJSON(&patient)
		if err != nil {
			web.Failure(ctx, http.StatusUnprocessableEntity, errors.New("invalid json"))
			return
		}
		valid, err := validateEmptysPatient(&patient)
		if !valid {
			web.Failure(ctx, http.StatusUnprocessableEntity, err)
			return
		}
		createdPatient, err := h.s.Create(patient)
		if err != nil {
			web.Failure(ctx, http.StatusInternalServerError, err)
			return
		}
		web.Success(ctx, http.StatusCreated, createdPatient)
	}
}

func (h *patientHandler) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, errors.New("invalid id"))
			return
		}
		var patient domain.Patient
		err = ctx.ShouldBindJSON(&patient)
		if err != nil {
			web.Failure(ctx, http.StatusUnprocessableEntity, errors.New("invalid json"))
			return
		}
		valid, err := validateEmptysPatient(&patient)
		if !valid {
			web.Failure(ctx, http.StatusUnprocessableEntity, err)
			return
		}
		createdPatient, err := h.s.Update(id, patient)
		if err != nil {
			web.Failure(ctx, http.StatusInternalServerError, err)
			return
		}
		web.Success(ctx, http.StatusCreated, createdPatient)
	}
}

func (h *patientHandler) Patch() gin.HandlerFunc {
	type Request struct {
		Surname          string `json:"surname,omitempty"`
		Name             string `json:"name,omitempty"`
		Rg               string `json:"rg,omitempty"`
		RegistrationDate string `json:"registration_date,omitempty"`
	}
	return func(ctx *gin.Context) {
		var request Request
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, errors.New("invalid id"))
			return
		}
		if err := ctx.ShouldBindJSON(&request); err != nil {
			web.Failure(ctx, http.StatusUnprocessableEntity, errors.New("invalid json"))
			return
		}

		update := domain.Patient{
			Surname:          request.Surname,
			Name:             request.Name,
			RG:               request.Rg,
			RegistrationDate: request.RegistrationDate,
		}

		updatedPatient, err := h.s.Patch(id, update)
		if err != nil {
			web.Failure(ctx, http.StatusInternalServerError, err)
			return
		}

		web.Success(ctx, http.StatusOK, updatedPatient)
	}
}

func (h *patientHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, errors.New("invalid id"))
			return
		}
		err = h.s.Delete(id)
		if err != nil {
			web.Failure(ctx, http.StatusNotFound, err)
			return
		}

		web.Success(ctx, http.StatusNoContent, nil)
	}
}
