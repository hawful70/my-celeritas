package main

import (
	"myapp/handlers"

	celeritas "github.com/hawful70/my-celeritas"
)

type application struct {
	App      *celeritas.Celeritas
	Handlers *handlers.Handlers
}

func main() {
	c := initApplication()
	c.App.ListenAndServe()
}
