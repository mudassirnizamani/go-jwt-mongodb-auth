package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/mudasiralinizamani/ums-api/controllers"
	"github.com/mudasiralinizamani/ums-api/middlewares"
)

func UserRouter(app *gin.Engine) {
	app.Use(middlewares.Authenticate())
	app.GET("/users", controllers.GetUsers())
	app.GET("/users/:user_id", controllers.GetUser())
}
