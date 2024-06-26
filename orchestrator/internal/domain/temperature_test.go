package domain

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTemperature(t *testing.T) {
	t.Run("given a new temperature when create a new temperature should calculate kelvin temperature", func(t *testing.T) {
		city := "Sao Paulo"
		celsius := 28.5
		fahrenheit := 83.3
		kelvin := 301.5

		temp := NewTemperature(city, celsius, fahrenheit)
		assert.Equal(t, city, temp.City)
		assert.Equal(t, celsius, temp.Celsius)
		assert.Equal(t, fahrenheit, temp.Fahrenheit)
		assert.Equal(t, kelvin, temp.Kelvin)
	})
}
