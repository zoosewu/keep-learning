package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Option func(*gin.RouterGroup)

var optionPair = map[string][]Option{}

// 注册app的路由配置
func Include(path string, opts ...Option) {
	option, isExist := optionPair[path]
	if isExist {
		optionPair[path] = append(option, opts...)
	} else {
		optionPair[path] = opts
	}
}

func showRouters(router *gin.Engine) {
	router.GET("/", func(ctx *gin.Context) {
		if gin.Mode() == gin.DebugMode {
			ctx.Data(200, "application/json; charset=utf-8", []byte(fmt.Sprintf("%v", router.Routes())))
		}
	})
}

// 初始化
func Init() *gin.Engine {
	rootRouter := gin.Default()
	showRouters(rootRouter)
	apiRouter := rootRouter.Group("/api")
	fmt.Printf("Router count: %v\n", len(optionPair))
	for path, options := range optionPair {
		fmt.Printf("Add router %v, count: %v\n", path, len(options))
		router := apiRouter.Group(path)
		for _, opt := range options {
			opt(router)
		}
	}
	return rootRouter
}
