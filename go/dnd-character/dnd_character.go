package dndcharacter

import (
	"math/rand"
)

type Character struct {
	Strength     int
	Dexterity    int
	Constitution int
	Intelligence int
	Wisdom       int
	Charisma     int
	Hitpoints    int
}

// Modifier calculates the ability modifier for a given ability score
func Modifier(score int) int {
	return 0
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
	return 0
}

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	character := Character{
		Strength:     GenerateSum(),
		Dexterity:    GenerateSum(),
		Constitution: GenerateSum(),
		Intelligence: GenerateSum(),
		Wisdom:       GenerateSum(),
		Charisma:     GenerateSum(),
	}

	character.Hitpoints = int(10 + float64((character.Constitution-10)/2))
	return character
}

func GenerateSum() int {
	rollValues := make([]int, 4)

	//4 rolls of the dice
	rollValues = append(rollValues, rand.Intn(6))
	rollValues = append(rollValues, rand.Intn(6))
	rollValues = append(rollValues, rand.Intn(6))
	rollValues = append(rollValues, rand.Intn(6))

	min := 6 //highest possible dice roll
	indextoRemove := 0
	for i, v := range rollValues {
		if v <= min {
			indextoRemove = i
		}
	}

	rollValues = append(rollValues[:indextoRemove], rollValues[indextoRemove+1:]...) //removing minumum value

	score := 0
	for _, v := range rollValues {
		score = score + v
	}

	return score
}
