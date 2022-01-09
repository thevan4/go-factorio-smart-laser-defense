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
	Ticker            int8 //59 ticks and 1 sleep
	ID                uuid.UUID
	InFirst, InSecond chan int64
	Out               chan int64
	LogicalOperation  circuitNetwork.LogicalOperation
	OutputType        circuitNetwork.CompareType
}

//Decide ...
func (d *Decider) Decide() {
	d.Lock()
	defer d.Unlock()
	d.Ticker++
	if d.Ticker == circuitNetwork.MagicSleepTick {
		d.Ticker = 0
		time.Sleep(circuitNetwork.TickTime)
	}
	f := <-d.InFirst
	s := <-d.InSecond
	loResult := d.LogicalOperation(f, s)
	switch d.OutputType {
	case circuitNetwork.BoolCompareType:
		d.Out <- formIntOutput(loResult)
	default:
		log.Fatalf("decider %v got unknown output form type %v",
			d.ID, d.OutputType)
	}
}

func formIntOutput(compareResult bool) (out int64) {
	if compareResult {
		out = 1
	}
	return out
}
