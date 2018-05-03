package eureka

import (
	"github.com/devopsfaith/krakend/config"
)

// Namespace is the key to use to store and access the custom config data
const BackendNamespace = "spt/gateway/backend/eureka"

type BackendExtraConfig struct {
	EurekaAppName string
}

var DefaultBackendExtraConfig = BackendExtraConfig{EurekaAppName: ""}

var EmptyBackendExtraConfig = BackendExtraConfig{}

func (e BackendExtraConfig) AsMap() map[interface{}]interface{} {
	data := make(map[interface{}]interface{})
	data[BACKEND_EUREKA_APP_NAME] = e.EurekaAppName
	return data
}

func GetBackendExtraConfig(e *config.Backend) BackendExtraConfig {
	return BackendConfigGetter(e.ExtraConfig).(BackendExtraConfig)
}

const BACKEND_EUREKA_APP_NAME = "eureka_app_name"

func BackendConfigGetter(e config.ExtraConfig) interface{} {
	v, ok := e[BackendNamespace]
	if !ok {
		return EmptyBackendExtraConfig
	}
	extra, ok := v.(map[interface{}]interface{})
	if !ok {
		return DefaultBackendExtraConfig
	}

	return NewBackendExtraConfigFromMap(extra)
}

func NewBackendExtraConfigFromMap(data map[interface{}]interface{}) BackendExtraConfig {
	eurekaAppName := DefaultBackendExtraConfig.EurekaAppName
	if data != nil {
		if eureka, ok := data[BACKEND_EUREKA_APP_NAME]; ok {
			eurekaAppName = eureka.(string)
		}
	}
	return BackendExtraConfig{eurekaAppName}
}
