package city

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
