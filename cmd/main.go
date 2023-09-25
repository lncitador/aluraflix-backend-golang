package main

import (
	"github.com/lncitador/alura-flix-backend/config"
	"github.com/lncitador/alura-flix-backend/internal/delivery/http"
)

func main() {
	cnf := config.NewConfig()
	defer cnf.Close()

	routes := http.InitRoutes(cnf)
	routes.Register()

	cnf.RunServe()
}
