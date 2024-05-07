package main

import (
	"os"

	"github.com/Corray333/portfolio/internal/app"
	"github.com/Corray333/portfolio/internal/config"
)

func main() {
	config.MustInit(os.Args[1])
	app.New().Run()
}
