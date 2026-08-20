//Package weather is used to know the wather of differnet regions.
package weather
var (
    //CurrentCondition of wather.
	CurrentCondition string 
    //CurrentLocation wather.
	CurrentLocation  string 
)
//Forecast where you pass city and condition and this know what is the whether condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
