package dndcharacter

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

	return Character{}
}
