package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/vivekweb2013/travel-planner/internal/httpservice"
	"github.com/vivekweb2013/travel-planner/internal/planner"
)

const (
	filename = "cities.json"
)

func main() {
	cities, err := parseCities()
	if err != nil {
		logrus.Fatal("error parsing cities.json file", err)
	}

	plannerService := planner.NewService(cities)
	httpservice.Run(plannerService)
}

func parseCities() (map[string][]planner.City, error) {
	var m map[string]planner.City
	jsonFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(byteValue, &m)
	qc := make(map[string][]planner.City)
	for _, c := range m {
		if val, ok := qc[c.ContID]; ok {
			qc[c.ContID] = append(val, c)
		} else {
			qc[c.ContID] = []planner.City{c}
		}
	}

	logrus.Infof("parsed cities from %d continents", len(qc))
	return qc, nil
}
