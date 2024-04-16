package impl

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/allanCordeiro/pos-fc-cloud-run/internal/domain"
)

type WeatherApi struct {
	Client *http.Client
}

type Output struct {
	Location struct {
		Name           string  `json:"name"`
		Region         string  `json:"region"`
		Country        string  `json:"country"`
		Lat            float64 `json:"lat"`
		Lon            float64 `json:"lon"`
		TzID           string  `json:"tz_id"`
		LocaltimeEpoch int     `json:"localtime_epoch"`
		Localtime      string  `json:"localtime"`
	} `json:"location"`
	Current struct {
		LastUpdatedEpoch int     `json:"last_updated_epoch"`
		LastUpdated      string  `json:"last_updated"`
		TempC            float64 `json:"temp_c"`
		TempF            float64 `json:"temp_f"`
		IsDay            int     `json:"is_day"`
		Condition        struct {
			Text string `json:"text"`
			Icon string `json:"icon"`
			Code int    `json:"code"`
		} `json:"condition"`
		WindMph    float64 `json:"wind_mph"`
		WindKph    float64 `json:"wind_kph"`
		WindDegree int     `json:"wind_degree"`
		WindDir    string  `json:"wind_dir"`
		PressureMb float64 `json:"pressure_mb"`
		PressureIn float64 `json:"pressure_in"`
		PrecipMm   float64 `json:"precip_mm"`
		PrecipIn   float64 `json:"precip_in"`
		Humidity   int     `json:"humidity"`
		Cloud      int     `json:"cloud"`
		FeelslikeC float64 `json:"feelslike_c"`
		FeelslikeF float64 `json:"feelslike_f"`
		VisKm      float64 `json:"vis_km"`
		VisMiles   float64 `json:"vis_miles"`
		Uv         float64 `json:"uv"`
		GustMph    float64 `json:"gust_mph"`
		GustKph    float64 `json:"gust_kph"`
	} `json:"current"`
}

// type ErrorOutput struct {
// 	Code    int    `json:"code"`
// 	Message string `json:"message"`
// }

func NewWeatherApi(client *http.Client) *WeatherApi {
	return &WeatherApi{Client: client}
}

func (w *WeatherApi) Retrieve(ctx context.Context, city string) (*domain.Temperature, error) {
	//KLUDGE:: find a way to put this url out of this
	url := "https://api.weatherapi.com/v1/current.json?key=3a86487cb0804004a3b10835241004&q=city_name"
	url = strings.Replace(url, "city_name", city, 1)
	url = strings.Replace(url, " ", "%20", -1)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("accept", "application/json")
	res, err := w.Client.Do(req)
	if err != nil {
		return nil, err
	}

	fmt.Println(res.Body)

	if res.StatusCode != http.StatusOK {
		// body, _ := io.ReadAll(res.Body)
		// log.Println(string(body))
		// var errMessage ErrorOutput
		// _ = json.Unmarshal(body, &errMessage)
		// log.Println(errMessage)
		return nil, errors.New("http error status code: " + res.Status)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	output, err := parser(body)
	if err != nil {
		return nil, err
	}

	return domain.NewTemperature(
		output.Current.TempC,
		output.Current.TempF), nil
}

func parser(body []byte) (Output, error) {
	var data Output
	err := json.Unmarshal(body, &data)
	if err != nil {
		return Output{}, err
	}
	return data, nil
}
