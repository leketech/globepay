	// Initialize service factory
	logger.Info("Initializing service factory...")
	serviceFactory := service.NewFactory(cfg, db, redisClient, awsConfig)
	logger.Info("Service factory initialized")