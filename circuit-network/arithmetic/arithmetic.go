package arithmetic

import (
	"sync"
	"time"

	"github.com/google/uuid"
	circuitNetwork "github.com/thevan4/go-factorio-smart-laser-defense/circuit-network"
)

//Combinator ...
type Combinator struct {
	sync.Mutex
	Ticker                int8 //59 ticks and 1 sleep
	ID                    uuid.UUID
	InFirst, InSecond     chan int64
	Out                   chan int64
	MathematicalOperation func(inFirst int64, inSecond int64) (output int64)
}

func (c *Combinator) Combine() {
	c.Lock()
	defer c.Unlock()
	c.Ticker++
	if c.Ticker == circuitNetwork.MagicSleepTick {
		c.Ticker = 0
		time.Sleep(circuitNetwork.TickTime)
	}
	f := <-c.InFirst
	s := <-c.InSecond
	c.Out <- c.MathematicalOperation(f, s)
}

func Addition(f int64, s int64) (output int64) {
	return f + s
}

func Subtraction(f int64, s int64) (output int64) {
	return f - s
}

func Multiplication(f int64, s int64) (output int64) {
	return f * s
}
