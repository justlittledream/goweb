package routers

import (
	controller "github/lhz/ginessential/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CollectRouter(r *gin.Engine) *gin.Engine {
	r.GET("/api/auth/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{})
	})
	r.GET("/api/auth/regis", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", gin.H{})
	})
	r.POST("/api/auth/login", controller.Login)
	r.POST("/api/auth/register", controller.Register)
	return r
}
