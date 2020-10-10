package main

import (
	common "github/lhz/ginessential/common"
	routers "github/lhz/ginessential/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	common.InitDB()
	r := gin.Default()
	r = routers.CollectRouter(r)
	r.LoadHTMLGlob("templates/*")
	panic(r.Run()) // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
