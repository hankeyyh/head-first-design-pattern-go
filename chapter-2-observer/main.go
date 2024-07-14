package main



func main() {
	wd := NewWeatherData()
	NewCurrentCondDisplay(wd)
	NewStatisticsDisplay(wd)
	NewForecastDisplay(wd)
	wd.SetMeasurements(80, 65, 30.4)
	wd.SetMeasurements(82, 70, 29.2)
	wd.SetMeasurements(78, 90, 29.2)
}