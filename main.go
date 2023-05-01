package main

import (
	"go_perpustakaan/config"
	m "go_perpustakaan/middlewares"
	"go_perpustakaan/routes"
)

func main() {
	config.InitDB()
	e := routes.New()
	m.Log(e)
	e.Logger.Fatal(e.Start(":8000"))
}
