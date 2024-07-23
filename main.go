package main

import (
	initializers "github.com/Shreedeepdatta/rankandmarks/Initializers"
	"github.com/Shreedeepdatta/rankandmarks/controllers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.Loadenv()
	initializers.DatabaseConn()
	initializers.SyncDB()
}
func main() {
	r := gin.Default()
	r.GET("/student", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "found",
		})
	})
	r.POST("/enroll", controllers.SignUp)
	r.Run()
}
