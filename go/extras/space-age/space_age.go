//Space package gives us a number of conversion functions for traveling time and space
package space

type Planet string

//Age function takes a value representing seconds and a planet and returns how old you would be in (Earth) years on the given planet
func Age(second float64, planet Planet) float64 {
	var age float64
	minutes := second / 60
	hours := minutes / 60
	days := hours / 24
	switch planet {
	case "Mercury":
		year := 365.25 * .2408467
		age = days / year
		break
	case "Venus":
		year := 365.25 * .61519726
		age = days / year
		break
	case "Earth":
		year := 365.25 * 1
		age = days / year
		break
	case "Mars":
		year := 365.25 * 1.8808158
		age = days / year
		break
	case "Jupiter":
		year := 365.25 * 11.862615
		age = days / year
		break
	case "Saturn":
		year := 365.25 * 29.447498
		age = days / year
		break
	case "Uranus":
		year := 365.25 * 84.016846
		age = days / year
		break
	case "Neptune":
		year := 365.25 * 164.79132
		age = days / year
		break
	}
	return age
}
