package main

import (
	"github.com/PBKKE08/FP-BE/echo-rest/db"
	"github.com/PBKKE08/FP-BE/echo-rest/routes"
)

func main() {
	db.Init()
	e := routes.Init()

	e.Logger.Fatal(e.Start(":1234"))
}
