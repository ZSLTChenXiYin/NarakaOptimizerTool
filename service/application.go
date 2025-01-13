package service

func Init() {
	if developerMode() {
		DebugLogger.Printf(NewLog("Configuration initialization begins."))
	}
	configInit()
	if developerMode() {
		DebugLogger.Printf(NewLog("Configuration initialization ends."))
	}

	if developerMode() {
		DebugLogger.Printf(NewLog("Router initialization begins."))
	}
	routerInit()
	if developerMode() {
		DebugLogger.Printf(NewLog("Router initialization ends."))
	}
}

func Start() {
	go func() {
		if developerMode() {
			DebugLogger.Printf(NewLog("Server initialization starts."))
		}
		err := r.Run()
		if err != nil {
			ErrorLogger.Fatalf(NewLog("Failed to start server: %v", err))
		}
	}()
}

func Stop() {
	if developerMode() {
		DebugLogger.Printf(NewLog("Configuration save begins."))
	}
	saveConfig()
	if developerMode() {
		DebugLogger.Printf(NewLog("Configuration save ends."))
	}
}
