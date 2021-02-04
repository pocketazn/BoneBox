package configuration

import (
	"encoding/json"
	"io/ioutil"
)

const (
	// ServiceName The name of the current service.
	ServiceName    = "bone-box-service"
	StagingEnv     = "STAGING"
	DevelopmentEnv = "DEVELOPMENT"
)

// CommonConfig ...
type CommonConfig struct {
	Host        string `json:"server_host"`
	Port        string `json:"port"`
	Environment string `json:"service_env"`
	AppEnv      string `json:"app_env"`
	AppRole     string `json:"app_role"`
	AppRoleName string `json:"app_role_name"`

	User     string `json:"postgres_user"`
	Password string `json:"postgres_password"`
	DBName   string `json:"postgres_db_name"`
}

type AppConfig struct {
	CommonConfig
	DocsPath string `json:"docs_path"`
}

var globalConfig AppConfig

func Configuration() *AppConfig {
	return &globalConfig
}

// Loads the AppConfig.
func Configure() (*AppConfig, error) {
	if globalConfig != (AppConfig{}) {
		return &globalConfig, nil
	}

	c := AppConfig{}

	file, _ := ioutil.ReadFile("boneboxConfig.json")

	_ = json.Unmarshal(file, &c)

	err := Load(ServiceName, &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func Load(serviceName string, config *AppConfig) error {
	// Change setting based off ENV

	return nil
}
