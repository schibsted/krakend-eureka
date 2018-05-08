//
// Copyright 2011 - 2018 Schibsted Products & Technology AS.
// Licensed under the terms of the Apache 2.0 license. See LICENSE in the project root.
//

package eureka

import (
	"testing"

	"github.com/devopsfaith/krakend/config"
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
