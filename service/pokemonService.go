package service

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"pustaka-api/domain"
)

func FindPokemonsByRegion(regionName string) domain.PokemonResponse {
	response, err := http.Get("http://pokeapi.co/api/v2/pokedex/" + regionName)
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject domain.PokemonResponse
	json.Unmarshal(responseData, &responseObject)

	fmt.Println(responseObject.Name)
	fmt.Println(len(responseObject.Pokemon))

	for i := 0; i < len(responseObject.Pokemon); i++ {
		fmt.Println(responseObject.Pokemon[i].Species.Name)
	}

	return responseObject
}
