//Cars package provides functions for helping with production of cars
package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	return (float64(successRate/100) * float64(productionRate))
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	//This line giving me issues but not sure why. Will just re-calculate the Per Hour number again within this function
	//When I learn the cause of error, and address I will re-enstate the use of the function
	//perHour := CalculateWorkingCarsPerMinute(productionRate, successRate)

	perHour := (float64(successRate/100) * float64(productionRate))
	perminute := int(perHour) / 60
	return perminute
}

// CalculateCost works out the cost of producing the given number of cars
func CalculateCost(carsCount int) uint {
	indivs := carsCount % 10
	tenners := carsCount / 10
	return (uint((indivs * 10000) + (tenners * 95000)))
}
