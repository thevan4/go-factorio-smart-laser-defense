package powerswitch

import (
	"log"
	"sync"
	"time"

	circuitNetwork "github.com/thevan4/go-factorio-smart-laser-defense/circuit-network"
	"github.com/thevan4/go-factorio-smart-laser-defense/defense"
	"github.com/thevan4/go-factorio-smart-laser-defense/energy"
)

//PowerSwitch ...
type PowerSwitch struct {
	sync.Mutex
	SwitchValue      int64
	In               chan circuitNetwork.Signal
	LogicalOperation circuitNetwork.LogicalOperation
	IsOn             bool

	Connections []interface{} //What connected to switch using some pole/substation
}

//NewPowerSwitch ...
func NewPowerSwitch(in chan circuitNetwork.Signal, logicalOperation circuitNetwork.LogicalOperation, connections []interface{}) *PowerSwitch {
	return &PowerSwitch{
		SwitchValue:      0, //FIXME:
		In:               in,
		LogicalOperation: logicalOperation,
		IsOn:             false,
		Connections:      connections,
	}
}

func (ps *PowerSwitch) Work() {
	for {
		ps.switchLogic()
		time.Sleep(circuitNetwork.TickTime)
	}
}

func (ps *PowerSwitch) switchLogic() {
	ps.Lock()
	defer ps.Unlock()
	in := <-ps.In
	isOn := ps.LogicalOperation(in.Value, ps.SwitchValue)
	ps.IsOn = isOn
	for connection := range ps.Connections {
		powerSwitchForConnections(connection, isOn)
	}
}

func powerSwitchForConnections(connection interface{}, isOn bool) {
	switch connect := connection.(type) {
	//TODO base?
	case *defense.LaserTurret:
		connect.EnableOrDisable(isOn)
	case *energy.Accumulator:
		connect.Work(isOn)
	default:
		log.Fatalf("powerSwitchForConnections unsupported type %T", connect)
	}
}
