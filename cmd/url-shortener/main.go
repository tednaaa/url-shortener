package main

import (
	"fmt"
	"url-shortener/internal/config"
)

const (
	envLocal       = "local"
	envDevelopment = "development"
	envProduction  = "production"
)

func main() {
	config := config.MustLoad()

	fmt.Println(config)
}
