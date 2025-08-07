package internal

import (
	"net/http"
	"sort"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetServices(c *gin.Context) {
	search := c.Query("search")
	sortKey := c.DefaultQuery("sort", "id")
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	services := GetCachedServices()
	filtered := []Service{}

	for _, s := range services {
		if search == "" || strings.Contains(strings.ToLower(s.Name), strings.ToLower(search)) ||
			strings.Contains(strings.ToLower(s.Description), strings.ToLower(search)) {
			filtered = append(filtered, s)
		}
	}

	switch sortKey {
	case "name":
		sort.Slice(filtered, func(i, j int) bool {
			return filtered[i].Name < filtered[j].Name
		})
	case "id":
		sort.Slice(filtered, func(i, j int) bool {
			return filtered[i].ID < filtered[j].ID
		})
	}

	start := offset
	if start > len(filtered) {
		start = len(filtered)
	}
	end := start + limit
	if end > len(filtered) {
		end = len(filtered)
	}

	c.JSON(http.StatusOK, filtered[start:end])
}

func GetServiceByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}
	if service, ok := GetCachedServiceByID(id); ok {
		c.JSON(http.StatusOK, service)
		return
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "service not found"})
}

func GetServiceVersions(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid ID"})
		return
	}
	if service, ok := GetCachedServiceByID(id); ok {
		c.JSON(http.StatusOK, service.Versions)
		return
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "service not found"})
}
