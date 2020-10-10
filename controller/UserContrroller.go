package controller

import (
	common "github/lhz/ginessential/common"
	"github/lhz/ginessential/model"
	util "github/lhz/ginessential/util"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"

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

	//创建用户
	hasedPasswd, err := bcrypt.GenerateFromPassword([]byte(passwd), bcrypt.DefaultCost)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "加密失败"})
		return
	}
	newUser := model.User{
		Name:     name,
		Telphone: tel,
		Passwd:   string(hasedPasswd),
	}
	DB.Create(&newUser)
	ctx.JSON(200, gin.H{
		"msg": "注册成功",
	})
}

func Login(ctx *gin.Context) {
	//获取参数
	DB := common.GetDB()
	tel := ctx.PostForm("tel")
	passwd := ctx.PostForm("passwd")

	//数据验证
	if len(tel) != 11 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "手机号长度不对"})
		return
	}
	if len(passwd) < 6 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "密码长度不能少于6位"})
		return
	}

	//判断手机号是否存在
	var user model.User
	DB.Where("telphone = ?", tel).First(&user)
	if user.ID == 0 {
		ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "用户不存在"})
		return
	}

	//判断密码是否正确
	if err := bcrypt.CompareHashAndPassword([]byte(user.Passwd), []byte(passwd)); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"code": 400, "msg": "密码错误"})
		return
	}

	//发放token
	token, err := common.ReleaseToken(user)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"code": 500, "msg": "系统异常"})
		log.Printf("token generate error: %v", err)
		return
	}

	//返回结果
	ctx.JSON(200, gin.H{
		"code": 200,
		"msg":  "登录成功",
		"data": gin.H{"token": token},
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

func Info(ctx *gin.Context) {
	user, _ := ctx.Get("user")
	ctx.JSON(http.StatusOK, gin.H{"code": 200, "data‘": gin.H{"user": user}})
}
