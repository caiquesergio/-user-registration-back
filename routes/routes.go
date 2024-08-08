package routes

import (
	"github.com/caiquesergio/api-go-gin/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func HandleRequests() {
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/usuarios", controllers.ExibeTodosUsuarios)
	r.POST("/usuarios", controllers.CriaNovoUsuario)
	r.DELETE("/usuarios/:id", controllers.DeletaUsuario)
	r.PATCH("/usuarios/:id", controllers.EditaUsuario)
	r.GET("/usuarios/:email", controllers.BuscarLoginUsuario)
	r.Run(":8080")
}
