package main

import "strconv"

type Observer interface {
	// v1 push 所有数据都传递给观察者
	// update(temp, humidity, pressure float32)

	// v2 pull 只传递一个主题对象，观察者可以自己获取需要的数据
	update()
}

type DisplayElement interface {
	display()
}

type CurrentCondDisplay struct {
	temperature float32
	humidity float32
	wd *WeatherData
}

func NewCurrentCondDisplay(wd *WeatherData) *CurrentCondDisplay {
	result := &CurrentCondDisplay{
		wd: wd,
	}
	wd.registerObserver("current", result)
	return result
}

func (c *CurrentCondDisplay) update() {
	c.temperature = c.wd.GetTemperature()
	c.humidity = c.wd.GetHumidity()
	c.display()
}

func (c *CurrentCondDisplay) display() {
	println("Current conditions: ", c.temperature, "F degrees and ", c.humidity, "% humidity")
}

type StatisticsDisplay struct {
	maxTemp float32
	minTemp float32
	tempSum float32
	numReadings int;
	wd *WeatherData;
}

func NewStatisticsDisplay(wd *WeatherData) *StatisticsDisplay {
	result := &StatisticsDisplay{
		minTemp: 200,
		wd: wd,
	}
	wd.registerObserver("statistics", result)
	return result
}

func (s *StatisticsDisplay) update() {
	temp := s.wd.GetTemperature()
	s.tempSum += temp
	s.numReadings++
	if temp > s.maxTemp {
		s.maxTemp = temp
	}

	if temp < s.minTemp {
		s.minTemp = temp
	}
	s.display()
}

func (s *StatisticsDisplay) display() {
	println("Avg/Max/Min temperature = " + strconv.FormatFloat(float64(s.tempSum)/float64(s.numReadings), 'f', 2, 32) + "/" + strconv.FormatFloat(float64(s.maxTemp), 'f', 2, 32) + "/" + strconv.FormatFloat(float64(s.minTemp), 'f', 2, 32))
}

type ForecastDisplay struct {
	currentPressure float32
	lastPressure float32
	wd *WeatherData
}

func NewForecastDisplay(wd *WeatherData) *ForecastDisplay {
	result := &ForecastDisplay{
		currentPressure: 29.92,
		wd: wd,
	}
	wd.registerObserver("forecast", result)
	return result
}

func (f *ForecastDisplay) update() {
	f.lastPressure = f.currentPressure
	f.currentPressure = f.wd.GetPressure()
	f.display()
}

func (f *ForecastDisplay) display() {
	print("Forecast: ")
	if f.currentPressure > f.lastPressure {
		println("Improving weather on the way!")
	} else if f.currentPressure == f.lastPressure {
		println("More of the same")
	} else if f.currentPressure < f.lastPressure {
		println("Watch out for cooler, rainy weather")
	}
}
