package utils


import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

// ReadJSONFile reads a JSON file and returns the data as a map
func ReadJSONFile(filename string) (map[string]interface{}, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	byteValue, err := ioutil.ReadAll(file)
	if err != nil {
		return nil, err
	}
	var result map[string]interface{}
	json.Unmarshal([]byte(byteValue), &result)
	return result, nil
}


// WriteJSONFile writes a JSON file from a map
func WriteJSONFile(filename string, data map[string]interface{}) error {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filename, file, 0644)
	if err != nil {
		return err
	}
	return nil
}

// PrintJSON prints a JSON file from a map
func PrintJSON(data map[string]interface{}) error {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		return err
	}
	fmt.Println(string(file))
	return nil
}

// MapToStruct converts a map to a struct
func MapToStruct(m map[string]interface{}, s interface{}) error {
	b, err := json.Marshal(m)
	if err != nil {
		return err
	}
	err = json.Unmarshal(b, s)
	if err != nil {
		return err
	}
	return nil
}

// StructToMap converts a struct to a map
func StructToMap(s interface{}) map[string]interface{} {
	b, err := json.Marshal(s)
	if err != nil {
		return nil
	}
	var result map[string]interface{}
	json.Unmarshal([]byte(b), &result)
	return result
}
