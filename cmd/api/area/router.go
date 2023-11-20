package area

import (
	"github.com/gin-gonic/gin"
)

func Router(router *gin.RouterGroup) {
	router.Group("/area").
		GET("", GetAreas).
		POST("", PostArea).
		GET("/:id", GetAreaByID)
}
