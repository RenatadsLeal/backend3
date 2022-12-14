package handler

import (
	"checkpoint2/internal/dentist"
	"checkpoint2/internal/domain"
	"checkpoint2/pkg/web"
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type dentistHandler struct {
	s dentist.Service
}

func NewDentistHandler(s dentist.Service) *dentistHandler {
	return &dentistHandler{
		s: s,
	}
}

func validateEmptysDentist(dentist *domain.Dentist) (bool, error) {
	switch {
	case dentist.Surname == "" || dentist.Name == "" || dentist.Registration == "":
		return false, errors.New("fields can't be empty")
	}
	return true, nil
}

func (h *dentistHandler) ReadById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, errors.New("invalid ID"))
			return
		}

		dentist, err := h.s.ReadById(id)
		if err != nil {
			web.Failure(ctx, http.StatusNotFound, err)
			return
		}

		web.Success(ctx, http.StatusOK, dentist)
	}
}

func (h *dentistHandler) ReadByRegistration() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		dentist, err := h.s.ReadByRegistration(ctx.Param("registration"))
		if err != nil {
			web.Failure(ctx, http.StatusNotFound, err)
			return
		}

		web.Success(ctx, http.StatusOK, dentist)
	}
}

func (h *dentistHandler) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dentist domain.Dentist
		err := ctx.ShouldBindJSON(&dentist)
		if err != nil {
			web.Failure(ctx, http.StatusUnprocessableEntity, errors.New("invalid json"))
			return
		}
		valid, err := validateEmptysDentist(&dentist)
		if !valid {
			web.Failure(ctx, http.StatusUnprocessableEntity, err)
			return
		}
		createdDentist, err := h.s.Create(dentist)
		if err != nil {
			web.Failure(ctx, http.StatusInternalServerError, err)
			return
		}
		web.Success(ctx, http.StatusCreated, createdDentist)
	}
}

func (h *dentistHandler) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, errors.New("invalid ID"))
			return
		}

		var dentist domain.Dentist

		err = ctx.ShouldBindJSON(&dentist)
		if err != nil {
			web.Failure(ctx, http.StatusUnprocessableEntity, errors.New("invalid json"))
			return
		}

		valid, err := validateEmptysDentist(&dentist)
		if !valid {
			web.Failure(ctx, http.StatusUnprocessableEntity, err)
			return
		}

		createdDentist, err := h.s.Update(id, dentist)
		if err != nil {
			web.Failure(ctx, http.StatusInternalServerError, err)
			return
		}

		web.Success(ctx, http.StatusCreated, createdDentist)
	}
}

func (h *dentistHandler) Patch() gin.HandlerFunc {
	type Request struct {
		Surname      string `json:"surname,omitempty"`
		Name         string `json:"name,omitempty"`
		Registration string `json:"registration,omitempty"`
	}
	return func(ctx *gin.Context) {
		var req Request
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			web.Failure(ctx, http.StatusBadRequest, errors.New("invalid id"))
			return
		}
		if err := ctx.ShouldBindJSON(&req); err != nil {
			web.Failure(ctx, http.StatusUnprocessableEntity, errors.New("invalid json"))
			return
		}

		update := domain.Dentist{
			Surname:      req.Surname,
			Name:         req.Name,
			Registration: req.Registration,
		}

		updatedDentist, err := h.s.Patch(id, update)
		if err != nil {
			web.Failure(ctx, http.StatusInternalServerError, err)
			return
		}

		web.Success(ctx, http.StatusOK, updatedDentist)
	}
}

func (h *dentistHandler) Delete() gin.HandlerFunc {
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
