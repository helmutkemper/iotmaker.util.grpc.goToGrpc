package server

import (
	"sync"
	"time"
)

type PullStatusList struct {
	m   map[string]BuildOrPullLog
	mux sync.Mutex
}

func (el *PullStatusList) Verify(key string) (found bool) {
	el.mux.Lock()
	defer el.mux.Unlock()

	if len(el.m) == 0 {
		return
	}

	_, found = el.m[key]
	return
}

func (el *PullStatusList) Set(key string, value BuildOrPullLog) {
	el.mux.Lock()
	defer el.mux.Unlock()

	if len(el.m) == 0 {
		el.m = make(map[string]BuildOrPullLog)
	}

	el.m[key] = value
}

func (el *PullStatusList) Get(key string) (value BuildOrPullLog, found bool) {
	el.mux.Lock()
	defer el.mux.Unlock()

	value, found = el.m[key]
	return
}

func (el *PullStatusList) TickerDeleteOldData() {
	el.mux.Lock()
	defer el.mux.Unlock()

	for k := range el.m {
		start := el.m[k].Start
		if time.Since(start) >= 2*time.Second*60*60 {
			delete(el.m, k)
		}
	}
}
