package models

type Poke struct {
	ID   int    `json:"id" gorm:"primary_key"`
	Name string `json:"name"`
}

type PokeStruct struct {
	Name string `json:"name" binding:"required"`
}

/*
type PokeSprites struct {
	Front_default string `json:"front_default"`
	Front_shiny   string `json:"front_shiny"`
}
*/
