package main

import "github.com/fsandov/logs"

func main() {
	logsService := logs.NewService(logs.Service{
		FileLog: true,
	})
	logsService.Trace("Hello World!", "from custom service")
	logsService.Debug("Hello World!")
	logsService.Info("Hello World!")
	logsService.Notice("Hello World!")
	logsService.Warning("Hello World!")
	logsService.Error("Hello World!")
	logsService.Fatal("Hello World!")
	logs.Trace("Hello World!, from default service")
	logs.Debug("Hello World!, from default service")
	logs.Info("Hello World!, from default service")
	logs.Notice("Hello World!, from default service")
	logs.Warning("Hello World!, from default service")
	logs.Error("Hello World!, from default service")
	logs.Fatal("Hello World!, from default service")

}
