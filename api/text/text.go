package text

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/workexercise2/model/text"
)

func New(context *gin.Context) {

	e := text.Newtext()
	if e != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "数据错误",
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"message": "创建成功",
	})
}
