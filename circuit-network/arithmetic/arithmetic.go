package arithmetic

import (
	"sync"
	"time"

	"github.com/google/uuid"
	circuitNetwork "github.com/thevan4/go-factorio-smart-laser-defense/circuit-network"
)

const energyCost = 1 //kW

//Combinator ...
type Combinator struct {
	sync.Mutex
	ID                    uuid.UUID
	GlobalTicker          chan *sync.WaitGroup
	EnergyWire            chan circuitNetwork.EnergyWire
	OutName               string
	GreenWire, RedWire    chan circuitNetwork.Signal
	Out                   []chan circuitNetwork.Signal
	MathematicalOperation func(inFirst int64, inSecond int64, outName string) (output circuitNetwork.Signal)
}

//NewCombinator ...
func (c *Combinator) NewCombinator(
	globalTicker chan *sync.WaitGroup,
	ev chan circuitNetwork.EnergyWire,
	outName string,
	greenWire, redWire chan circuitNetwork.Signal,
	out []chan circuitNetwork.Signal,
	mathematicalOperation func(inFirst int64, inSecond int64, outName string) (output circuitNetwork.Signal),
) *Combinator {
	return &Combinator{
		GlobalTicker:          globalTicker,
		ID:                    uuid.New(),
		EnergyWire:            ev,
		OutName:               outName,
		GreenWire:             greenWire,
		RedWire:               redWire,
		Out:                   out,
		MathematicalOperation: mathematicalOperation,
	}
}

func (c *Combinator) Work() {
	for {
		t := <-c.GlobalTicker
		t.Done()

		c.Lock()
		enoughEnergy := c.drainEnergy()
		if enoughEnergy {
			c.combine()
		} //else error?
		c.Unlock()
		time.Sleep(circuitNetwork.TickTime)
	}
}

func (c *Combinator) drainEnergy() (enoughEnergy bool) {
	//wait energy
	e := <-c.EnergyWire
	if e.Charge >= energyCost {
		enoughEnergy = true
	}
	return enoughEnergy
}

func (c *Combinator) combine() {
	g := <-c.GreenWire
	r := <-c.RedWire
	//TODO: debug log names?
	result := c.MathematicalOperation(g.Value, r.Value, c.OutName)
	for _, o := range c.Out {
		o <- result
	}
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
