package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// https://api.chucknorris.io/jokes/random
// {
//   "categories": [
//     "travel"
//   ],
//   "created_at": "2020-01-05 13:42:19.576875",
//   "icon_url": "https://assets.chucknorris.host/img/avatar/chuck-norris.png",
//   "id": "9rqqcar_s5mqpujm0yu5za",
//   "updated_at": "2020-01-05 13:42:19.576875",
//   "url": "https://api.chucknorris.io/jokes/9rqqcar_s5mqpujm0yu5za",
//   "value": "For Spring Break '05, Chuck Norris drove to Madagascar, riding a chariot pulled by two electric eels."
// }
//
// https://mholt.github.io/json-to-go/
// type AutoGenerated struct {
// 	Categories []any  `json:"categories"`
// 	CreatedAt  string `json:"created_at"`
// 	IconURL    string `json:"icon_url"`
// 	ID         string `json:"id"`
// 	UpdatedAt  string `json:"updated_at"`
// 	URL        string `json:"url"`
// 	Value      string `json:"value"`
// }

type ChuckNorrisJoke struct {
	CreatedAt  string `json:"created_at"`
	IconURL    string `json:"icon_url"`
	ID         string `json:"id"`
	UpdatedAt  string `json:"updated_at"`
	URL        string `json:"url"`
	Value      string `json:"value"`
	Categories []any  `json:"categories"`
}

func get_joke() string {
	url := "https://api.chucknorris.io/jokes/random"

	res, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	var joke ChuckNorrisJoke
	json.Unmarshal(body, &joke)

	return joke.Value
}

func main() {
	joke := get_joke()
	fmt.Printf("%s\n", joke)
}
