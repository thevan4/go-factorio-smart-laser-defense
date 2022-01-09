package constant

import (
	"sync"
	"time"

	"github.com/google/uuid"
	circuitNetwork "github.com/thevan4/go-factorio-smart-laser-defense/circuit-network"
)

//Constant ...
type Constant struct {
	sync.Mutex
	ID     uuid.UUID
	Signal circuitNetwork.Signal
	Out    chan circuitNetwork.Signal
}

//NewConstant ...
func NewConstant(signal circuitNetwork.Signal) *Constant {
	return &Constant{
		ID:     uuid.New(),
		Signal: signal,
		Out:    make(chan circuitNetwork.Signal, 3),
	}
}

//Work ...
func (c *Constant) Work() {
	for {
		c.sendSignal()
		time.Sleep(circuitNetwork.TickTime)
	}
}

func (c *Constant) sendSignal() {
	c.Lock()
	defer c.Unlock()
	c.Out <- c.Signal
}
