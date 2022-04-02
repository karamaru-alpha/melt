package main

import (
	"os"

	"github.com/karamaru-alpha/melt/pkg/api"
)

func main() {
	os.Exit(cmd())
}

func cmd() (code int) {
	c := &api.Config{
		Port: os.Getenv("PORT"),
	}

	api.Serve(c)
	return 0
}
