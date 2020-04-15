package service

import (
	"../utils"
	"../validation"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

const symbolurl = "https://api.hitbtc.com/api/2/public/symbol/"
const tickerurl = "https://api.hitbtc.com/api/2/public/ticker/"
const currencyurl = "https://api.hitbtc.com/api/2/public/currency/"

/**
 * Make downstream API calls to fetch Currency, Symbol, Ticker by symbol
 * aggregate response and return back to client. if bad response return error
 */
func GetCryptoCurrencyBySymbol(symbol string) []byte {
	var symbolresultMap map[string]interface{}
	var tickerresultMap map[string]interface{}
	var currencyresultMap map[string]interface{}

	symbolresultMap = externalrestcall(symbolurl + symbol)
	response := validation.CheckForErrorResponse(symbolresultMap)
	if response != nil {
		return utils.ConstructJsonResponse(response)
	}

	tickerresultMap = externalrestcall(tickerurl + symbol)
	response = validation.CheckForErrorResponse(tickerresultMap)
	if response != nil {
		return utils.ConstructJsonResponse(response)
	}

	baseCurrency := fmt.Sprintf("%v", symbolresultMap["baseCurrency"])
	currencyresultMap = externalrestcall(currencyurl + baseCurrency)
	response = validation.CheckForErrorResponse(currencyresultMap)
	if response != nil {
		return utils.ConstructJsonResponse(response)
	}

	aggregateResponse := utils.ConstructAggregateResponse(symbolresultMap, tickerresultMap, currencyresultMap)
	return utils.ConstructJsonResponse(aggregateResponse)

}

/**
 * Make downstream API calls to fetch Currency, Symbol, Ticker for ALL symbols
 * aggregate response and return back to client. if bad response return error
 */
func GetAllCryptoCurrency() []byte {
	symbolresultMap := externalrestcallAllData(symbolurl)
	response := validation.CheckForErrorResponseForAll(symbolresultMap)
	if response != nil {
		return utils.ConstructJsonResponse(response)
	}

	tickerresultMap := externalrestcallAllData(tickerurl)
	response = validation.CheckForErrorResponseForAll(tickerresultMap)
	if response != nil {
		return utils.ConstructJsonResponse(response)
	}

	currencyresultMap := externalrestcallAllData(currencyurl)
	response = validation.CheckForErrorResponseForAll(currencyresultMap)
	if response != nil {
		return utils.ConstructJsonResponse(response)
	}

	signalsFromConfig := utils.ReadFromConfigFile()
	var signalMap map[string]interface{}
	json.Unmarshal([]byte(signalsFromConfig), &signalMap)

	aggregateResponse := utils.ConstructAggregateResponseForAll(symbolresultMap, tickerresultMap, currencyresultMap, signalMap)

	responseJson, err := json.Marshal(aggregateResponse)
	if err != nil {
		panic(err)
	}

	return responseJson
}

func externalrestcall(url string) map[string]interface{} {
	var restResponse map[string]interface{}

	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {

		result := validation.Validateresponse(response)
		if result != nil {
			return result
		}

		data, _ := ioutil.ReadAll(response.Body)
		//Convert response data to Map
		json.Unmarshal([]byte(data), &restResponse)

		return restResponse
	}
	return nil
}

func externalrestcallAllData(url string) map[string]map[string]interface{} {
	result := map[string]map[string]interface{}{}
	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {

		data, _ := ioutil.ReadAll(response.Body)

		//Convert array of response data to Map
		var mapSlice []map[string]interface{}
		if err := json.Unmarshal(data, &mapSlice); err != nil {
			panic(err)
		}

		for i := 0; i < len(mapSlice); i++ {
			singleMap := mapSlice[i]
			key := fmt.Sprintf("%v", singleMap["id"])
			if key == "" || key == "<nil>" { // this is for ticker map
				key = fmt.Sprintf("%v", singleMap["symbol"])
			}

			result[key] = singleMap
		}
	}
	return result
}
