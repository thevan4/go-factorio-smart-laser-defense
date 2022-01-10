package decider

import (
	"log"
	"sync"

	"github.com/google/uuid"
	circuitNetwork "github.com/thevan4/go-factorio-smart-laser-defense/circuit-network"
)

const energyCost = 1 //kW

//Decider ...
type Decider struct {
	sync.Mutex
	ID                 uuid.UUID
	GlobalTicker       chan *sync.WaitGroup
	EnergyWire         chan circuitNetwork.EnergyWire
	OutName            string
	GreenWire, RedWire chan circuitNetwork.Signal
	Out                []chan circuitNetwork.Signal
	LogicalOperation   circuitNetwork.LogicalOperation
	OutputType         circuitNetwork.CompareType
}

//NewDecider ...
func NewDecider(
	globalTicker chan *sync.WaitGroup,
	energyWire chan circuitNetwork.EnergyWire,
	outName string,
	greenWire, redWire chan circuitNetwork.Signal,
	out []chan circuitNetwork.Signal,
	logicalOperation circuitNetwork.LogicalOperation,
	outputType circuitNetwork.CompareType,
) *Decider {
	return &Decider{
		ID:               uuid.New(),
		GlobalTicker:     globalTicker,
		EnergyWire:       energyWire,
		OutName:          outName,
		GreenWire:        greenWire,
		RedWire:          redWire,
		Out:              out,
		LogicalOperation: logicalOperation,
		OutputType:       outputType,
	}
}

func (d *Decider) Work() {
	for {
		t := <-d.GlobalTicker
		t.Done()

		d.Lock()
		d.decide()
		d.Unlock()
	}
}

func (d *Decider) decide() {
	g := <-d.GreenWire
	r := <-d.RedWire
	//TODO: debug log names?
	loResult := d.LogicalOperation(g.Value, r.Value)
	switch d.OutputType {
	case circuitNetwork.BoolCompareType:
		result := formOutputFromBool(loResult, d.OutName)
		for _, o := range d.Out {
			o <- result
		}
	default:
		log.Fatalf("decider %v got unknown output form type %v",
			d.ID, d.OutputType)
	}
}

func formOutputFromBool(compareResult bool, outName string) (out circuitNetwork.Signal) {
	out.Name = outName
	if compareResult {
		out.Value = 1
	}
	return out
}
