package impl

import (
	"context"
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"net/url"

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
		TempC     float64  `json:"temp_c"`
		TempF     float64  `json:"temp_f"`
		IsDay     int      `json:"is_day"`
		Condition struct{} `json:"condition"`
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

	sanitizedCity := url.QueryEscape(city)
	//KLUDGE:: find a way to put this url out of this
	url := "https://api.weatherapi.com/v1/current.json?key=3a86487cb0804004a3b10835241004&q=" + sanitizedCity

	log.Println(url)
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("accept", "application/json")
	res, err := w.Client.Do(req)
	if err != nil {
		return nil, err
	}

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
