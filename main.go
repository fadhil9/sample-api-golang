package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Test struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

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

	//digrupin "v1" masukin ke var v1
	v1 := r.Group("v1")
	// cara gunain grupnya
	v1.GET("/user/:name", func(ctx *gin.Context) {
		param := ctx.Param("name")
		ctx.JSON(http.StatusOK, gin.H{
			"status": "succes",
			"value":  param,
		})
	})
	v1.POST("user")
	return r
}

func main() {
	r := setupRouter()

	r.Run(":8080")
}
