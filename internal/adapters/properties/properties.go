package properties

import (
	"encoding/json"
	"errors"
	"game-inventory-management/internal/adapters/database/connection"
	"io/ioutil"

	"go.uber.org/zap"
)

func Get(log *zap.SugaredLogger) (*Properties, error) {
	file, err := ioutil.ReadFile("properties.json")
	if err != nil {
		log.Error("Error reading file:", err)
		return nil, errors.New("Error reading file")
	}

	var props Properties
	err = json.Unmarshal(file, &props)
	if err != nil {
		log.Error("Error parsing JSON:", err)
		return nil, errors.New("Error parsing JSON")
	}

	return &props, nil
}

type Properties struct {
	Database connection.Connection `json:"database"`
}
