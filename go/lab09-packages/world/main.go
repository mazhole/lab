package world

import "fmt"

const StartingRoom = "lobby"

const startingLevel = 1

//GetStartingLevel returns startingLevel value
func GetStartingLevel() int {
	return startingLevel
}

func checkStartingLevel(level int) bool {
	return level == startingLevel
}

//PrintStartRoom prints StartingRoom value
func PrintStartRoom() {
	fmt.Println("Start level is", StartingRoom)
}
