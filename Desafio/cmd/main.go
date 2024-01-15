package main

import (
	"gotesting/Desafio/cmd/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	router.MapRoutes(r)

	r.Run(":18085")
}
