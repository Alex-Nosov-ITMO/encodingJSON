package encoding

import (
	"encoding/json"
	"os"

	"github.com/Alex-Nosov-ITMO/encodingJSON/models"
	"gopkg.in/yaml.v3"

	"fmt"
)

// JSONData тип для перекодирования из JSON в YAML
type JSONData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// YAMLData тип для перекодирования из YAML в JSON
type YAMLData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// MyEncoder интерфейс для структур YAMLData и JSONData
type MyEncoder interface {
	Encoding() error
}

// Encoding перекодирует файл из JSON в YAML
func (j *JSONData) Encoding() error {
	y, err := os.ReadFile(j.FileInput)
	if err != nil {
		fmt.Printf("ошибка при чтении файла: %s\n", err.Error())
		return err
	}
	err = json.Unmarshal(y, &j.DockerCompose)
	if err != nil {
		fmt.Printf("ошибка при декодировании объекта: %s\n", err.Error())
		return err
	}
	yData, err := yaml.Marshal(&j.DockerCompose)
	if err != nil {
		fmt.Printf("ошибка при кодировании объекта: %s\n", err.Error())
		return err
	}
	yamlFile, err := os.Create(j.FileOutput)
	if err != nil {
		fmt.Printf("ошибка при создании файла: %s\n", err.Error())
		return err
	}
	defer yamlFile.Close()
	_, err = yamlFile.Write(yData)
	if err != nil {
		fmt.Printf("ошибка при записи в файл: %s\n", err.Error())
		return err
	}
	return nil
}

// Encoding перекодирует файл из YAML в JSON
func (y *YAMLData) Encoding() error {
	j, err := os.ReadFile(y.FileInput)
	if err != nil {
		fmt.Printf("ошибка при чтении файла: %s\n", err.Error())
		return err
	}
	err = yaml.Unmarshal(j, &y.DockerCompose)
	if err != nil {
		fmt.Printf("ошибка при декодировании объекта: %s\n", err.Error())
		return err
	}
	jData, err := json.Marshal(&y.DockerCompose)
	if err != nil {
		fmt.Printf("ошибка при кодировании объекта: %s\n", err.Error())
		return err
	}
	jsonFile, err := os.Create(y.FileOutput)
	if err != nil {
		fmt.Printf("ошибка при создании файла: %s\n", err.Error())
		return err
	}
	defer jsonFile.Close()
	_, err = jsonFile.Write(jData)
	if err != nil {
		fmt.Printf("ошибка при записи в файл: %s\n", err.Error())
		return err
	}
	return nil
}
