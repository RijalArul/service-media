package main

import (
	"service-media/databases"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	databases.StartDB()
	r.Run()
}
