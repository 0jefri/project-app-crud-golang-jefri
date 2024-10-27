package service

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sync"

	"lumosh_klinik/model"
)

var (
	appointments []model.Appointment
	appMux       sync.Mutex
	appFile      = "appointments.json"
)

func LoadAppointments() error {
	data, err := ioutil.ReadFile(appFile)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	return json.Unmarshal(data, &appointments)
}

func SaveAppointments() error {
	data, err := json.MarshalIndent(appointments, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(appFile, data, 0644)
}

func AddAppointment(appointment model.Appointment) error {
	appMux.Lock()
	defer appMux.Unlock()
	appointments = append(appointments, appointment)
	return SaveAppointments()
}

func GetAppointments() []model.Appointment {
	return appointments
}
