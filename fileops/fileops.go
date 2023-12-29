package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteFloatToFile(value float64, filename string) {
	valueTxt := fmt.Sprint(value)
	os.WriteFile(filename, []byte(valueTxt), 0644)
}

func GetFloatFromFile(filename string) (float64, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return 1000, errors.New("Failed to find file.")
	}

	value, err := strconv.ParseFloat(string(data), 64)
	if err != nil {
		return 1000, errors.New("Failed to parse stored value.")
	}

	return value, nil
}
