package main

import (
	"GinPoke/initializers"
	"GinPoke/routes"
	"log"
	/*"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"*/)

/*
	func ShowMyTeam(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"My Team:": team,
		})
	}

	func SearchByName(c *gin.Context) {
		response, err := http.Get("http://pokeapi.co/api/v2/pokemon/" + c.Param("name") + "/")
		if err != nil {
			c.JSON(404, gin.H{
				"message:": err,
			})
			return
		}

		responsedata, err2 := ioutil.ReadAll(response.Body)
		if err2 != nil {
			c.JSON(404, gin.H{
				"message:": err2,
			})
			return
		}

		var responseObject Response
		json.Unmarshal(responsedata, &responseObject)

		c.JSON(302, gin.H{
			"data": responseObject,
		})

}

	func AddToTeam(c *gin.Context) {
		response, err := http.Get("http://pokeapi.co/api/v2/pokemon/" + c.Param("name") + "/")
		if err != nil {
			c.JSON(404, gin.H{
				"message:": err,
			})
			return
		}

		responsedata, err2 := ioutil.ReadAll(response.Body)
		if err2 != nil {
			c.JSON(404, gin.H{
				"message:": err2,
			})
			return
		}

		var responseObject Response
		json.Unmarshal(responsedata, &responseObject)

		team = append(team, responseObject)

		c.JSON(http.StatusCreated, gin.H{
			"Se añadio:": responseObject,
		})

}

	func DelFromTeam(c *gin.Context) {
		var name = c.Param("name")

		for i, val := range team {
			if val.Name == name {
				team = append(team[:i], team[i+1:]...)
				c.JSON(http.StatusOK, gin.H{
					"message": "Pokemon is now free!! UwU",
				})
				return
			}
		}

		c.JSON(404, gin.H{
			"message": "Pokemon no existente en tu team!",
		})

}
*/
func main() {

	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
	routes.InitRoutes()

}

/*var team = []Response{}

type Response struct {
	Id      int              `json:"id"`
	Name    string           `json:"name"`
	Sprites []PokemonSprites `json:"sprites"`
}

type PokemonSprites struct {
	Front_default string `json:"front_default"`
	Front_shiny   string `json:"front_shiny"`
}
*/
