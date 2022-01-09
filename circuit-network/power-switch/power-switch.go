package powerswitch

import (
	"sync"
	"time"

	circuitNetwork "github.com/thevan4/go-factorio-smart-laser-defense/circuit-network"
)

//PowerSwitch ...
type PowerSwitch struct {
	sync.Mutex
	Ticker           int8 //59 ticks and 1 sleep
	SwitchValue      int64
	In               chan int64
	LogicalOperation circuitNetwork.LogicalOperation
	IsOn             bool
}

func (ps *PowerSwitch) SwitchLogic() {
	ps.Lock()
	defer ps.Unlock()
	in := <-ps.In
	ps.Ticker++
	if ps.Ticker == circuitNetwork.MagicSleepTick {
		ps.Ticker = 0
		time.Sleep(circuitNetwork.TickTime)
	}
	ps.IsOn = ps.LogicalOperation(in, ps.SwitchValue)
}
