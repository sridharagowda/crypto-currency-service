package validation

import (
	"../models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

const symbolurl = "https://api.hitbtc.com/api/2/public/symbol/"

//global variable
var symbolMapCache = map[string]interface{}{}

/**
 * Suppose to use websocket, since i dont have right URL, using http call to backend service (https://api.hitbtc.com/api/2/public/symbol/)
 * and caching the request if it is valid if it is invalid then return error.
 */
func ValidateSymbol(symbol string) (bool, map[string]interface{}) {
	if symbolMapCache[symbol] == nil {
		symbolMap := externalrestcall(symbolurl + symbol)
		if symbolMap["id"] != nil {
			symbolMapCache[symbol] = symbolMap["id"]
			return true, nil
		} else {
			return false, symbolMap
		}
	}
	return true, nil
}

/**
 * Validate symbol exists as URL parameter, if not return error
 */
func ValidateRequest(request *http.Request) (string, string) {
	errorMessage := ""
	params := mux.Vars(request)
	symbol := params["symbol"]
	if len(symbol) < 1 {
		errorMessage := "Url Param 'symbol' is missing"
		log.Println(errorMessage)
		return symbol, errorMessage
	}
	return symbol, errorMessage
}

/**
 * Validate Rest response
 */
func Validateresponse(response *http.Response) map[string]interface{} {
	result := make(map[string]interface{})
	if response.StatusCode != http.StatusOK {
		result["statusCode"] = response.StatusCode
		result["message"] = response.Status

		return result
	}
	return nil
}

/**
 * Checking error response from downstream api's
 */
func CheckForErrorResponse(response map[string]interface{}) interface{} {
	if response["statusCode"] != nil {
		errorResponse := new(models.ErrorResponse)
		statusCode, _ := response["statusCode"].(int)
		errorResponse.StatusCode = statusCode
		errorResponse.Message = fmt.Sprintf("%v", response["message"])

		return errorResponse
	}
	return nil
}

/**
 * Checking error response from downstream api's
 */
func CheckForErrorResponseForAll(response map[string]map[string]interface{}) interface{} {
	if response["statusCode"] != nil {
		errorResponse := new(models.ErrorResponse)
		stringStatus := fmt.Sprintf("%v", response["statusCode"])
		errorResponse.StatusCode, _ = strconv.Atoi(stringStatus)
		errorResponse.Message = fmt.Sprintf("%v", response["message"])

		return errorResponse
	}
	return nil
}

func externalrestcall(url string) map[string]interface{} {
	var restResponse map[string]interface{}

	response, err := http.Get(url)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
	} else {

		result := Validateresponse(response)
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
