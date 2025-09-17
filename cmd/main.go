package main

import (
	"pulsechain-assets/internal/manager"
)

func main() {
	manager.InitCommands()
	manager.Execute()
}
