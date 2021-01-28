package configuration

const (
	// ServiceName The name of the current service.
	ServiceName    = "bone-box-service"
	StagingEnv     = "STAGING"
	DevelopmentEnv = "DEVELOPMENT"
)

// CommonConfig ...
type CommonConfig struct {
	Host             string `config:"alias=server_host" validation:"hostname"`
	Port             string `config:"alias=port" validation:"port"`
	Environment      string `config:"alias=service_env"`
	AppEnv           string `config:"alias=app_env"`
	AppRole          string `config:"alias=app_role"`
	AppRoleName      string `config:"alias=app_role_name"`
	EnvironmentConst int
	ECRegion         string
	ECIP             string
	ECID             string
}

type AppConfig struct {
	CommonConfig
	DocsPath 	string `config:"alias=docs_path"`
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
