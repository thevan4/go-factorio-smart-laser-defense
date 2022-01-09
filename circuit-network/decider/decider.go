package decider

import (
	"log"
	"sync"
	"time"

	"github.com/google/uuid"
	circuitNetwork "github.com/thevan4/go-factorio-smart-laser-defense/circuit-network"
)

//Decider ...
type Decider struct {
	sync.Mutex
	ID                uuid.UUID
	OutName           string
	InFirst, InSecond chan circuitNetwork.Signal
	Out               chan circuitNetwork.Signal
	LogicalOperation  circuitNetwork.LogicalOperation
	OutputType        circuitNetwork.CompareType
}

//NewDecider ...
func NewDecider(
	outName string,
	inFirst, inSecond chan circuitNetwork.Signal,
	logicalOperation circuitNetwork.LogicalOperation,
	outputType circuitNetwork.CompareType,
) *Decider {
	return &Decider{
		ID:               uuid.New(),
		OutName:          outName,
		InFirst:          inFirst,
		InSecond:         inSecond,
		Out:              make(chan circuitNetwork.Signal, 3),
		LogicalOperation: logicalOperation,
		OutputType:       outputType,
	}
}

func (d *Decider) Work() {
	for {
		d.decide()
		time.Sleep(circuitNetwork.TickTime)
	}
}

func (d *Decider) decide() {
	d.Lock()
	defer d.Unlock()
	f := <-d.InFirst
	s := <-d.InSecond
	//TODO: debug log names?
	loResult := d.LogicalOperation(f.Value, s.Value)
	switch d.OutputType {
	case circuitNetwork.BoolCompareType:
		d.Out <- formOutputFromBool(loResult, d.OutName)
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
