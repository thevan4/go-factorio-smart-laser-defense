package energy

import (
	"log"
	"time"

	circuitNetwork "github.com/thevan4/go-factorio-smart-laser-defense/circuit-network"
	"github.com/thevan4/go-factorio-smart-laser-defense/defense"
)

const (
	MaxAccumulatorCapacity = 5000 //kW
	input                  = 300  //kW
	output                 = 300  //kW
)

type Accumulator struct {
	Charge uint16 //0-5000 kJ

	Connections []interface{} //What connected to accumulator using some pole/substation
}

//NewAccumulator ...
func NewAccumulator(connections []interface{}) *Accumulator {
	return &Accumulator{
		Charge:      0,
		Connections: connections,
	}
}

func (a *Accumulator) Work(isOn bool) {
	for {
		a.charging(isOn)
		a.discharging()
		time.Sleep(circuitNetwork.TickTime)
	}
}

//Charging give 300kW to accumulator charge
func (a *Accumulator) charging(isOn bool) {
	if isOn {
		if a.Charge != MaxAccumulatorCapacity {
			a.Charge += input
			if a.Charge >= MaxAccumulatorCapacity {
				a.Charge = MaxAccumulatorCapacity
				//isCharged = true
			}
		}
		//else {
		//	isCharged = true
		//}
	}
	//return isCharged
}

//Discharging take 300kW from accumulator charge
func (a *Accumulator) discharging() (isPowerLow bool) {
	if a.Charge < output {
		a.Charge = 0
		isPowerLow = true
	} else {
		a.Charge -= output
	}

	//FIXME: somehow rework
	//if isPowerLow {
	//	for connection := range a.Connections {
	//		powerChangeForConnections(connection, isPowerLow)
	//	}
	//}
	return isPowerLow
}

//FIXME need idea how to send energy, need some feedback from connections
func (a *Accumulator) sendPowerToConnections(connection interface{}, isPowerLow bool) {
	switch connect := connection.(type) {
	case *defense.LaserTurret:
		if connect.IsOn {
			if connect.IsFire {
				if a.Charge >= defense.FireDrain {
					a.Charge -= defense.FireDrain
				} else {
					connect.IsOn = false
					a.Charge = 0
				}
				return
			}
		}
		connect.EnableOrDisable(!isPowerLow)
	default:
		log.Fatalf("powerChangeForConnections unsupported type %T", connect)
	}
}
