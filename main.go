package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Test struct { //untuk nampung post body
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

	//digrupin "v1" masukin ke var v1,, jadinya localhost:8080/v1/endpoint
	v1 := r.Group("v1")
	// cara gunain grupnya
	v1.GET("/user/:name", func(ctx *gin.Context) {
		param := ctx.Param("name") //simplenya ini req.param
		ctx.JSON(http.StatusOK, gin.H{
			"status": "succes",
			"value":  param,
		})
	})

	v1.POST("user", func(ctx *gin.Context) {
		var data Test //struct untuk nangkep body

		ctx.BindJSON(&data) //kenapa ada '&' alias pointer? karna hasil dari bind dibalikan ke data melalui pointer
		//karna kalo kita liat bind itu response nya error makanya harus melalu '&' alias pointer
		//simplenya ini Adalah req.body
		ctx.JSON(http.StatusOK, gin.H{
			"status": "succes",
			"value":  data,
		})
	})

	v1.GET("/user", func(ctx *gin.Context) {

		query := ctx.Query("name") //simplenya ini req.query
		ctx.JSON(http.StatusOK, gin.H{
			"status": "succes",
			"value":  query,
		})
	})

	return r
}

func main() {
	r := setupRouter()

	r.Run(":8080")
}
