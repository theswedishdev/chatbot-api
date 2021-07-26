package commands

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/theswedishdev/chatbot-api/app/types"
)

func fetchChuckNorris() (types.ChuckNorrisResponse, error) {
	URL := "https://api.chucknorris.io/jokes/random"

	response, err := http.Get(URL); if err != nil {
		return types.ChuckNorrisResponse{}, err
	}

	bodyBytes, err := ioutil.ReadAll(response.Body); if err != nil {
		return types.ChuckNorrisResponse{}, err
	}

	var responseData types.ChuckNorrisResponse
	json.Unmarshal(bodyBytes, &responseData)

	return responseData, nil
}

func GetChuckNorrisJoke() (string, error) {
	joke, err := fetchChuckNorris(); if err != nil {
		return "", err
	}

	return joke.Value, nil
}
