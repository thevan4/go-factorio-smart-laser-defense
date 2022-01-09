package energy

import "sync"

type Base struct{ sync.Mutex }

func (b *Base) GiveEnergy() {
	b.Lock()
	defer b.Unlock()
}
