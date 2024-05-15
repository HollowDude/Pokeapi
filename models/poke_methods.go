package models

import (
	"GinPoke/initializers"

	"github.com/gin-gonic/gin"
)

func ShowIt() (team Poke) {
	initializers.DB.Find(&team)
	return
}

func AddIt(input Poke, c *gin.Context) (poke Poke, error error) {

	poke = Poke{Name: input.Name}

	initializers.DB.Create(&poke)

	return
}
func DelIt(id string) (error error) {
	var poke Poke
	if error = initializers.DB.Where("id = ?", id).First(&poke).Error; error != nil {
		return
	}
	initializers.DB.Delete(&poke)
	return
}
