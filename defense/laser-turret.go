package defense

import (
	"sync"
	"time"

	circuitNetwork "github.com/thevan4/go-factorio-smart-laser-defense/circuit-network"
)

const (
	sleepDrain uint16 = 24   //kW
	fireDrain  uint16 = 1200 //kW
)

//LaserTurret ...
type LaserTurret struct {
	sync.Mutex
	Drain  chan uint16
	isOn   bool
	isFire bool
}

//SetIsFire ...
func (l *LaserTurret) SetIsFire(isFire bool) {
	l.Lock()
	defer l.Unlock()
	l.isFire = isFire
}

//EnableOrDisable ...
func (l *LaserTurret) EnableOrDisable(isOn bool) {
	l.Lock()
	defer l.Unlock()
	l.isOn = isOn
}

//Work in routine. Can be turned off/on, sleep and fire
func (l *LaserTurret) Work() {
	for {
		l.Lock()

		//Turret off
		if !l.isOn {
			l.Drain <- 0
			l.Unlock()
			time.Sleep(circuitNetwork.TickTime)
			continue
		}

		//sleep
		if !l.isFire {
			l.Drain <- sleepDrain
			l.Unlock()
			time.Sleep(circuitNetwork.TickTime)
			continue
		}

		//fire
		l.Drain <- fireDrain
		l.Unlock()
		time.Sleep(circuitNetwork.TickTime)
	}
}
