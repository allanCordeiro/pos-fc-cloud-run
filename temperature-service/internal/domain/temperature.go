package domain

type Temperature struct {
	City       string
	Celsius    float64
	Fahrenheit float64
	Kelvin     float64
}

func NewTemperature(city string, celsius float64, fahrenheit float64) *Temperature {
	t := &Temperature{
		City:       city,
		Celsius:    celsius,
		Fahrenheit: fahrenheit,
	}
	t.CalculateKelvinTemp()

	return t
}

func (t *Temperature) CalculateKelvinTemp() {
	t.Kelvin = t.Celsius + 273.0
}
