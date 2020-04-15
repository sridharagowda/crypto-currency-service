// +build unit

package service_test

import (
	"../service"
	"testing"
)

func TestGetCryptoCurrencyBySymbol(t *testing.T) {
	responseJson := service.GetCryptoCurrencyBySymbol("ETHBTC")
	if responseJson == nil {
		t.Errorf("handler returned unexpected body: got %v ", responseJson)
	}
}

func TestGetCryptoCurrencyAll(t *testing.T) {
	responseJson := service.GetCryptoCurrencyBySymbol("all")
	if responseJson == nil {
		t.Errorf("handler returned unexpected body: got %v ", responseJson)
	}
}
