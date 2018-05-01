package main

import (
	"github.com/urfave/negroni"
	"github.com/wty/api"
	"github.com/wty/models"
	"github.com/wty/routes"
)

func main() {
	db := models.NewSqliteDB("data.db")
	api := api.NewAPI(db)
	routes := routes.NewRoutes(api)
	n := negroni.Classic()
	n.UseHandler(routes)
	n.Run(":3000")
}
