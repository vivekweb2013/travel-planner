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

// GetTravelPlan creates a travel plan that starts and ends to the provided city.
// It returns the travel plan as a http response.
func (c *CityHandler) GetTravelPlan(g *gin.Context) {
	var cityPayload city.City
	g.BindJSON(&cityPayload)
	logrus.Infof("%+v", cityPayload)
	err := cityPayload.Validate()
	if err != nil {
		abortRequestWithError(g, NewAppError(ErrorCodeValidationFailed, err.Error()))
		return
	}

	logrus.WithField("start city", cityPayload.Name).Info("request to make travel plan started")

	plan := c.cityService.GetTravelPlan(cityPayload)
	g.JSON(http.StatusOK, plan)

	logrus.WithField("start city", cityPayload.Name).Info("request to make travel plan successful")
}
