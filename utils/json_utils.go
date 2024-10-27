package utils

import "encoding/json"

func MarshalJSON(data interface{}) (string, error) {
	marshaledData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return "", err
	}
	return string(marshaledData), nil
}
