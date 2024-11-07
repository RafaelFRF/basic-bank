package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteFloatToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}

func GetFloatFromFile(fileName string) (float64, error) {
	// _ represents a value that I know I could receive but I don't want to work with it
	data, err := os.ReadFile(fileName)
	if err != nil {
		return 0, errors.New("Failed to find file.")
	}

	valueText := string(data)
	value, err := strconv.ParseFloat(valueText, 64)
	if err != nil {
		return 0, errors.New("Failed to parse stored value.")
	}

	return value, nil
}
