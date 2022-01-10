package ticker

import (
	"errors"
	"sync"
	"time"
)

var singleTicker *GlobalTicker
var lock = &sync.Mutex{}

//GlobalTicker ...
type GlobalTicker struct {
	Ticker      *time.Ticker
	WG          *sync.WaitGroup
	TickerChans []chan *sync.WaitGroup
}

//NewGlobalTicker ...
func NewGlobalTicker(tickTime time.Duration) (*GlobalTicker, error) {
	lock.Lock()
	defer lock.Unlock()
	if singleTicker != nil {
		return nil, errors.New("globalTicker already exist")
	}
	ticker := time.NewTicker(tickTime)
	wg := new(sync.WaitGroup)
	return &GlobalTicker{
		Ticker: ticker,
		WG:     wg,
	}, nil
}

func (t *GlobalTicker) NewTickChan() chan *sync.WaitGroup {
	newChan := make(chan *sync.WaitGroup, 3)
	t.TickerChans = append(t.TickerChans, newChan)
	return newChan
}

func (t *GlobalTicker) Work() {
	for {
		<-t.Ticker.C
		t.SendTicks()
	}
}

func (t *GlobalTicker) SendTicks() {
	for _, tickChan := range t.TickerChans {
		t.WG.Add(1)
		tickChan <- t.WG
	}
	t.WG.Wait()
}
