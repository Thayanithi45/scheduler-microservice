package controllers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"scheduler-microservice/database"
	"scheduler-microservice/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

// Setup router for tests
func setupTestRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)

	if err := godotenv.Load("../.env"); err != nil {
		log.Println("No .env file found")
	}

	database.Init() // uses env vars
	router := gin.Default()
	routes.SetupRoutes(router)
	return router
}

// Test Get All Jobs API
func TestGetJobs(t *testing.T) {
	router := setupTestRouter()

	req, _ := http.NewRequest("GET", "/jobs/", nil)
	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var jobs []map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &jobs)

	fmt.Println("responce",jobs)
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(jobs), 1, "Expected at least 1 job in response")
}
