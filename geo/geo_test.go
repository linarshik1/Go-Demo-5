package geo_test

import (
	"demo/weather/geo"
	"testing"
)

func TestGetMyLocation(t *testing.T) {
	// Arange - подготовка, expected результат, данные функции
	city := "London"
	expected := geo.GeoData{
		City: "London",
	}
	// Act - выполняем функцию
	got, err := geo.GetMyLocation(city)
	// Asserst - проверка результата с expected
	if err != nil {
		t.Error(err)
	}
	if got.City != expected.City {
		t.Errorf("Ожидалось %v, получение %v", expected, got)
	}
}
func TestGetMyLocationNoCity(t *testing.T) {
	city := "LondonQG"
	_, err := geo.GetMyLocation(city)
	if err != geo.ErrorNoCity {
		t.Errorf("Ожидалось %v, получение %v", geo.ErrorNoCity, err)
	}

}
