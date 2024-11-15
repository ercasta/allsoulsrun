package main

import (
	"github.com/ercasta/allsoulsrun/pkg/game/run"
	"github.com/gin-gonic/gin"
)

// func main() {
// 	allsouls.Rungame()
// }

func main() {

	r := gin.Default()

	public := r.Group("/api")

	//public.POST("/register", controllers.Register)

	public.POST("/dorun", run.Rungame)
	// protected := r.Group("/api/admin")
	// protected.Use(middleware.JwtAuthMiddleware())
	// protected.GET("/user", controllers.CurrentUser)

	r.Run(":8080")

}
