package weather

import (
	"context"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Show(cityName string) *WeatherList {
	jsonData := WeatherCommandHandler(cityName)
	wl := &WeatherList{}
	err := json.Unmarshal(jsonData, wl)
	if err != nil {
		log.Println(err)
	}
	return wl
}

func WeatherCommandHandler(city string) []byte {
	httpClient := &http.Client{Timeout: 5 * time.Second}
	ctxBase := context.Background()
	ctx, cancel := context.WithTimeout(ctxBase, time.Second*5)
	defer cancel()

	req := requestBuilder(ctx, city)
	resp, err := httpClient.Do(req)
	if err != nil {
		log.Fatal(err) // actual request error
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	return respBody
}

func requestBuilder(ctx context.Context, cityName string) *http.Request {
	req, err := http.NewRequestWithContext(
		ctx,
		http.MethodGet,
		"https://api.openweathermap.org/data/2.5/forecast",
		nil,
	)
	q := url.Values{}
	//q.Add("id", "323777")
	q.Add("q", cityName)
	q.Add("cnt", "1")
	q.Add("units", "metric")
	q.Add("appid", os.Getenv("API"))
	req.URL.RawQuery = q.Encode()

	if err != nil {
		log.Fatal(err) // nil context or invalid method
	}
	return req
}
