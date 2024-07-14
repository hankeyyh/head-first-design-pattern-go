package main

type Subject interface {
	registerObserver(key string, o Observer)
	removeObserver(key string)
	notifyObservers()	
}

type WeatherData struct {
	observers map[string]Observer
	temperature float32
	humidity float32
	pressure float32
}

func NewWeatherData() *WeatherData {
	return &WeatherData{
		observers: make(map[string]Observer),
	}
}

func (wd *WeatherData) registerObserver(key string, o Observer) {
	wd.observers[key] = o
}

func (wd *WeatherData) removeObserver(key string) {
	delete(wd.observers, key)
}

func (wd *WeatherData) notifyObservers() {
	for _, observer := range wd.observers {
		observer.update()
	}
}

func (wd *WeatherData) SetMeasurements(temperature, humidity, pressure float32) {
	wd.temperature = temperature
	wd.humidity = humidity
	wd.pressure = pressure
	wd.measurementsChanged()
}

func (wd *WeatherData) GetTemperature() float32 {
	return wd.temperature
}

func (wd *WeatherData) GetHumidity() float32 {
	return wd.humidity
}

func (wd *WeatherData) GetPressure() float32 {
	return wd.pressure
}

func (wd *WeatherData) measurementsChanged() {
	wd.notifyObservers()
}

