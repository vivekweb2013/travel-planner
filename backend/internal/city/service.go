package city

import (
	"strings"
)

type Service interface {
	GetMatchingCities(query string) []City
}

type service struct {
	cities map[string][]City
}

func NewService(cities map[string][]City) Service {
	return &service{
		cities: cities,
	}
}

func (s *service) GetMatchingCities(query string) []City {
	var match []City

	for _, cities := range s.cities {
		for _, c := range cities {
			if strings.HasPrefix(c.Name, query) {
				match = append(match, c)
			}
			if len(match) == 5 {
				break
			}
		}
	}
	return match
}
