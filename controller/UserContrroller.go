package controller

import (
	"github/lhz/ginessential/common"
	"github/lhz/ginessential/model"
	util "github/lhz/ginessential/util"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

var db *gorm.DB

func Register(ctx *gin.Context) {
	//获取参数P
	DB := common.GetDB()
	name := ctx.PostForm("name")
	tel := ctx.PostForm("tel")
	passwd := ctx.PostForm("passwd")
	if len(tel) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号长度不对"})
		return
	}
	if len(passwd) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码长度不能少于6位"})
		return
	}
	//如果名称没有，给1个10位的随机符
	if len(name) == 0 {
		name = util.RandomString(10)
	}
	log.Println(name, tel, passwd)

	//判断手机号是否存在
	if isTelExists(DB, tel) {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "该用户已经存在"})
		return
	}
	newUser := model.User{
		Name:     name,
		Telphone: tel,
		Passwd:   passwd,
	}
	DB.Create(&newUser)
	ctx.JSON(200, gin.H{
		"msg": "注册成功",
	})
}

func isTelExists(db *gorm.DB, tel string) bool {
	var user model.User
	db.Where("telphone = ?", tel).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
