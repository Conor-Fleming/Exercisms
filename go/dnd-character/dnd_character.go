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
	return int(10 + float64((score-10)/2))
}

// Ability uses randomness to generate the score for an ability
func Ability() int {
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

// GenerateCharacter creates a new Character with random scores for abilities
func GenerateCharacter() Character {
	character := Character{
		Strength:     Ability(),
		Dexterity:    Ability(),
		Constitution: Ability(),
		Intelligence: Ability(),
		Wisdom:       Ability(),
		Charisma:     Ability(),
	}

	character.Hitpoints = Modifier(character.Constitution)
	return character
}
