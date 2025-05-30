package main

import (
	"myapp/data"
	"myapp/handlers"

	celeritas "github.com/hawful70/my-celeritas"
)

type application struct {
	App      *celeritas.Celeritas
	Handlers *handlers.Handlers
	Models   data.Models
}

func main() {
	c := initApplication()
	c.App.ListenAndServe()
}
