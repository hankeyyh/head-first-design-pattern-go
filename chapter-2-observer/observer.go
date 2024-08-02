package main

import (
	"strconv"
)

type Observer interface {
	update()
}

type DisplayElement interface {
	display()
}

type CurrentCondDisplay struct {
	temperature float32
	humidity float32
	wd Subject
}

func NewCurrentCondDisplay(wd Subject) *CurrentCondDisplay {
	result := &CurrentCondDisplay{
		wd: wd,
	}
	wd.registerObserver("current", result)
	return result
}

func (c *CurrentCondDisplay) update() {
	if wd, ok := c.wd.(*WeatherData); ok {
		c.temperature = wd.GetTemperature()
		c.humidity = wd.GetHumidity()
		c.display()
	}
}

func (c *CurrentCondDisplay) display() {
	println("Current conditions: ", c.temperature, "F degrees and ", c.humidity, "% humidity")
}

type StatisticsDisplay struct {
	maxTemp float32
	minTemp float32
	tempSum float32
	numReadings int;
	wd Subject;
}

func NewStatisticsDisplay(wd Subject) *StatisticsDisplay {
	result := &StatisticsDisplay{
		minTemp: 200,
		wd: wd,
	}
	wd.registerObserver("statistics", result)
	return result
}

func (s *StatisticsDisplay) update() {
	if wd, ok := s.wd.(*WeatherData); ok {
		temp := wd.GetTemperature()
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
}

func (s *StatisticsDisplay) display() {
	println("Avg/Max/Min temperature = " + strconv.FormatFloat(float64(s.tempSum)/float64(s.numReadings), 'f', 2, 32) + "/" + strconv.FormatFloat(float64(s.maxTemp), 'f', 2, 32) + "/" + strconv.FormatFloat(float64(s.minTemp), 'f', 2, 32))
}

type ForecastDisplay struct {
	currentPressure float32
	lastPressure float32
	wd Subject
}

func NewForecastDisplay(wd Subject) *ForecastDisplay {
	result := &ForecastDisplay{
		currentPressure: 29.92,
		wd: wd,
	}
	wd.registerObserver("forecast", result)
	return result
}

func (f *ForecastDisplay) update() {
	if wd, ok := f.wd.(*WeatherData); ok {
		f.lastPressure = f.currentPressure
		f.currentPressure = wd.GetPressure()
		f.display()
	}
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
