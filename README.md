# krakend-eureka
krakend eureka integration.

The integration is based on a custom subscriber using the [fargo library](https://github.com/hudl/fargo)

## Usage Example
```go
	logger, err := logging.NewLogger("INFO", os.Stdout, "[KRAKEND]")
	if err != nil {
		log.Fatal("ERROR:", err.Error())
	}
	parser := config.NewParser()
	serviceConfig, err := parser.Parse("./test.json")
	if err != nil {
		log.Fatal("ERROR:", err.Error())
	}

	eurekaLocalAppInstance, err := NewLocalAppInstance(8000, "HELLO")
	if err != nil {
		log.Fatal("ERROR:", err.Error())
	}
	eurekaClient := NewFargoEurekaClient(eurekaLocalAppInstance, "http://localhost:8080", logger)

  subscriber := eurekaClient.NewSubscriber(sd.GetSubscriber)
	routerFactory := kgin.DefaultFactory(proxy.DefaultFactoryWithSubscriber(logger, subscriber), logger)
	routerFactory.New().Run(serviceConfig)
```

We need to create an app instance to register the current application to eureka:
```go
eurekaLocalAppInstance, err := NewLocalAppInstance(8000, "HELLO")
```
An aws aware application instance can be used: **NewAwsAppInstance**

Then, a eureka client should be create with:
```go
eurekaClient := NewFargoEurekaClient(eurekaLocalAppInstance, "http://localhost:8080", logger)
```

And wire it with krakend with a new Subscriber
```go
subscriber := eurekaClient.NewSubscriber(sd.GetSubscriber)
routerFactory := kgin.DefaultFactory(proxy.DefaultFactoryWithSubscriber(logger, subscriber), logger)
```

## Config Example
```yml
{
  "version": 2,
  "max_idle_connections": 250,
  "timeout": "3000ms",
  "read_timeout": "0s",
  "write_timeout": "0s",
  "idle_timeout": "0s",
  "read_header_timeout": "0s",
  "name": "Test",
  "endpoints": [
    {
      "endpoint": "/cbcrash",
      "method": "GET",
      "backend": [
        {
          "url_pattern": "/crash",
          "host": [
            "http://localhost:8000"
          ],
          "extra_config": {
            "github.com/schibsted/krakend-eureka": {
              "eureka_app_name": "crash",
            }
          }
        }
      ],
      "timeout": "1500ms",
      "max_rate": "10000"
    }
  ]
}
```

The only configuration needed it's the eureka backend application name