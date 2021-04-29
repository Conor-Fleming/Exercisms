//Raindrops is a package that allows you appreciate nature and the little things
package raindrops

import "fmt"

//Convert takes numbers and creates music. In other words:
//if a number is divisible by 3, 5, or 7 a corresponding string is returned.
func Convert(num int) string {
	result := ""
	if num%3 != 0 && num%5 != 0 && num%7 != 0 {
		result = fmt.Sprintf("%d", num)
	}
	if num%3 == 0 {
		result = result + "Pling"
	}
	if num%5 == 0 {
		result = result + "Plang"
	}
	if num%7 == 0 {
		result = result + "Plong"
	}
	return result
}
