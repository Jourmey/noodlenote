package router

import (
	"github.com/gin-gonic/gin"
	"noodlenote/controller"
)

func NewRouter() *gin.Engine {
	e := gin.Default()

	e.GET("/ping", func(context *gin.Context) {
		context.Writer.WriteString("Pong")
	})

	e.NoRoute(func(context *gin.Context) {
		context.Writer.WriteString("Sorry! No Route")
	})

	FolderRouter(e.Group("/folder"))

	return e
}

func FolderRouter(g *gin.RouterGroup) {
	g.GET("/current", controller.GetCurrent)
	g.GET("/sub_file/:page", controller.GetSubFile)
	g.GET("/sub_folder", controller.GetSubFolders) //用于编辑文章选择目录时请求
	g.GET("/update", controller.Update)
	g.GET("/add", controller.Add)
	g.GET("/delete", controller.Delete)
}
