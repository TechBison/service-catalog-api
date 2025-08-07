package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/shreeshg/service-catalog-api/internal"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	db := internal.InitDB()
	internal.SeedData(db)
	internal.LoadCache(db)

	r := gin.Default()
	r.GET("/services", internal.GetServices)
	r.GET("/services/:id", internal.GetServiceByID)
	r.GET("/services/:id/versions", internal.GetServiceVersions)
	return r
}

func TestGetServices(t *testing.T) {
	r := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/services?search=auth&sort=name&limit=1&offset=0", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var result []internal.Service
	err := json.Unmarshal(w.Body.Bytes(), &result)
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, len(result), 1)
	assert.Contains(t, result[0].Name, "Auth")
}

func TestGetServiceByID(t *testing.T) {
	r := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/services/1", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var s internal.Service
	err := json.Unmarshal(w.Body.Bytes(), &s)
	assert.NoError(t, err)
	assert.Equal(t, 1, s.ID)
}

func TestGetServiceVersions(t *testing.T) {
	r := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/services/1/versions", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)

	var versions []internal.Version
	err := json.Unmarshal(w.Body.Bytes(), &versions)
	assert.NoError(t, err)
	assert.GreaterOrEqual(t, len(versions), 1)
}

func TestServiceNotFound(t *testing.T) {
	r := setupRouter()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/services/9999", nil)
	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNotFound, w.Code)
}
