package routers

import "github.com/gin-gonic/gin"

func ApiRoutersInit(r *gin.Engine) {
	apiRouters := r.Group("/api")
	{
		apiRouters.GET("/user", func(c *gin.Context) {
			c.String(200, "user")
		})
		apiRouters.GET("/user/:id", func(c *gin.Context) {
			id := c.Param("id")
			c.String(200, "user id is %s", id)
		})
	}
}
