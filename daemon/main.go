package main

import (
	"github.com/nicholaskh/cottage/config"
	"github.com/nicholaskh/cottage/server"
	serverlib "github.com/nicholaskh/golib/server"
)

func init() {
	parseFlags()

	if options.showVersion {
		serverlib.ShowVersionAndExit()
	}

	serverlib.SetupLogging(options.logFile, options.logLevel, options.crashLogFile)

	conf := serverlib.LoadConfig(options.configFile)
	config.Cottage = new(config.CottageConfig)
	config.Cottage.LoadConfig(conf)
}

func main() {
	cottageServer := server.NewCottageServer(config.Cottage)
	cottageServer.Launch(config.Cottage.ListenAddr)
}
