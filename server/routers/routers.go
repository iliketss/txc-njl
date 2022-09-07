package routers

import (
	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"txc-njl/server/DB"
	"txc-njl/server/model"
	"txc-njl/server/utils"

	"net/http"
)

var db = DB.InitDB()

func login(c *gin.Context) {

	user := model.User{}
	err := c.Bind(&user)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	reUser := model.User{}

	db.Where("name=?", user.Name).First(&reUser)

	if reUser.ID == 0 {
		c.JSON(202, gin.H{
			"code": 202,
			"msg":  "用户不存在",
		})
		return
	} else {

		//todo
		token, err := utils.Login(&reUser)
		if err != nil {
			panic(err)
		}

		if utils.ComparePwd(reUser.Password, user.Password) {
			c.JSON(200, gin.H{
				"code":  200,
				"msg":   "login success",
				"user":  reUser,
				"token": token,
			})
		} else {
			c.JSON(201, gin.H{
				"code": 201,
				"msg":  "密码错误",
			})
		}

	}
}

//注册
func register(c *gin.Context) {
	user := model.User{}

	err := c.Bind(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}

	//fmt.Println("-----------------", user)
	reUSer := model.User{}
	db.Where("name=?", user.Name).First(&reUSer)
	if reUSer.ID == 0 {
		pwd, _ := utils.GenPwd(user.Password)
		user.Password = string(pwd)
		db.Create(&user)
		c.JSON(200, gin.H{
			"code": 200,
			"msg":  "register success",
			"user": user,
		})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{
			"code": 1,
			"msg":  "已存在该用户",
		})
		user = model.User{}
		return
	}

	/*reUser := model.User{}
	db.Where("name = ?", user.Name, user.Password).First(&reUser)*/

}

func getGoodsList(c *gin.Context) {
	var goods []model.Goods
	db.Preload("Comment").Preload("Comment.Reply").Find(&goods)
	c.JSON(http.StatusOK, gin.H{
		"data": goods,
	})
}

func getById(c *gin.Context) {
	id := c.PostForm("id")
	goods := model.Goods{}
	db.First(&goods, id)
	c.JSON(200, gin.H{
		"data": goods,
	})

}

func createGoods(c *gin.Context) {
	goods := model.Goods{}
	c.Bind(&goods)
	db.Create(&goods)
	c.JSON(200, gin.H{
		"code": "200",
		"data": goods,
	})
}

func createComment(c *gin.Context) {
	comment := model.Comment{}
	c.Bind(&comment)
	db.Create(&comment)
	c.JSON(200, gin.H{
		"code": "200",
		"data": comment,
	})
}

func createReply(c *gin.Context) {
	reply := model.Reply{}
	c.Bind(&reply)
	db.Create(&reply)
}

func uploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.String(500, "上传图片出错")
	}
	// c.JSON(200, gin.H{"message": file.Header.Context})
	c.SaveUploadedFile(file, file.Filename)
	c.String(http.StatusOK, file.Filename)
}
