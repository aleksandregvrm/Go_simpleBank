package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	utils "example.com/banking/utils/token"
)

func TestAuthMiddleware(t *testing.T) {
	testCases := []struct {
		name          string
		setupAuth     func(t *testing.T, request *http.Request, tokenMaker utils.Maker)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{}
	fmt.Println(testCases)
}
