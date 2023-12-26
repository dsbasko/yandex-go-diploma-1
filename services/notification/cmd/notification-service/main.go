package main

import (
	"log"

	"github.com/dsbasko/yandex-go-diploma-1/services/notification/internal/app"
)

func main() {
	if err := app.Run(); err != nil {
		log.Panicf("app.Run: %v", err)
	}
}
