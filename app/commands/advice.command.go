package commands

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/theswedishdev/chatbot-api/app/types"
)

func addExtraBrace(b []byte) []byte {
	if strings.HasSuffix(string(b[:]), "}}") {
		return b
	}
	return append(b, '}')
}

func fetchAdvice(ID string) (types.AdviceResponse, error) {
	URL := "https://api.adviceslip.com/advice"
	if ID != "" {
		URL = fmt.Sprintf("%s/%s", URL, ID)
	}

	response, err := http.Get(URL); if err != nil {
		return types.AdviceResponse{}, err
	}

	bodyBytes, err := ioutil.ReadAll(response.Body); if err != nil {
		return types.AdviceResponse{}, err
	}

	bodyBytes = addExtraBrace(bodyBytes)

	var responseData types.AdviceResponse
	json.Unmarshal(bodyBytes, &responseData)

	if responseData.HasError() {
		return types.AdviceResponse{}, errors.New(responseData.GetErrorMessage())
	}

	return responseData, nil
}

func GetAdvice(ID string) (string, error) {
	advice, err := fetchAdvice(ID); if err != nil {
		return "", err
	}

	return fmt.Sprintf("#%d: %s", advice.GetID(), advice.GetAdvice()), nil
}
