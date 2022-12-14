package domain

type Patient struct {
	Id               int    `json:"id"`
	Surname          string `json:"surname" binding:"required"`
	Name             string `json:"name" binding:"required"`
	RG               string `json:"rg" binding:"required"`
	RegistrationDate string `json:"registration_date" binding:"required"`
}
