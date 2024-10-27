package view

import (
	"encoding/json"
	"fmt"
	"lumosh_klinik/model"
)

func ShowPatients(patients []model.Patient) {
	data, _ := json.MarshalIndent(patients, "", "  ")
	fmt.Println(string(data))
}
