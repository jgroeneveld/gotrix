package cfg

import (
	"encoding/json"
	"os"

	"log"

	"gotrix/lib/errors"
)

type config struct {
	ApplicationName string
	DatabaseURL     string
	Port            string
}

var Config = loadConfig()

// loadConfig loads the configuration based on ENV, ConfigFile and Defaults. In that order.
func loadConfig() *config {
	c := &config{}
	defaults := configDefaults[Env()]
	configFile, err := loadConfigFile(os.ExpandEnv("$HOME/.gotrix/config.json"))
	if err != nil {
		log.Fatal(err)
	}

	c.ApplicationName = configValue("APPLICATION_NAME", configFile.ApplicationName, defaults.ApplicationName)
	c.DatabaseURL = configValue("DATABASE_URL", configFile.DatabaseURL, defaults.DatabaseURL)
	c.Port = configValue("PORT", configFile.Port, defaults.Port)

	return c
}

func configValue(envKey string, fallbacks ...string) string {
	env := os.Getenv(envKey)
	if env != "" {
		return env
	}
	for _, fb := range fallbacks {
		if fb != "" {
			return fb
		}
	}
	return ""
}

func loadConfigFile(cfgPath string) (*config, error) {
	f, err := os.Open(cfgPath)

	if os.IsNotExist(err) {
		return &config{}, nil
	} else if err != nil {
		return nil, errors.New("failed to open config file %q: %s", cfgPath, err)
	}
	defer f.Close()

	var c *config
	if err := json.NewDecoder(f).Decode(&c); err != nil {
		return nil, errors.New("failed to parse config file %q: %s", cfgPath, err)
	}

	return c, nil
}
