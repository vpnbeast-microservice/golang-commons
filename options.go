package commons

import (
	"fmt"
	"github.com/spf13/cast"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"net/http"
	"os"
	"reflect"
)

// unmarshalConfig creates a new *viper.Viper and unmarshalls the config into struct using *viper.Viper
func unmarshalConfig(key string, opts interface{}) error {
	sub := viper.Sub(key)
	sub.AutomaticEnv()
	sub.SetEnvPrefix(key)
	// t := reflect.TypeOf(opts)
	bindEnvs(sub, opts)
	return sub.Unmarshal(opts)
}

// bindEnvs takes *viper.Viper as argument and binds structs fields to environments variables to be able to override
// them using environment variables at the runtime
func bindEnvs(sub *viper.Viper, opts interface{}) {
	elem := reflect.ValueOf(opts).Type().Elem()
	fieldCount := elem.NumField()
	for i := 0; i < fieldCount; i++ {
		env := elem.Field(i).Tag.Get("env")
		name := elem.Field(i).Name
		_ = sub.BindEnv(name, env)
	}
}

// getStringEnv gets the specific environment variables with default value, returns default value if variable not set
func getStringEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}

// getIntEnv gets the specific environment variables with default value, returns default value if variable not set
func getIntEnv(key string, defaultValue int) int {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return cast.ToInt(value)
}

// InitOptions gets options interface and appName as parameter and loads the configuration from remote config store
func InitOptions(opts interface{}, name string) error {
	activeProfile := getStringEnv("ACTIVE_PROFILE", "local")
	appName := getStringEnv("APP_NAME", name)
	if activeProfile == "unit-test" {
		logger.Info("active profile is unit-test, reading configuration from static file")
		// TODO: better approach for that?
		viper.AddConfigPath("./../../config")
		viper.SetConfigName("unit_test")
		viper.SetConfigType("yaml")
		if err := viper.ReadInConfig(); err != nil {
			return err
		}
	} else {
		configHost := getStringEnv("CONFIG_SERVER_HOST", "localhost")
		configPort := getIntEnv("CONFIG_SERVER_PORT", 8888)
		logger.Info("loading configuration from remote server", zap.String("host", configHost),
			zap.Int("port", configPort), zap.String("appName", appName),
			zap.String("activeProfile", activeProfile))
		confAddr := fmt.Sprintf("http://%s:%d/%s-%s.yaml", configHost, configPort, appName, activeProfile)
		resp, err := http.Get(confAddr)
		if err != nil {
			return err
		}

		defer func() {
			err := resp.Body.Close()
			if err != nil {
				panic(err)
			}
		}()

		viper.SetConfigName("application")
		viper.SetConfigType("yaml")
		if err = viper.ReadConfig(resp.Body); err != nil {
			return err
		}
	}

	if err := unmarshalConfig(appName, opts); err != nil {
		return err
	}

	return nil
}
