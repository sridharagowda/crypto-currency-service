package controller

import (
	"../service"
	"../utils"
	"../validation"
	"log"
	"net/http"
)

/**
 * This endpoint to support http://localhost:8080/v1.0/api/currency/{symbol}
 * Accept symbol from request url as param and if value is all then process all the symbol, if it is specific symbol then process that symbol
 * if it is invalid then return respective error.
 */

func GetCryptoCurrency(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	symbol, errorMessage := validation.ValidateRequest(request)
	if errorMessage != "" {
		response.WriteHeader(http.StatusBadRequest)
		response.Write([]byte(errorMessage))
		return
	}

	if symbol == "all" {
		GetAllCryptoCurrency(response, request)
		return
	} else {
		// Validate symbol from realtime market data
		isValid, errResp := validation.ValidateSymbol(symbol)

		if isValid {
			responseJson := service.GetCryptoCurrencyBySymbol(symbol)
			response.Write(responseJson)
		} else {
			log.Println(errorMessage)
			response.WriteHeader(errResp["statusCode"].(int))
			response.Write([]byte(utils.ConstructJsonResponse(utils.ConstructErrorResponse(errResp["statusCode"].(int), errResp["message"].(string)))))
		}
	}
}

/**
 * This endpoint to support http://localhost:8080/v1.0/api/currency
 * without symbol, this is not part of requirement, just added as additional use case
 */
func GetAllCryptoCurrency(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	response.WriteHeader(http.StatusOK)
	responseJson := service.GetAllCryptoCurrency()
	response.Write(responseJson)
}
