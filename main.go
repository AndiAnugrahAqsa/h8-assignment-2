package main

import (
	"assignment2/database"
	"assignment2/route"

	"github.com/gin-gonic/gin"
)

func main() {
	database.InitDB()

	r := gin.Default()

	route.Router(r)

	r.Run() // listen and serve on 0.0.0.0:8080
}
