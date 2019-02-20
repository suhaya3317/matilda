package main

import (
	"matilda/backend/infrastructure"

	"google.golang.org/appengine"
)

func main() {
	infrastructure.RegisterHandlers()
	appengine.Main()
}
