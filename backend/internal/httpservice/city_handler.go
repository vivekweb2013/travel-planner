package httpservice

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	validation "github.com/go-ozzo/ozzo-validation"
	"github.com/sirupsen/logrus"
	"github.com/vivekweb2013/travel-planner/internal/planner"
)

// PlannerHandler represents http handler for serving requests related to travel planning.
type PlannerHandler struct {
	plannerService planner.Service
}

// NewPlannerHandler creates and returns a new note handler.
func NewPlannerHandler(plannerService planner.Service) *PlannerHandler {
	return &PlannerHandler{
		plannerService: plannerService,
	}
}

// GetMatchingCities performs a city search against given query.
// It returns the result of search operation as a http response.
func (c *PlannerHandler) GetMatchingCities(g *gin.Context) {
	query := g.Query("query")
	if err := validation.Validate(query, validation.Required); err != nil {
		abortRequestWithError(g, NewAppError(ErrorCodeValidationFailed, fmt.Sprintf("query: %s", err.Error())))
		return
	}
	logrus.WithField("query", query).Info("request to search cities started")
	matchingCities := c.plannerService.GetMatchingCities(query)
	g.JSON(http.StatusOK, matchingCities)
	logrus.WithField("query", query).Info("request to search cities successful")
}

// GetTravelPlan creates a travel plan that starts and ends to the provided planner.
// It returns the travel plan as a http response.
func (c *PlannerHandler) GetTravelPlan(g *gin.Context) {
	var cityPayload planner.City
	g.BindJSON(&cityPayload)
	logrus.Infof("%+v", cityPayload)
	err := cityPayload.Validate()
	if err != nil {
		abortRequestWithError(g, NewAppError(ErrorCodeValidationFailed, err.Error()))
		return
	}

	logrus.WithField("start city", cityPayload.Name).Info("request to make travel plan started")

	plan := c.plannerService.GetTravelPlan(cityPayload)
	g.JSON(http.StatusOK, plan)

	logrus.WithField("start city", cityPayload.Name).Info("request to make travel plan successful")
}
