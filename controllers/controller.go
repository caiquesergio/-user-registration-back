package controllers

import (
	"net/http"

	"github.com/caiquesergio/api-go-gin/database"
	"github.com/caiquesergio/api-go-gin/models"
	"github.com/gin-gonic/gin"
)

func ExibeTodosUsuarios(c *gin.Context) {
	var usuarios []models.User
	database.DB.Find(&usuarios)
	c.JSON(200, usuarios)
}

func CriaNovoUsuario(c *gin.Context) {
	var usuario models.User
	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Create(&usuario)
	c.JSON(http.StatusOK, usuario)
}

func DeletaUsuario(c *gin.Context) {
	var usuario models.User
	id := c.Params.ByName("id")
	database.DB.Delete(&usuario, id)
	c.JSON(http.StatusOK, gin.H{"data": "usuario deletado com sucesso"})
}

func EditaUsuario(c *gin.Context) {
	var usuario models.User
	id := c.Params.ByName("id")
	database.DB.First(&usuario, id)
	if err := c.ShouldBindJSON(&usuario); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error()})
		return
	}
	database.DB.Model(&usuario).UpdateColumns(usuario)
	c.JSON(http.StatusOK, usuario)
}

func BuscarLoginUsuario(c *gin.Context) {
	var usuario models.User
	email := c.Param("email")
	database.DB.Where(&models.User{Email: email}).First(&usuario)

	if usuario.ID == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"Notfound": "Usuario não encontrado"})
		return
	}

	c.JSON(http.StatusOK, usuario)
}
