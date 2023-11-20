package content

import (
	"github.com/gin-gonic/gin"
)

func Router(router *gin.RouterGroup) {
	router.Group("/content").
		GET("", GetContents).
		POST("", PostContent).
		GET("/:id", GetContentByID)
}
