package tmpl

var InitWeb = `package routers

import (
	"github.com/gin-gonic/gin"
)

func Service(app *gin.Engine)  {
	app.Use(gin.Logger())
	app.Use(gin.Recovery())
	router(app)
}

`

var InitRouter = `package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/{{.Group}}/{{.Project}}/controllers"
)

func router(app *gin.Engine) {
	srv := app.Group("/api/v1")
	srv.GET("/", func(c *gin.Context) {c.JSON(200, gin.H{"msg": "ok"})})  //健康检查
	srv.GET("/user/info", controllers.GetUser)                            //获取用户信息
    srv.POST("/user/info", controllers.CreateUser)                        //创建用户
}

`
