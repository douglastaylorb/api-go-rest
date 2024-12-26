package main

import (
	"github.com/douglastaylorb/api-go-rest/database"
	"github.com/douglastaylorb/api-go-rest/routes"
)

func main() {

	database.ConectaBanco()

	routes.HandleRequest()
}
