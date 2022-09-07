package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"txc-njl/server/Cors"
	"txc-njl/server/jwtdemo"
)

func SetRouters() *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies([]string{"localhost"})

	r.Use(Cors.CORS())
	r.Use(gin.Recovery())
	r.POST("/login", login)
	r.POST("/register", register)
	r.GET("/getgoods", getGoodsList)
	r.POST("/getgoodsbyid", getById)

	r.POST("/creategoods", createGoods)
	r.POST("/createcomment", createComment)

	r.POST("/createReply", createReply)

	r.POST("/upload", uploadFile)

	r.GET("/ping", jwtdemo.Middleware(), func(c *gin.Context) {
		if userName, exists := c.Get("userName"); exists {
			fmt.Println("当前用户", userName)
		}
		c.JSON(http.StatusOK, gin.H{
			"code":    0,
			"message": "验证成功",
		})

	})

	return r
}
