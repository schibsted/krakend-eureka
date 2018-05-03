package eureka

import (
	"github.com/hudl/fargo"
	"testing"
)

func TestGetHosts_ok(t *testing.T) {
	app := &fargo.Application{
		Name: "APPNAME",
		Instances: []*fargo.Instance{
			{
				HostName: "host1",
				Port:     111,
				Status:   fargo.UP,
			},
			{
				HostName: "hosts21",
				Port:     222,
				Status:   fargo.STARTING,
			},
			{
				HostName: "host2",
				Port:     333,
				Status:   fargo.UP,
			},
		},
	}

	hosts, err := hosts(app)
	if err != nil {
		t.Errorf("Unexpected error parsing app data %v", err)
	}
	if hosts == nil {
		t.Errorf("Hosts slice should be not nil")
	}
	if len(hosts) != 2 || hosts[0] != "http://host1:111" || hosts[1] != "http://host2:333" {
		t.Errorf("Unexpected Hosts list. Got %v", hosts)
	}
}

func TestGetHosts_ko(t *testing.T) {
	hosts, err := hosts(nil)
	if hosts != nil {
		t.Errorf("Hosts slice should be nil. Got: %v`", hosts)
	}
	if err == nil {
		t.Errorf("We should have got an error: Failed to get application metadata from eureka")
	}
}
