package pkg

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

type Provider struct {
	client  http.Client
	fullURL string
}

const (
	PIRATE_WEATHER_URL = "https://api.pirateweather.net"
	FORECAST_ENDPOINT  = "/forecast/%s"
	HTTP_TIMEOUT       = 10
)

func (p *Provider) NewProvider(apikey string) {
	p.client = http.Client{Timeout: time.Duration(HTTP_TIMEOUT) * time.Second}
	endpoint := fmt.Sprintf(FORECAST_ENDPOINT, apikey)
	p.fullURL = fmt.Sprintf("%s%s", PIRATE_WEATHER_URL, endpoint)
}

type Request struct {
	Provider *Provider
	Unit     string
	Lat      float64
	Lon      float64

	Data  *Data
	Error error
}

func (r *Request) Get() {
	// Create URL with specific LatLng
	url := fmt.Sprintf("%s/%v,%v", r.Provider.fullURL, r.Lat, r.Lon)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		r.Error = err
		return
	}
	// Add query parameters
	q := req.URL.Query()
	q.Add("units", r.Unit)
	req.URL.RawQuery = q.Encode()

	// Make request
	resp, err := r.Provider.client.Do(req)
	if err != nil {
		r.Error = err
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		r.Error = errors.New("Received a non-200 response")
		return
	}

	// Parse response
	var data Data
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		r.Error = err
		return
	}
	r.Data = &data
}
