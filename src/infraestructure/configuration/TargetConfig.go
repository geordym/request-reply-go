package configuration

import (
	"encoding/json"
	"fmt"
	"os"
)

type Target struct {
	TARGET_KEY         string                 `json:"TARGET_KEY"`
	TARGET_TYPE        string                 `json:"TARGET_TYPE"`
	TARGET_CONFIG      map[string]interface{} `json:"TARGET_CONFIG"`
	MESSAGE_TEMPLATE   string                 `json:"MESSAGE_TEMPLATE"`
	MESSAGE_SERIALIZER string
}

type Config struct {
	Targets []Target `json:"targets"`
}

var config Config

func InitializeTargets(filePath string) error {
	fmt.Printf("Entrando a la función InitializeTargets con el archivo: %s\n", filePath)

	_, err := os.Stat(filePath)
	if os.IsNotExist(err) {
		return fmt.Errorf("El archivo %s no existe", filePath)
	}
	if err != nil {
		return fmt.Errorf("Error al verificar el archivo %s: %v", filePath, err)
	}

	file, err := os.ReadFile(filePath)
	if err != nil {
		return fmt.Errorf("Error leyendo el archivo JSON: %v", err)
	}

	fmt.Printf("Contenido del archivo leído: \n%s\n", string(file))

	err = json.Unmarshal(file, &config)
	if err != nil {
		return fmt.Errorf("Error deserializando el JSON: %v", err)
	}

	fmt.Printf("Se cargaron %d targets\n", len(config.Targets))

	return nil
}

func FindTargetByKey(targetKey string) (*Target, error) {
	for _, target := range config.Targets {
		if target.TARGET_KEY == targetKey {
			return &target, nil
		}
	}
	return nil, fmt.Errorf("Target con clave %s no encontrado", targetKey)
}
