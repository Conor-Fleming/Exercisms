// Package twofer implements functions that allow you to share.
// https://golang.org/doc/effective_go.html#commentary
//
package twofer

// ShareWith takes a name and gives one to that person, but always two for me
func ShareWith(name string) string {
	if name != "" {
		name = "One for " + name + ", one for me."
	} else {
		name = "One for you, one for me."
	}
	return name
}
