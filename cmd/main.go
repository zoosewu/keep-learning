package main

import (
	"fmt"
	"zoosewu/vocabulary/cmd/api/area"
	"zoosewu/vocabulary/cmd/api/content"
	"zoosewu/vocabulary/internal/router"
)

func main() {
	// 加载多个APP的路由配置
	router.Include("/v1", content.Router)
	router.Include("/v1", area.Router)
	// 初始化路由
	r := router.Init()
	if err := r.Run("localhost:8080"); err != nil {
		fmt.Printf("startup service failed, err:%v\n", err)
	}
}
