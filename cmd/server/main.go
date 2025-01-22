package main

import (
	"vega-server/cmd/server/wire"
	"vega-server/pkg/config"

	"context"
	"flag"
)

// @title Vega-Server API
// @version 1.0
// @description This is a sample server for Vega-Server API.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8080
// @BasePath /api
func main() {
	// Initialize path
	path := flag.String("path", "config/dev.yaml", "path, eg: -path=config/dev.yaml")
	flag.Parse()
	// Initialize conf
	conf, err := config.LoadConfig(*path)
	if err != nil {
		panic(err)
	}
	// Initialize app
	app, err := wire.InitializeApp(conf)
	if err != nil {
		panic(err)
	}
	// Run app
	if err = app.Run(context.Background()); err != nil {
		panic(err)
	}
}
