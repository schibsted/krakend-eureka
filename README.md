# krakend-eureka
krakend-eureka integration


Example
``` 
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

	routerFactory := kgin.DefaultFactory(proxy.DefaultFactoryWithSubscriber(logger, eurekaClient.NewSubscriber(sd.GetSubscriber)), logger)
	routerFactory.New().Run(serviceConfig)
```

Config Example
```
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
            "github.com/tgracchus/krakend-eureka": {
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