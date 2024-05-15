package routes

import (
	"GinPoke/controllers"
	"os"

	//"GinPoke/models"

	"github.com/gin-gonic/gin"
)

func InitRoutes() {
	gin.DefaultWriter = os.Stdout
	r := gin.Default()

	//models.ConnectDB()

	r.GET("/pokemon", controllers.ShowMyTeam)
	r.GET("/pokemon/:name", controllers.SearchByName)
	r.POST("/pokemon/:name", controllers.AddToTeam)
	r.DELETE("/pokemon/:id", controllers.DelFromTeam)

	err := r.Run()
	if err != nil {
		return
	}

}
