package phonenumber

import (
	"fmt"
	"regexp"
)

func Number(phoneNumber string) (string, error) {
	result, err := CleanNumber(phoneNumber)
	if err != nil {
		return "", err
	}

	return result[1] + result[2] + result[3], nil
}

func AreaCode(phoneNumber string) (string, error) {
	result, err := CleanNumber(phoneNumber)
	if err != nil {
		return "", err
	}

	return result[1], nil
}

func Format(phoneNumber string) (string, error) {
	result, err := CleanNumber(phoneNumber)
	if err != nil {
		return "", err
	}

	dashAdded := result[2][2:] + "-" + result[2][:3]

	return "(" + result[1] + ") " + dashAdded, nil
}

func CleanNumber(number string) ([]string, error) {
	re := regexp.MustCompile(`(?:\D*?\(?(\d{3})\)?\D*?)?(?:(\d{3})\D*(\d{4}))`)
	matches := re.FindStringSubmatch(number)

	if len(matches) < 3 {
		return nil, fmt.Errorf("invalid phone number format")
	}

	if len (matches[1]+matches[2])

	return matches, nil
}
