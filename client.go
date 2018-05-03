package eureka

import (
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/sd/eureka"
	"github.com/hudl/fargo"

	"errors"
	"fmt"
	"github.com/devopsfaith/krakend/config"
	"github.com/devopsfaith/krakend/logging"
	"github.com/devopsfaith/krakend/sd"

	"os"
	"time"
)

const POLL_INTERVAL = 30 // seconds

type EurekaClient interface {
	EnableInstance()
	NewSubscriber(subscriber sd.SubscriberFactory) sd.SubscriberFactory
}

type fargoEurekaClient struct {
	conn        *fargo.EurekaConnection
	appInstance *fargo.Instance
}

func (e *fargoEurekaClient) NewSubscriber(subscriber sd.SubscriberFactory) sd.SubscriberFactory {
	return func(config *config.Backend) sd.Subscriber {
		backendExtraConfig := GetBackendExtraConfig(config)
		if backendExtraConfig == EmptyBackendExtraConfig {
			return subscriber(config)
		}
		return &EurekaSubscriber{
			appSource: e.conn.NewAppSource(backendExtraConfig.EurekaAppName, true),
		}
	}
}

func (e *fargoEurekaClient) EnableInstance() {
	e.conn.UpdateInstanceStatus(e.appInstance, fargo.UP)
}

func NewFargoEurekaClient(appInstance *fargo.Instance, url string, logger logging.Logger) EurekaClient {
	logger.Info("Starting EurekaClient client")
	conn := startEurekaConnection(url, appInstance)
	return &fargoEurekaClient{conn, appInstance}
}

func startEurekaConnection(url string, appInstance *fargo.Instance) *fargo.EurekaConnection {
	conn := fargo.NewConn(url)
	conn.PollInterval = time.Second * time.Duration(POLL_INTERVAL)

	logger := log.NewLogfmtLogger(log.NewSyncWriter(os.Stderr))
	registrar := eureka.NewRegistrar(&conn, appInstance, logger)
	registrar.Register()

	return &conn
}

type EurekaConfig struct {
	AppName string `mapstructure:"app_name"`
	URL     string `mapstructure:"url"`
}

type EurekaSubscriber struct {
	appSource *fargo.AppSource
}

func (e EurekaSubscriber) Hosts() ([]string, error) {
	app := e.appSource.Latest()
	return hosts(app)
}

func hosts(app *fargo.Application) ([]string, error) {
	if app == nil {
		return nil, errors.New("Failed to get application metadata from eureka")
	}
	var hosts []string
	for _, instance := range app.Instances {
		if instance.Status == fargo.UP {
			hosts = append(hosts, fmt.Sprintf("http://%s:%d", instance.HostName, instance.Port))
		}
	}

	return hosts, nil
}
