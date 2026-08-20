package purchase
import "strings"
import "fmt"
// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
    kind = strings.ToUpper(kind)
	if kind!= ""{
        if kind=="CAR" || kind =="TRUCK" {
            return true
        } else if kind == "BIKE"{
            return false
        }
        
    }
    return false 
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	if option1 < option2 {
        return option1 + " is clearly the better choice."
    }else{
        return option2 + " is clearly the better choice."
    }
    
}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	if age < 3{
        return (originalPrice * 80  ) / 100
    }
    if age >= 3 &&  age < 10{

        return (originalPrice * 70  ) / 100

    }
    if age >= 10 {
         
        fmt.Printf("paso aqui 10 %g\n", (originalPrice * 50  ) / 100)
        return (originalPrice * 50  ) / 100
    }
    return 0
    
}
