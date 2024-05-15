package controllers

import (
	"GinPoke/initializers"
	"GinPoke/middlewares"
	"GinPoke/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ShowMyTeam(c *gin.Context) {
	var team []models.Poke
	initializers.DB.Find(&team)
	c.JSON(http.StatusOK, gin.H{
		"My Team:": team,
	})
}

func AddToTeam(c *gin.Context) {
	input, err := middlewares.Pulling(c.Param("name"))
	if err != nil {
		c.JSON(404, gin.H{
			"message": err,
		})
		return
	}
	poke2, err2 := models.AddIt(input, c)
	if err2 != nil {
		c.JSON(404, gin.H{
			"message": err,
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"Se añadio:": poke2,
	})
}

func SearchByName(c *gin.Context) {
	response, err := middlewares.Pulling(c.Param("name"))
	if err != nil {
		c.JSON(404, gin.H{
			"message": err,
		})
		return
	}

	c.JSON(302, gin.H{
		"data": response,
	})
}

func DelFromTeam(c *gin.Context) {
	/*var id, error = strconv.Atoi(c.Param("id"))
	if error != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Bad data",
		})
		return
	}*/

	err := models.DelIt(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": true})
}
