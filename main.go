package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joshuaautawi/shuttle/db"
	"github.com/joshuaautawi/shuttle/user"
)

func main() {
	r := gin.Default()
	DB := db.InitDB()
	user.Routes(r, DB)

	r.GET("/api/", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"message": "Hello"}) })
	r.Run()
}
