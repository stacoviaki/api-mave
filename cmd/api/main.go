package main

import (
	"github.com/gin-gonic/gin"
	"github.com/stacoviaki/api-mave.git/db"
)

func main() {
	db.Connect()

	r := gin.Default()

	r.Run(":8080")
}
