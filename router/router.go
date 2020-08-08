package router

import (
	"github.com/gin-gonic/gin"
	"noodlenote/controller/note"
	"noodlenote/controller/notebook"
	"noodlenote/controller/resource"
)

func NewRouter() *gin.Engine {
	e := gin.Default()

	e.GET("/ping", func(context *gin.Context) {
		context.Writer.WriteString("Pong")
	})

	e.NoRoute(func(context *gin.Context) {
		context.Writer.WriteString("Sorry! No Route")
	})

	noteBookRouter(e.Group("/notebook"))
	noteRouter(e.Group("/note"))
	resourceRouter(e.Group("/resource"))
	return e
}

func noteBookRouter(g *gin.RouterGroup) {
	g.GET("/all", notebookController.All)          //查看用户全部笔记本
	g.GET("/list/:id", notebookController.List)    //列出笔记本下的笔记
	g.PUT("/create", notebookController.Create)    //创建笔记本
	g.DELETE("/delete", notebookController.Delete) //删除笔记本
}
func noteRouter(g *gin.RouterGroup) {
	g.PUT("/create", noteController.Create)        //创建笔记 ?id=0(非必须) 默认0
	g.GET("/get/:id", noteController.Get)          //查看笔记
	g.PUT("/update", noteController.Update)        //修改笔记
	g.PUT("/move/:id", noteController.Move)        //移动笔记 ?id=0(非必须) 默认0
	g.DELETE("/delete/:id", noteController.Delete) //删除笔记
}

func resourceRouter(g *gin.RouterGroup) {
	g.POST("/upload", resourceController.Upload)        //上传附件或图片
	g.GET("/download/:id", resourceController.Download) //下载附件/图片/图标
}
