package planner

import (
	"math"
	"strings"

	"golang.org/x/exp/slices"
)

type Service interface {
	GetMatchingCities(query string) []City
	GetTravelPlan(startCity City) []Route
}

type service struct {
	continents map[string][]City
}

func NewService(continents map[string][]City) Service {
	return &service{
		continents: continents,
	}
}

func (s *service) GetMatchingCities(query string) []City {
	match := make([]City, 0)

	for _, q := range s.continents {
		for _, c := range q {
			if strings.HasPrefix(strings.ToLower(c.Name), strings.ToLower(query)) {
				match = append(match, c)
			}
			if len(match) == 5 {
				break
			}
		}
	}
	return match
}

func (s *service) GetTravelPlan(startCity City) []Route {
	continentsToExclude := []string{startCity.ContID}
	initialRoute := Route{
		Order:  1,
		Source: &startCity,
	}
	travelPlan := []Route{initialRoute}
	return plan(s.continents, continentsToExclude, travelPlan)
}

func plan(continents map[string][]City, continentsToExclude []string, travelPlan []Route) []Route {
	lastRoute := &travelPlan[len(travelPlan)-1]
	sourceCity := lastRoute.Source
	for cont, cities := range continents {
		if slices.Contains(continentsToExclude, cont) {
			continue
		}
		for _, city := range cities {
			// calculate distance of source city with current city
			d := distance(sourceCity.Location.Lat, sourceCity.Location.Lon, city.Location.Lat, city.Location.Lon)
			if lastRoute.Dest == nil || d < *lastRoute.Distance {
				lastRoute.Dest = &city
				lastRoute.Distance = &d
			}
		}
	}
	if len(continentsToExclude) == (len(continents) - 1) {
		// return to the home city
		lastCity := travelPlan[len(travelPlan)-1].Dest
		homeCity := travelPlan[0].Source
		d := distance(homeCity.Location.Lat, homeCity.Location.Lon, lastCity.Location.Lat, lastCity.Location.Lon)
		return append(travelPlan, Route{Order: lastRoute.Order + 1, Source: lastCity, Dest: homeCity, Distance: &d})
	}
	return plan(continents, append(continentsToExclude, lastRoute.Dest.ContID), append(travelPlan, Route{Order: lastRoute.Order + 1, Source: lastRoute.Dest}))
}

func distance(lat1 float64, lon1 float64, lat2 float64, lon2 float64) float64 {
	var R = float64(6371)           // Radius of the earth in km
	var dLat = deg2rad(lat2 - lat1) // deg2rad below
	var dLon = deg2rad(lon2 - lon1)
	var a = math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(deg2rad(lat1))*math.Cos(deg2rad(lat2))*
			math.Sin(dLon/2)*math.Sin(dLon/2)

	var c = 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	var d = R * c // Distance in km
	return d
}

func deg2rad(deg float64) float64 {
	return deg * (math.Pi / 180)
}
