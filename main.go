package main

import (
	"go_perpustakaan/config"
	"go_perpustakaan/routes"
)

func main() {
	config.InitDB()
	e := routes.New()

	e.Logger.Fatal(e.Start(":8000"))
}
