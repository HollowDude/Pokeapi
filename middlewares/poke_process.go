package middlewares

import (
	"GinPoke/models"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func Pulling(name string) (input models.Poke, error error) {
	response, err := http.Get("http://pokeapi.co/api/v2/pokemon/" + name)
	if err != nil {
		error = err
		return
	}

	responsedata, err2 := ioutil.ReadAll(response.Body)
	if err2 != nil {
		error = err2
		return
	}
	json.Unmarshal(responsedata, &input)
	return
}
