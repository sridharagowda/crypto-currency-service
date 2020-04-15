// +build unit

package controller_test

import (
	"../controller"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGetCryptoCurrencyBySymbolSuccess(t *testing.T) {
	req, _ := http.NewRequest("GET", "/v1.0/api/currency/ETHBTC", nil)
	res := httptest.NewRecorder()

	//Hack to try to fake gorilla/mux vars
	vars := map[string]string{
		"symbol": "ETHBTC",
	}

	// CHANGE THIS LINE!!!
	req = mux.SetURLVars(req, vars)

	controller.GetCryptoCurrency(res, req)

	expected := "{\"id\":\"ETH\",\"fullName\":\"Ethereum\",\"ask\":\"0.022989\",\"bid\":\"0.022986\",\"last\":\"0.022982\",\"open\":\"0.022781\",\"low\":\"0.022740\",\"high\":\"0.023197\",\"feeCurrency\":\"BTC\"}"

	if res.Code != http.StatusOK || !strings.Contains(res.Body.String(), "id") {
		t.Errorf("handler returned unexpected body: got %v want %v", res.Body.String(), expected)
	}
}

func TestGetCryptoCurrencyBySymbolBadRequest(t *testing.T) {
	req, _ := http.NewRequest("GET", "/v1.0/api/currency/BLAH", nil)
	res := httptest.NewRecorder()

	//Hack to try to fake gorilla/mux vars
	vars := map[string]string{
		"symbol": "BLAH",
	}

	// CHANGE THIS LINE!!!
	req = mux.SetURLVars(req, vars)

	controller.GetCryptoCurrency(res, req)

	if res.Code != http.StatusBadRequest {
		t.Errorf("handler returned unexpected body: got %v ", res.Body.String())
	}
}
