package app

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ResponseError(ctx *gin.Context, err error) {
	if ctx == nil {
		fmt.Println(err)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": err.Error(),
	})
}

func ResponseSuccess(ctx *gin.Context, message string) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": message,
	})
}

func ResponseSearchSuccess(ctx *gin.Context, total interface{}, rows interface{}) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "查询成功",
		"total":   total,
		"list":    rows,
	})
}
