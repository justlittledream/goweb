package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"time"
	"gorm.io/driver/mysql"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `gorm:"type:varchar(20);not null"`
	Telphone string `gorm:"type:varchar(110);not null;unique"`
	Passwd   string `gorm:"type:varchar(255);not null"`
}

func main() {
	db := InitDB()
	//defer db.Close()
	// if err != nil {
	// 	panic("连接数据库失败")
	// }
	//defer db.Close()

	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/api/auth", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", gin.H{})
	})
	r.POST("/api/auth/register", func(ctx *gin.Context) {
		//获取参数
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
			name = RandomString(10)
		}
		log.Println(name, tel, passwd)

		//判断手机号是否存在
		if isTelExists(db, tel) {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{"code": 422, "msg": "该用户已经存在"})
			return
		}
		newUser := User{
			Name:     name,
			Telphone: tel,
			Passwd:   passwd,
		}
		db.Create(&newUser)
		ctx.JSON(200, gin.H{
			"msg": "注册成功",
		})
	})
	panic(r.Run()) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func RandomString(n int) string {
	var letters = []byte("asdfghjklzxcvbnmqwertyuiopZXCVBNMASDFGHJKLQWERTYUIOP")
	result := make([]byte, n)
	rand.Seed(time.Now().Unix())
	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

func InitDB() *gorm.DB {
	//driverName := "mysql"
	host := "localhost"
	port := "3306"
	database := "go"
	username := "root"
	password := "root"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true",
		username,
		password,
		host,
		port,
		database,
		charset)
	db, err := gorm.Open(mysql.Open(args), &gorm.Config{})
	if err != nil {
		panic("failed to connect to mysql, err :" + err.Error())
	}
	db.AutoMigrate(&User{})
	return db
}

func isTelExists(db *gorm.DB, tel string) bool {
	var user User
	db.Where("Telphone = ?", tel).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
