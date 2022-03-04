package main

import (
	"sync"
	"testing"
)

var _cs1 = 0
var _mu1 sync.Mutex

var _cs2 = 0
var _mu2 sync.RWMutex

func BenchmarkWirteSyncByMutex(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			_mu1.Lock()
			_cs1++
			_mu1.Unlock()
		}
	})
}

func BenchmarkWirteSyncByRWMutex(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			_mu2.Lock()
			_cs2++
			_mu2.Unlock()
		}
	})
}

func BenchmarkReadSyncByMutex(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			_mu1.Lock()
			_ = _cs1
			_mu1.Unlock()
		}
	})
}

func BenchmarkReadSyncByRWMutex(b *testing.B) {
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			_mu2.RLock()
			_ = _cs2
			_mu2.RUnlock()
		}
	})
}
