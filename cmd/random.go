/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/spf13/cobra"
)

// randomCmd represents the random command
var randomCmd = &cobra.Command{
	Use:   "random",
	Short: "gives a random comic character name",
	Long:  "this command gives a random comic character name",
	Run: func(cmd *cobra.Command, args []string) {
		getRandomCharacter()
	},
}

func init() {
	rootCmd.AddCommand(randomCmd)
}

type ComicCharacter struct {
	ID          int64       `json:"id"`
	Name        string      `json:"name"`
	Slug        string      `json:"slug"`
	Powerstats  Powerstats  `json:"powerstats"`
	Appearance  Appearance  `json:"appearance"`
	Biography   Biography   `json:"biography"`
	Work        Work        `json:"work"`
	Connections Connections `json:"connections"`
	Images      Images      `json:"images"`
}

type Appearance struct {
	Gender    string   `json:"gender"`
	Race      string   `json:"race"`
	Height    []string `json:"height"`
	Weight    []string `json:"weight"`
	EyeColor  string   `json:"eyeColor"`
	HairColor string   `json:"hairColor"`
}

type Biography struct {
	FullName        string   `json:"fullName"`
	AlterEgos       string   `json:"alterEgos"`
	Aliases         []string `json:"aliases"`
	PlaceOfBirth    string   `json:"placeOfBirth"`
	FirstAppearance string   `json:"firstAppearance"`
	Publisher       string   `json:"publisher"`
	Alignment       string   `json:"alignment"`
}

type Connections struct {
	GroupAffiliation string `json:"groupAffiliation"`
	Relatives        string `json:"relatives"`
}

type Images struct {
	Xs string `json:"xs"`
	Sm string `json:"sm"`
	Md string `json:"md"`
	Lg string `json:"lg"`
}

type Powerstats struct {
	Intelligence int64 `json:"intelligence"`
	Strength     int64 `json:"strength"`
	Speed        int64 `json:"speed"`
	Durability   int64 `json:"durability"`
	Power        int64 `json:"power"`
	Combat       int64 `json:"combat"`
}

type Work struct {
	Occupation string `json:"occupation"`
	Base       string `json:"base"`
}

func getRandomCharacter() {
	base_url := "https://akabab.github.io/superhero-api/api/id/"
	comic_character := ComicCharacter{}
	response := getCharacterData(base_url)

	err := json.Unmarshal(response, &comic_character)
	if err != nil {
		fmt.Println("cound not unmarshal response %v", err)
	}

	fmt.Println(string(comic_character.Name))
}

func getCharacterData(base_url string) []byte {
	// generate a random number between 1 and 731 since we only have 731 comic characters available in the superhero api.
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)
	rand_num := r1.Intn(731)

	// form the request url
	url := base_url + strconv.Itoa(rand_num) + ".json"

	// getting the response from the superhero api
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("could not make request. %v", err)
	}

	// reading the response body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("could not read response. %v", err)
	}
	return body
}
