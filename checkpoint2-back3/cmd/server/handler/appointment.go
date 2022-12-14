package handler

import (
	"checkpoint2/internal/appointment"
	"checkpoint2/internal/domain"
	"checkpoint2/pkg/web"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type appointmentHandler struct {
	s appointment.Service
}

func NewAppointmentHandler(s appointment.Service) *appointmentHandler {
	return &appointmentHandler{
		s: s,
	}
}

func validateEmptysAppointment(appointment *domain.Appointment) (bool, error) {
	switch {
	case appointment.Date == "" || appointment.Description == "":
		return false, errors.New("date and description can't be empty")
	}
	return true, nil
}

func (h *appointmentHandler) ReadById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, errors.New("invalid id"))
			return
		}
		appointment, err := h.s.ReadById(id)
		if err != nil {
			web.Failure(ctx, http.StatusNotFound, err)
			return
		}
		web.Success(ctx, http.StatusOK, appointment)
	}
}

func (h *appointmentHandler) ReadByRg() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		rg := ctx.Param("rg")
		appointments, err := h.s.ReadByRg(rg)
		if err != nil {
			web.Failure(ctx, http.StatusNotFound, err)
			return
		}
		web.Success(ctx, http.StatusOK, appointments)
	}
}

func (h *appointmentHandler) CreateById() gin.HandlerFunc {
	type Request struct {
		Date        string `json:"date" binding:"required"`
		Description string `json:"description" binding:"required"`
	}
	return func(ctx *gin.Context) {
		var request Request
		idPatient, err := strconv.Atoi(ctx.Param("patient-id"))
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, errors.New("invalid patient id"))
			return
		}
		idDentist, err := strconv.Atoi(ctx.Param("dentist-id"))
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, errors.New("invalid dentist id"))
			return
		}

		err = ctx.ShouldBindJSON(&request)
		if err != nil {
			web.Failure(ctx, http.StatusUnprocessableEntity, errors.New("invalid json"))
			return
		}
		appointment := domain.Appointment{
			Date:        request.Date,
			Description: request.Description,
		}
		valid, err := validateEmptysAppointment(&appointment)
		if !valid {
			web.Failure(ctx, http.StatusUnprocessableEntity, err)
			return
		}

		newAppointment, err := h.s.CreateById(appointment, idPatient, idDentist)
		if err != nil {
			web.Failure(ctx, http.StatusInternalServerError, err)
			return
		}
		web.Success(ctx, http.StatusCreated, newAppointment)
	}
}

func (h *appointmentHandler) CreateByRgAndRegistration() gin.HandlerFunc {
	type Request struct {
		Date        string `json:"date" binding:"required"`
		Description string `json:"description" binding:"required"`
	}
	return func(ctx *gin.Context) {
		var req Request
		rgPatient := ctx.Param("patient-rg")
		registrationDentist := ctx.Param("dentist-registration")

		err := ctx.ShouldBindJSON(&req)
		if err != nil {
			web.Failure(ctx, http.StatusUnprocessableEntity, errors.New("invalid json"))
			return
		}
		appointment := domain.Appointment{
			Date:        req.Date,
			Description: req.Description,
		}
		valid, err := validateEmptysAppointment(&appointment)
		if !valid {
			web.Failure(ctx, http.StatusUnprocessableEntity, err)
			return
		}

		createdAppointment, err := h.s.CreateByRgAndRegistration(appointment, rgPatient, registrationDentist)
		if err != nil {
			web.Failure(ctx, http.StatusInternalServerError, err)
			return
		}
		web.Success(ctx, http.StatusCreated, createdAppointment)
	}
}

func (h *appointmentHandler) Update() gin.HandlerFunc {
	type Request struct {
		PatientId   int    `json:"patient_id" binding:"required"`
		DentistId   int    `json:"dentist_id" binding:"required"`
		Date        string `json:"date" binding:"required"`
		Description string `json:"description" binding:"required"`
	}
	return func(ctx *gin.Context) {
		var req Request
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, errors.New("invalid id"))
			return
		}
		err = ctx.ShouldBindJSON(&req)
		if err != nil {
			web.Failure(ctx, http.StatusUnprocessableEntity, errors.New("invalid json"))
			return
		}
		if req.PatientId == 0 || req.DentistId == 0 || req.Date == "" || req.Description == "" {
			web.Failure(ctx, http.StatusUnprocessableEntity, err)
			return
		}
		updateRequestAppointment := domain.Appointment{
			Patient: domain.Patient{
				Id: req.PatientId,
			},
			Dentist: domain.Dentist{
				Id: req.DentistId,
			},
			Date:        req.Date,
			Description: req.Description,
		}
		updatedAppointment, err := h.s.Update(id, updateRequestAppointment)
		if err != nil {
			web.Failure(ctx, http.StatusInternalServerError, err)
			return
		}
		web.Success(ctx, http.StatusOK, updatedAppointment)
	}
}

func (h *appointmentHandler) Patch() gin.HandlerFunc {
	type Request struct {
		PatientId   int    `json:"patient_id"`
		DentistId   int    `json:"dentist_id"`
		Date        string `json:"date"`
		Description string `json:"description"`
	}
	return func(ctx *gin.Context) {
		var req Request
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, errors.New("invalid id"))
			return
		}
		err = ctx.ShouldBindJSON(&req)
		if err != nil {
			web.Failure(ctx, http.StatusUnprocessableEntity, errors.New("invalid json"))
			return
		}
		appointment := domain.Appointment{
			Patient: domain.Patient{
				Id: req.PatientId,
			},
			Dentist: domain.Dentist{
				Id: req.DentistId,
			},
			Date:        req.Date,
			Description: req.Description,
		}
		updatedAppointment, err := h.s.Patch(id, appointment)
		if err != nil {
			web.Failure(ctx, http.StatusInternalServerError, err)
			return
		}
		web.Success(ctx, http.StatusOK, updatedAppointment)
	}
}

func (h *appointmentHandler) Delete() gin.HandlerFunc {
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
