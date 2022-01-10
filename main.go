package main

import (
	"fmt"

	circuitNetwork "github.com/thevan4/go-factorio-smart-laser-defense/circuit-network"
	powerswitch "github.com/thevan4/go-factorio-smart-laser-defense/circuit-network/power-switch"
	"github.com/thevan4/go-factorio-smart-laser-defense/defense"
	"github.com/thevan4/go-factorio-smart-laser-defense/energy"
)

func main() {
	//laserDetectorEnergyWire := energy.NewDefaultEnergyWire()
	laserDetector := &defense.LaserTurret{}
	battery := energy.NewAccumulator([]interface{}{laserDetector})

	inPowerSwitchForBatterySignal := circuitNetwork.NewDefaultSignalChan()
	powerSwitchForBattery := powerswitch.NewPowerSwitch(
		inPowerSwitchForBatterySignal,
		circuitNetwork.LTE,
		[]interface{}{battery},
	)

	mainDefenceSignal := make(chan circuitNetwork.Signal, 3) //FIXME
	laserDefence1 := &defense.LaserTurret{}
	laserDefence2 := &defense.LaserTurret{}
	powerSwitchForMainDefence := powerswitch.NewPowerSwitch(
		mainDefenceSignal,
		circuitNetwork.EQ,
		[]interface{}{laserDefence1, laserDefence2},
	)

	//inFirst, inSecond := make(chan circuitNetwork.Signal, 3), make(chan circuitNetwork.Signal, 3)
	//deciderForLaserDetector := decider.NewDecider(
	//	"enemyDetected",
	//	inFirst, inSecond,
	//	circuitNetwork.EQ,
	//	circuitNetwork.BoolCompareType,
	//)

	fmt.Println(powerSwitchForBattery, powerSwitchForMainDefence, laserDetector, laserDefence1, laserDefence2, battery)
}
