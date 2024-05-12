package configs

import (
	"encoding/json"
	"os"
	"swift/pkg/models"
)
var Settings *models.ConfigModel

func InitConfigs(path string) (err error) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	//нужен для того чтоб декодировать или десериализация bytes в json и переобразует его структуру config
	err = json.Unmarshal(bytes, &Settings)
	if err != nil {
		return err
	}

	return err
}
