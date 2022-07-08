//Package weather implements functions that allow for the viewing of weather conditions.
package weather

//CurrentCondition contains a string describing the weather right now.
var CurrentCondition string

//CurrentLocation contains a string describing where we are right now.
var CurrentLocation string

//Forecast takes two strings as parameters and returns a string telling the forecast.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
