package meteorology
import "fmt"
type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

// Add a String method to the TemperatureUnit type
func (temp TemperatureUnit) String() string {
    units :=[]string{"°C", "°F"}
    return units[temp]
}

type Temperature struct {
	degree int
	unit   TemperatureUnit
}

// Add a String method to the Temperature type
func (temp Temperature) String() string{
    return fmt.Sprintf("%v %v", temp.degree, temp.unit)
}
type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

// Add a String method to SpeedUnit
func (speedun SpeedUnit) String() string{
    units := []string{"km/h", "mph"}
    return units[speedun]
}
type Speed struct {
	magnitude int
	unit      SpeedUnit
}

// Add a String method to Speed
func (sp Speed) String() string{
    return fmt.Sprintf("%v %v", sp.magnitude, sp.unit)
}
type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

// Add a String method to MeteorologyData
func (meto MeteorologyData) String() string {
    return fmt.Sprintf("%v: %v %v, Wind %v at %v %v, %v%% Humidity", meto.location , meto.temperature.degree, meto.temperature.unit , meto.windDirection ,  meto.windSpeed.magnitude ,  meto.windSpeed.unit , meto.humidity)
}

