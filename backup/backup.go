package backup

import (
	"encoding/json"
	"errors"
	"os"
)

func GetJSONFromData(data map[string]map[string]interface{}) ([]byte, error) {
	j, err := json.MarshalIndent(data, "", "  ")
	if string(j) == "null" {
		return nil, errors.New("Invalid data")
	}
	return j, err
}

func SaveJSONToFile(filename string, data []byte) error {
	err := os.WriteFile(filename, data, 0644)
	return err
}
