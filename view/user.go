package view

import (
	"encoding/json"
	"fmt"
	"lumosh_klinik/model"
)

func ShowUsers(users []model.User) {
	data, _ := json.MarshalIndent(users, "", "  ")
	fmt.Println(string(data))
}
