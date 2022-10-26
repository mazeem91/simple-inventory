package main

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/mazeem91/trackman-poc/application/routers"
	"github.com/mazeem91/trackman-poc/config"
	"github.com/mazeem91/trackman-poc/infrastructure/database"
	"github.com/mazeem91/trackman-poc/infrastructure/logger"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func performRequest(r http.Handler, method, path string, body io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, body)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
func TestAddLocation(t *testing.T) {
	// Build our expected body
	body := gin.H{
		"name": "loc1",
		"area": gin.H{
			"id":   "1",
			"name": "area1",
		},
	}
	//set timezone
	viper.SetDefault("SERVER_TIMEZONE", "UTC")
	loc, _ := time.LoadLocation(viper.GetString("SERVER_TIMEZONE"))
	time.Local = loc

	if err := config.SetupConfig(); err != nil {
		logger.Fatalf("config SetupConfig() error: %s", err)
	}
	masterDSN, replicaDSN := config.DbConfiguration()

	if err := database.DbConnection(masterDSN, replicaDSN); err != nil {
		logger.Fatalf("database DbConnection error: %s", err)
	}

	// Grab our router
	router := routers.SetupRoute()
	// Perform a POST request with that handler.
	jsonBody := []byte(`{"name": "loc1", "area":"area1"}`)
	bodyReader := bytes.NewReader(jsonBody)
	w := performRequest(router, "POST", "/locations", bodyReader)
	// Assert we encoded correctly,
	// the request gives a 200
	assert.Equal(t, http.StatusCreated, w.Code)
	// Convert the JSON response to a map
	var response map[string]string
	json.Unmarshal([]byte(w.Body.String()), &response)
	// Grab the value & whether or not it exists
	value, exists := response["name"]
	// Make some assertions on the correctness of the response.
	assert.True(t, exists)
	assert.Equal(t, body["name"], value)
}
