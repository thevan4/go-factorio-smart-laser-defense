package constant

import (
	"sync"

	"github.com/google/uuid"
	circuitNetwork "github.com/thevan4/go-factorio-smart-laser-defense/circuit-network"
)

//Constant ...
type Constant struct {
	sync.Mutex
	GlobalTicker chan *sync.WaitGroup
	ID           uuid.UUID
	Signal       circuitNetwork.Signal
	Out          []chan circuitNetwork.Signal
}

//NewConstant ...
func NewConstant(
	globalTicker chan *sync.WaitGroup,
	signal circuitNetwork.Signal,
	out []chan circuitNetwork.Signal) *Constant {
	return &Constant{
		GlobalTicker: globalTicker,
		ID:           uuid.New(),
		Signal:       signal,
		Out:          out,
	}
}

//Work ...
func (c *Constant) Work() {
	for {
		t := <-c.GlobalTicker
		t.Done()
		c.Lock()
		c.sendSignal()
		c.Unlock()
	}
}

func (c *Constant) sendSignal() {
	for _, o := range c.Out {
		o <- c.Signal
	}
}
