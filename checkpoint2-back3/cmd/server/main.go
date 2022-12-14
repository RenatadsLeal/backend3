package main

import (
	"checkpoint2/cmd/server/handler"
	"checkpoint2/connections"
	"net/http"

	"checkpoint2/internal/appointment"
	"checkpoint2/internal/dentist"
	"checkpoint2/internal/patient"

	"checkpoint2/pkg/store"

	"github.com/gin-gonic/gin"
)

func main() {

	sqlStore := connections.NewSQLStore();

	r := gin.Default()

	sqlStoragePatient := store.NewSQLStorePatient(sqlStore)
	repoPatient := patient.NewRepository(sqlStoragePatient)
	servicePatient := patient.NewService(repoPatient)
	patientHandler := handler.NewPatientHandler(servicePatient)

	r.GET("/ping", func(c *gin.Context) { c.String(http.StatusOK, "pong") })

	patients := r.Group("/patients")
	{
		patients.GET("/id/:id", patientHandler.ReadById())
		patients.GET("/rg/:rg", patientHandler.ReadByRg())
		patients.POST("", patientHandler.Create())
		patients.PUT(":id", patientHandler.Update())
		patients.PATCH(":id", patientHandler.Patch())
		patients.DELETE(":id", patientHandler.Delete())
	}

	sqlStorageDentist := store.NewSQLStoreDentist(sqlStore)
	repoDentist := dentist.NewRepository(sqlStorageDentist)
	serviceDentist := dentist.NewService(repoDentist)
	dentistHandler := handler.NewDentistHandler(serviceDentist)

	dentists := r.Group("/dentists")
	{
		dentists.GET("/id/:id", dentistHandler.ReadById())
		dentists.GET("/registration/:registration", dentistHandler.ReadByRegistration())
		dentists.POST("", dentistHandler.Create())
		dentists.PUT(":id", dentistHandler.Update())
		dentists.PATCH(":id", dentistHandler.Patch())
		dentists.DELETE(":id", dentistHandler.Delete())
	}

	sqlStorageAppointment := store.NewSQLStoreAppointment(sqlStore)
	repoAppointment := appointment.NewRepository(sqlStorageAppointment)
	serviceAppointment := appointment.NewService(repoAppointment)
	appointmentHandler := handler.NewAppointmentHandler(serviceAppointment)

	appointments := r.Group("/appointments")
	{
		appointments.GET("/id/:id", appointmentHandler.ReadById())
		appointments.GET("/rg/:rg", appointmentHandler.ReadByRg())
		appointments.POST("/id/:patient-id/:dentist-id", appointmentHandler.CreateById())
		appointments.POST("/rg-registration/:patient-rg/:dentist-registration", appointmentHandler.CreateByRgAndRegistration())
		appointments.PUT(":id", appointmentHandler.Update())
		appointments.PATCH(":id", appointmentHandler.Patch())
		appointments.DELETE(":id", appointmentHandler.Delete())
	}

	r.Run(":8080")
}