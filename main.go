package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	// deklare framework gin
	r := gin.Default()

	r.GET("/ping", func(ctx *gin.Context) {
		// ctx.String(http.StatusOK, "hello world") ini kalo outputnya string
		ctx.JSON(http.StatusOK, gin.H{
			"status": "succes",
			"value":  "hello  fadhil ganteng",
		}) //gin.H untuk generate json bisa juga pake map : ctx.JSON(http.StatusOK,map[string]interface{}{"blabla"}

	})

	return r
}

func main() {
	r := setupRouter()

	r.Run(":8080")
}
