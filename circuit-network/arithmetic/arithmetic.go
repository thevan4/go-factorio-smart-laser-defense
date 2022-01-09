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
	ID                    uuid.UUID
	OutName               string
	InFirst, InSecond     chan circuitNetwork.Signal
	Out                   chan circuitNetwork.Signal
	MathematicalOperation func(inFirst int64, inSecond int64, outName string) (output circuitNetwork.Signal)
}

func (c *Combinator) Work() {
	for {
		c.combine()
		time.Sleep(circuitNetwork.TickTime)
	}
}

func (c *Combinator) combine() {
	c.Lock()
	defer c.Unlock()
	f := <-c.InFirst
	s := <-c.InSecond
	//TODO: debug log names?
	c.Out <- c.MathematicalOperation(f.Value, s.Value, c.OutName)
}

func Addition(f int64, s int64, outName string) (output circuitNetwork.Signal) {
	return circuitNetwork.Signal{Name: outName, Value: f + s}
}

func Subtraction(f int64, s int64, outName string) (output circuitNetwork.Signal) {
	return circuitNetwork.Signal{Name: outName, Value: f - s}
}

func Multiplication(f int64, s int64, outName string) (output circuitNetwork.Signal) {
	return circuitNetwork.Signal{Name: outName, Value: f * s}
}
