package main

import (
	"fmt"
	"log"
	"tracker-manager-go/conf"
	_ "tracker-manager-go/conf"
	"tracker-manager-go/router"
)

func main() {
	app := router.New()
	log.Fatal(app.Listen(fmt.Sprintf(":%d", conf.App.Port)))
}
