package httpservice

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/sirupsen/logrus"
	"github.com/vivekweb2013/travel-planner/internal/city"
)

// CityHandler represents http handler for serving requests related to cities.
type CityHandler struct {
	cityService city.Service
}

// NewCityHandler creates and returns a new note handler.
func NewCityHandler(cityService city.Service) *CityHandler {
	return &CityHandler{
		cityService: cityService,
	}
}

// GetMatchingCities performs a city search against given query.
// It returns the result of search operation as a http response.
func (c *CityHandler) GetMatchingCities(g *gin.Context) {
	query := g.Query("query")
	if err := validation.Validate(query, validation.Required); err != nil {
		abortRequestWithError(g, NewAppError(ErrorCodeValidationFailed, fmt.Sprintf("query: %s", err.Error())))
		return
	}
	logrus.WithField("query", query).Info("request to search cities started")
	matchingCities := c.cityService.GetMatchingCities(query)
	g.JSON(http.StatusOK, matchingCities)
	logrus.WithField("query", query).Info("request to search cities successful")
}
