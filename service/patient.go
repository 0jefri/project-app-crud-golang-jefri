package service

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"sync"

	"lumosh_klinik/model"
)

var (
	patients    []model.Patient
	patientMux  sync.Mutex
	patientFile = "patients.json"
)

func LoadPatients() error {
	data, err := ioutil.ReadFile(patientFile)
	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}
	return json.Unmarshal(data, &patients)
}

func SavePatients() error {
	data, err := json.MarshalIndent(patients, "", "  ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(patientFile, data, 0644)
}

func AddPatient(patient model.Patient) error {
	patientMux.Lock()
	defer patientMux.Unlock()
	patients = append(patients, patient)
	return SavePatients()
}

func GetPatients() []model.Patient {
	return patients
}
