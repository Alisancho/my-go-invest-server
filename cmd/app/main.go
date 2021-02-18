package main

import "github.com/alicancho/my-go-invest-server/internal/app"

const configPath = "configs/main"

func main() {
	app.Run(configPath)
}
