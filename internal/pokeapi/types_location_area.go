package pokeapi

// https://mholt.github.io/json-to-go/ to autogenerate a struct for a json response body
// LocationAreasResp struct to hold https://pokeapi.co/api/v2/location-area/ response.
type LocationAreasResp struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}
