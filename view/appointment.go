package view

import (
	"encoding/json"
	"fmt"
	"lumosh_klinik/model"
)

func ShowAppointments(appointments []model.Appointment) {
	data, _ := json.MarshalIndent(appointments, "", "  ")
	fmt.Println(string(data))
}
