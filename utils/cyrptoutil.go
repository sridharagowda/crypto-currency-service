package utils

import (
	"../models"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strings"
)

func ConstructAggregateResponse(symbolresultMap map[string]interface{}, tickerresultMap map[string]interface{}, currencyresultMap map[string]interface{}) *models.Response {
	currencyResponse := new(models.Response)
	currencyResponse.Id = fmt.Sprintf("%v", currencyresultMap["id"])
	currencyResponse.FullName = fmt.Sprintf("%v", currencyresultMap["fullName"])
	currencyResponse.Ask = fmt.Sprintf("%v", tickerresultMap["ask"])
	currencyResponse.Bid = fmt.Sprintf("%v", tickerresultMap["bid"])
	currencyResponse.Last = fmt.Sprintf("%v", tickerresultMap["last"])
	currencyResponse.Open = fmt.Sprintf("%v", tickerresultMap["open"])
	currencyResponse.Low = fmt.Sprintf("%v", tickerresultMap["low"])
	currencyResponse.High = fmt.Sprintf("%v", tickerresultMap["high"])
	currencyResponse.FeeCurrency = fmt.Sprintf("%v", symbolresultMap["feeCurrency"])

	return currencyResponse
}

func ConstructAggregateResponseForAll(symbolresultsMap map[string]map[string]interface{}, tickerresultsMap map[string]map[string]interface{},
	currencyresultsMap map[string]map[string]interface{}, signalMap map[string]interface{}) []models.Response {

	var currencyResponses []models.Response
	symbolsString := fmt.Sprintf("%v", signalMap["symbols"])

	maxlen := len(symbolsString) - 1
	symbolsString = symbolsString[1:maxlen]
	symbols := strings.Split(symbolsString, " ")

	for i := 0; i < len(symbols); i++ {
		symbolMap := symbolresultsMap[symbols[i]]
		tickerMap := tickerresultsMap[symbols[i]]
		currencyMap := currencyresultsMap[fmt.Sprintf("%v", symbolMap["baseCurrency"])]

		currencyResponse := ConstructAggregateResponse(symbolMap, tickerMap, currencyMap)
		currencyResponses = append(currencyResponses, *currencyResponse)
	}

	return currencyResponses
}

func ConstructJsonResponse(response interface{}) []byte {
	responseJson, err := json.Marshal(response)
	if err != nil {
		panic(err)
	}
	return responseJson
}

func ReadFromConfigFile() []byte {
	workingdir, err := os.Getwd()
	envName := "dev" // Default to dev
	if os.Getenv("ENV") != "" {
		envName = os.Getenv("ENV")
	}
	filename := fmt.Sprintf("crypto-currency-%s.json", envName)
	signalsFromConfig, err := ioutil.ReadFile(path.Join(workingdir, filename))
	if err != nil {
		errorMessage := "could not find config file"
		log.Println(errorMessage)
		return []byte(errorMessage)
		//panic(err)
	}
	return signalsFromConfig
}

func ConstructErrorResponse(statusCode int, message string) interface{} {
	errorResponse := new(models.ErrorResponse)
	errorResponse.StatusCode = statusCode
	if statusCode == 404 {
		message = "URL not found, pls correct"
	} else if statusCode == 500 {
		message = "Service not found or not responsive"
	} else if statusCode == 400 {
		message = "Bad request, symbol not found"
	}
	errorResponse.Message = message

	return errorResponse
}
