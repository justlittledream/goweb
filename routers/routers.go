package routers

import (
	controller "github/lhz/ginessential/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CollectRouter(r *gin.Engine) *gin.Engine {
	r.GET("/api/auth", func(c *gin.Context) {
		c.HTML(http.StatusOK, "register.html", gin.H{})
	})
	r.POST("/api/auth/register", controller.Register)
	return r
}
