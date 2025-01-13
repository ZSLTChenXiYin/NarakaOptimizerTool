package main

import (
	"github/ZSLTChenXiYin/NarakaOptimizerTool/service"
)

func init() {
	service.Init()

	service.InfoLogger.Printf(service.NewLog("Naraka Optimizer Tool Initialized"))
}

func main() {
	service.Start()

	service.InfoLogger.Printf(service.NewLog("Naraka Optimizer Tool Started"))

	service.Process(func() {
		service.Stop()

		service.InfoLogger.Printf(service.NewLog("Naraka Optimizer Tool Saved Config"))
	})

	service.InfoLogger.Printf(service.NewLog("Naraka Optimizer Tool Stopped"))
}
