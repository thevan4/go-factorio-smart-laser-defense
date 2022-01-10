package defense

import (
	"sync"
)

const (
	SleepDrain uint16 = 24   //kW
	FireDrain  uint16 = 1200 //kW
)

//LaserTurret ...
type LaserTurret struct {
	sync.Mutex
	GlobalTicker chan *sync.WaitGroup
	IsOn         bool
	IsFire       bool
}

//NewLaserTurret ...
func NewLaserTurret(globalTicker chan *sync.WaitGroup) *LaserTurret {
	return &LaserTurret{
		GlobalTicker: globalTicker,
	}
}

//SetIsFire ...
func (l *LaserTurret) SetIsFire(isFire bool) {
	l.Lock()
	defer l.Unlock()
	l.IsFire = isFire
}

//EnableOrDisable ...
func (l *LaserTurret) EnableOrDisable(isOn bool) {
	l.Lock()
	defer l.Unlock()
	l.IsOn = isOn
}
