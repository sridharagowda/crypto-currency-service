// +build unit

package validation_test

import (
	"../validation"
	"github.com/gorilla/mux"
	"net/http"
	"testing"
)

func TestValidateSymbolSuccess(t *testing.T) {
	isValid, errResp := validation.ValidateSymbol("ETHBTC")
	if !isValid {
		t.Errorf("handler returned unexpected body: got %v ", errResp["message"])
	}
}

func TestValidateSymbolError(t *testing.T) {
	isValid, errResp := validation.ValidateSymbol("BLAH")
	if isValid && errResp == nil {
		t.Errorf("handler returned unexpected body: got %v ", errResp["message"])
	}
}

func TestValidateRequestSuccess(t *testing.T) {
	req, _ := http.NewRequest("GET", "/v1.0/api/currency/ETHBTC", nil)
	//Hack to try to fake gorilla/mux vars
	vars := map[string]string{
		"symbol": "ETHBTC",
	}
	// CHANGE THIS LINE!!!
	req = mux.SetURLVars(req, vars)

	_, errormessage := validation.ValidateRequest(req)
	if errormessage != "" {
		t.Errorf("handler returned unexpected body: got %v ", errormessage)
	}
}

func TestValidateRequestError(t *testing.T) {
	req, _ := http.NewRequest("GET", "/v1.0/api/currency/ETHBTC", nil)
	//Hack to try to fake gorilla/mux vars
	vars := map[string]string{
		"symbol": "",
	}
	// CHANGE THIS LINE!!!
	req = mux.SetURLVars(req, vars)

	_, errormessage := validation.ValidateRequest(req)
	if errormessage == "" {
		t.Errorf("handler returned unexpected body: got %v ", errormessage)
	}
}
