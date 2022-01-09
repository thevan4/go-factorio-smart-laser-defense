package main

import (
	"fmt"

	circuitNetwork "github.com/thevan4/go-factorio-smart-laser-defense/circuit-network"
	"github.com/thevan4/go-factorio-smart-laser-defense/circuit-network/decider"
	powerswitch "github.com/thevan4/go-factorio-smart-laser-defense/circuit-network/power-switch"
	"github.com/thevan4/go-factorio-smart-laser-defense/defense"
	"github.com/thevan4/go-factorio-smart-laser-defense/energy"
)

func main() {
	laserDetector := &defense.LaserTurret{Drain: make(chan uint16, 3)}
	battery := energy.NewAccumulator([]interface{}{laserDetector})

	inPowerSwitchForBatterySignal := make(chan circuitNetwork.Signal, 3)
	powerSwitchForBattery := powerswitch.NewPowerSwitch(
		inPowerSwitchForBatterySignal,
		circuitNetwork.LTE,
		[]interface{}{battery},
	)

	mainDefenceSignal := make(chan circuitNetwork.Signal, 3) //FIXME
	laserDefence1 := &defense.LaserTurret{Drain: make(chan uint16, 3)}
	laserDefence2 := &defense.LaserTurret{Drain: make(chan uint16, 3)}
	powerSwitchForMainDefence := powerswitch.NewPowerSwitch(
		mainDefenceSignal,
		circuitNetwork.EQ,
		[]interface{}{laserDefence1, laserDefence2},
	)

	inFirst, inSecond := make(chan circuitNetwork.Signal, 3), make(chan circuitNetwork.Signal, 3)
	deciderForLaserDetector := decider.NewDecider(
		"enemyDetected",
		inFirst, inSecond,
		circuitNetwork.EQ,
		circuitNetwork.BoolCompareType,
	)

	fmt.Println(powerSwitchForBattery, powerSwitchForMainDefence, laserDetector, laserDefence1, laserDefence2, battery, deciderForLaserDetector)
}
