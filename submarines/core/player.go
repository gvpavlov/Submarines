package Submarines

// Player structure
type Player struct {
	// player x coordinate
	x int
	// player y coordinate
	y int
	// sonar width
	sonarWidth int
	// number of weapon shots remaining
	weaponShots int
	// percentage of hull integrity
	hull int
}

// Assigns initial random coordinates to the player
func (p Player) AssignCoordinates(x int, y int) {

}

// Shoots a torpedo and calculates damage based on distance from target
func (p Player) Shoot() {

}
