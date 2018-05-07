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