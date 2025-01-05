package main

import (
	"fmt"
	"os"

	"github.com/milesjpool/kahd/cmd/api-server/internal/database"
	"github.com/milesjpool/kahd/cmd/api-server/internal/logging"
	"github.com/milesjpool/kahd/cmd/api-server/internal/server"
)

var Logger logging.Logger = &logging.StdIOLogger{}
var service server.Server = &APIServer{
	Logger:            Logger,
	ConfigLoader:      &EnvConfigLoader{},
	DatabaseConnector: &database.PostgresDatabaseConnector{},
	ServerFactory: &server.HTTPServerFactory{
		Logger: Logger,
	},
}

func main() {
	if err := service.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
