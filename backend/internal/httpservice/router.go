package httpservice

import (
	"net"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/vivekweb2013/travel-planner/internal/planner"
)

// Run starts the http server.
func Run(cityService planner.Service) error {
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	cityHandler := NewPlannerHandler(cityService)

	v1 := router.Group("api/v1")
	v1.GET("/cities", cityHandler.GetMatchingCities)
	v1.POST("/travelplan", cityHandler.GetTravelPlan)

	address := net.JoinHostPort("localhost", "8080")
	server := http.Server{
		Addr:           address,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   2 * time.Minute,
		MaxHeaderBytes: 1 << 20,
	}
	return server.ListenAndServe()
}
