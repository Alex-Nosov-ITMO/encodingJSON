package main

import (
	"fmt"

	"github.com/Alex-Nosov-ITMO/encodingJSON/encoding"
	"github.com/Alex-Nosov-ITMO/encodingJSON/utils"
)

func Encode(data encoding.MyEncoder) error {
	return data.Encoding()
}

func main() {
	utils.CreateJSONFile()
	utils.CreateYAMLFile()

	jsonData := encoding.JSONData{FileInput: "jsonInput.json", FileOutput: "yamlOutput.yml"}
	err := Encode(&jsonData)
	if err != nil {
		fmt.Printf("ошибка при перекодировании данных из JSON в YAML: %s", err.Error())
	}

	yamlData := encoding.YAMLData{FileInput: "yamlInput.yml", FileOutput: "jsonOutput.json"}
	err = Encode(&yamlData)
	if err != nil {
		fmt.Printf("ошибка при перекодировании данных из YAML в JSON: %s", err.Error())
	}
}
