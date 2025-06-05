package two_map

import (
	"sync"
	"sync/atomic"
)

var (
	fenceAllMap = make(map[int]map[int]*int64, 50)
	frLock      sync.RWMutex
)

func Number(n int) {
	frLock.RLock()
	fenceMap, has := fenceAllMap[n]
	frLock.RUnlock()
	if !has {
		frLock.Lock()
		if _, h := fenceAllMap[n]; !h {
			var mm = make(map[int]*int64, 0)
			mm[1] = new(int64)
			fenceAllMap[n] = mm
		}
		fenceMap = fenceAllMap[n]
		frLock.Unlock()
	}

	atomic.AddInt64(fenceMap[1], 1)
}
