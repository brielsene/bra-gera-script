package main

import (
	"chg-gera-script-brad/database"
	"chg-gera-script-brad/routes"
)

func main() {
	database.ConectaComDB()
	routes.HandleRequests()
}
