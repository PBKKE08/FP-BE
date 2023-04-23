package main

import (
	"github.com/PBKKE08/FP-BE/echo-rest/routes"
)

func main() {
	e := routes.Init()

	e.Logger.Fatal(e.Start(":8080"))
}
