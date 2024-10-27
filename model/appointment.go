package model

import "time"

type Appointment struct {
	ID        int       `json:"id"`
	PatientID int       `json:"patient_id"`
	DoctorID  int       `json:"doctor_id"`
	Date      time.Time `json:"date"`
	Notes     string    `json:"notes"`
}
