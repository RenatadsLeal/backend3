package domain

type Appointment struct {
	Id          int     `json:"id"`
	Patient     Patient `json:"patient"`
	Dentist     Dentist `json:"dentist"`
	Date        string  `json:"date" binding:"required"`
	Description string  `json:"description" binding:"required"`
}
