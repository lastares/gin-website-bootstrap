package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Category struct {
	Id int
	JumpLink string
	CategoryName string
}

func main()  {
	// 生成一个实例
	engine := gin.Default()
	// 加载静态文件设置
	engine.Static("/static", "./static")
	// 加载视图文件设置
	engine.LoadHTMLGlob("templates/*")

	// 官网导航数据
	categories := []Category{
		{Id: 2, JumpLink:"#home", CategoryName:"首页"},
		{Id: 3, JumpLink:"#bbs", CategoryName:"论坛"},
		{Id: 4, JumpLink:"#html5", CategoryName:"前端开发"},
		{Id: 5, JumpLink:"#course", CategoryName:"最新课程"},
		{Id: 6, JumpLink:"#app", CategoryName:"移动APP"},
		{Id: 7, JumpLink:"#contact", CategoryName:"联系我们"},
	}

	// 官网首页路由
	engine.GET("/go-website", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"categories": categories,
		})
	})
	engine.Run(":8081")
}


