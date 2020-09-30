package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.New()
	routers(r)
	r.Run(":8080")
}


func routers(r *gin.Engine){
	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{"desc":"ok"})
	})
	v1 := r.Group("/api/v1/todos")
	{
		v1.POST("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK,gin.H{
				"desc":"post",
			})
		})
		v1.GET("/", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK,gin.H{
				"desc":"get",
			})
		})
	}
}