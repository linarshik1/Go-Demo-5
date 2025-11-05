package weather_test

import (
	"demo/weather/geo"
	"demo/weather/weather"
	"strings"
	"testing"
)

func TestGetWeather(t *testing.T) {
	expected := "Ufa"
	GeoData := geo.GeoData{
		City: expected,
	}
	format := 4
	result, err := weather.GetWeather(GeoData, format)
	if err != nil {
		t.Errorf("Пришла ошибка %v", err)
	}
	if strings.Contains(result, expected) {
		t.Errorf("Ожидалось %v, получение %v", expected, result)
	}
}

var testCases = []struct {
	name   string
	format int
}{
	{name: "Big format", format: 147},
	{name: "0 format", format: 0},
	{name: "Minus format", format: -1},
}

func TestGetWeatherWrongFormat(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			expected := "Moscow"
			GeoData := geo.GeoData{
				City: expected,
			}
			_, err := weather.GetWeather(GeoData, tc.format)
			if err != weather.ErrWrongFormat {
				t.Errorf("Ожидалось %v, получение %v", weather.ErrWrongFormat, err)
			}
		})
	}

}
