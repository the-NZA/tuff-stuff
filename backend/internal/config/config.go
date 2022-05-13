package config

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	log "github.com/go-pkgz/lgr"
)

// Config struct includes application configuration.
type Config struct {
	Domain      string `json:"domain"`
	DBPath      string `json:"db_path"`
	SecretKey   string `json:"secret_key"`
	Port        int    `json:"port"`
	IsDebugLogs bool   `json:"is_debug_log"`
	ServeStatic bool   `json:"serve_static"`
}

// NewConfigFromFile reads config from file with given path.
func NewConfigFromFile(path string) (Config, error) {
	var config Config

	file, err := os.Open(path)
	if err != nil {
		return config, err
	}

	defer func(file *os.File) {
		internalError := file.Close()
		if internalError != nil {
			log.Printf("[ERROR] Can't close config file: %v\n", internalError)
		}
	}(file)

	data, err := io.ReadAll(file)
	if err != nil {
		return config, err
	}

	err = json.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}

// GetBindAddr returns concatenated domain and port.
func (c Config) GetBindAddr() string {
	return fmt.Sprintf("%s:%d", c.Domain, c.Port)
}
