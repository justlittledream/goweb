package main

import (
	common "github/lhz/ginessential/common"
	routers "github/lhz/ginessential/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	common.InitDB()

	// if err != nil {
	// 	panic("连接数据库失败")
	// }
	//defer db.Close()

	r := gin.Default()
	r = routers.CollectRouter(r)
	r.LoadHTMLGlob("templates/*")
	panic(r.Run()) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
