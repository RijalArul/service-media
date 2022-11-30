package main

import (
	"service-media/databases"
	"service-media/routes"
)

func main() {
	databases.StartDB()
	routes.Routes()
}
