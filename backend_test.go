package eureka

import (
	"github.com/devopsfaith/krakend/config"
	"testing"
)

func TestBackendExtraConfig(t *testing.T) {
	cfg := &config.Backend{
		ExtraConfig: config.ExtraConfig{
			BackendNamespace: BackendExtraConfig{EurekaAppName: "testEureka"}.AsMap(),
		},
	}
	res := GetBackendExtraConfig(cfg)
	if v := res.EurekaAppName; v != "testEureka" {
		t.Errorf("unexpected value for key `EurekaAppName`: %v", v)
		return
	}
}

func TestDefaultBackendExtraConfig(t *testing.T) {
	cfg := &config.Backend{}
	res := GetBackendExtraConfig(cfg)
	if v := res.EurekaAppName; v != "" {
		t.Errorf("unexpected value for key `EurekaAppName`: %v", v)
		return
	}
}
