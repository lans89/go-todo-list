package conf

import (
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

type DBConf struct {
	Username           string `yaml:"username"`
	Password           string `yaml:"password"`
	Server             string `yaml:"server"`
	Port               int    `yaml:"port"`
	DataBaseName       string `yaml:"databasename"`
	MaxOpenConnections int    `yaml:"maxOpenConnections"`
	MaxIdleConnections int    `yaml:"maxIdleConnections"`
	MaxLifetimeMin     int    `yaml:"maxLifetimeMin"`
	MaxIdleTimeMin     int    `yaml:"maxIdleTimeMin"`
}

type Config struct {
	DB DBConf `yaml:"db"`
}

var config *Config

func GetConfig() *Config {
	if config == nil {
		// Leer el archivo YAML
		data, err := os.ReadFile("conf/conf.yaml")
		if err != nil {
			log.Fatalf("Error al leer el archivo: %v", err)
		}

		// Deserializar el contenido YAML en la estructura
		err = yaml.Unmarshal(data, &config)
		if err != nil {
			log.Fatalf("Error al parsear el YAML: %v", err)
		}
	}
	return config
}
