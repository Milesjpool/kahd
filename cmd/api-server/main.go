package main

import (
	"fmt"
	"os"

	"github.com/milesjpool/kahd/cmd/api-server/internal/database"
	"github.com/milesjpool/kahd/cmd/api-server/internal/logging"
	"github.com/milesjpool/kahd/cmd/api-server/internal/server"
)

var service server.Server = &APIServer{
	ConfigLoader:      &EnvConfigLoader{},
	DatabaseConnector: &database.PostgresDatabaseConnector{},
	ServerFactory: &server.HTTPServerFactory{
		Logger: &logging.StdIOLogger{},
	},
}

func main() {
	if err := service.Start(); err != nil {
		fmt.Fprintf(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}
