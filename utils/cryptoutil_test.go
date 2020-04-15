// +build unit

package utils_test

import (
	"../utils"
	"testing"
)

func TestConstructAggregateResponse(t *testing.T) {
	var symbolresultMap map[string]interface{}
	var tickerresultMap map[string]interface{}
	var currencyresultMap map[string]interface{}

	response := utils.ConstructAggregateResponse(symbolresultMap, tickerresultMap, currencyresultMap)

	if response == nil {
		t.Errorf("handler returned unexpected body: got %v ", response)
	}
}

func TestConstructErrorResponse(t *testing.T) {
	errorResponse1 := utils.ConstructErrorResponse(400, "Bad request")
	if errorResponse1 == nil {
		t.Errorf("handler returned unexpected body: got %v ", errorResponse1)
	}

	errorResponse2 := utils.ConstructErrorResponse(500, "server error")
	if errorResponse2 == nil {
		t.Errorf("handler returned unexpected body: got %v ", errorResponse2)
	}

	errorResponse3 := utils.ConstructErrorResponse(404, "not found")
	if errorResponse3 == nil {
		t.Errorf("handler returned unexpected body: got %v ", errorResponse3)
	}
}
