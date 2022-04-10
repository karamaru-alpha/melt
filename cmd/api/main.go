package main

import (
	"os"

	"github.com/karamaru-alpha/melt/pkg/api"
	"github.com/karamaru-alpha/melt/pkg/logging/app"
)

func main() {
	os.Exit(cmd())
}

func cmd() (code int) {
	if err := app.SetLogger(os.Getenv("ENV") == "local"); err != nil {
		panic(err)
	}
	c := &api.Config{
		Port: os.Getenv("PORT"),
	}

	api.Serve(c)
	return 0
}
