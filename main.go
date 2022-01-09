package main

import (
	"fmt"

	"github.com/thevan4/go-factorio-smart-laser-defense/defense"
	"github.com/thevan4/go-factorio-smart-laser-defense/energy"
)

func main() {
	laserDetector := new(defense.LaserTurret)
	laserDefence1 := new(defense.LaserTurret)
	laserDefence2 := new(defense.LaserTurret)
	battery := new(energy.Accumulator)

	fmt.Println(laserDetector, laserDefence1, laserDefence2, battery)
}
