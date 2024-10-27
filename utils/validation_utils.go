package utils

import "errors"

func ValidateUser(username, password, role string) error {
	if username == "" || password == "" || role == "" {
		return errors.New("semua field harus diisi")
	}
	return nil
}

func ValidatePatient(name string, age int, gender, address string) error {
	if name == "" || age <= 0 || gender == "" || address == "" {
		return errors.New("semua field harus diisi dengan benar")
	}
	return nil
}
