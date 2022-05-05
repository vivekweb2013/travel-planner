package planner

import validation "github.com/go-ozzo/ozzo-validation"

type Location struct {
	Lat float64 `json:"lat"`
	Lon float64 `json:"lon"`
}

type City struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	Location    Location `json:"location"`
	CountryName string   `json:"countryName"`
	ContID      string   `json:"contId"`
}

type Route struct {
	Order    uint     `json:"order"`
	Source   *City    `json:"source"`
	Dest     *City    `json:"dest"`
	Distance *float64 `json:"distance"`
}

func (c City) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.ID, validation.Required),
		validation.Field(&c.ContID, validation.Required),
		validation.Field(&c.Location, validation.Required),
	)
}

func (l Location) Validate() error {
	return validation.ValidateStruct(&l,
		validation.Field(&l.Lat, validation.Required),
		validation.Field(&l.Lon, validation.Required),
	)
}
