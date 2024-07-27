package synch

import "sync"

type Locker interface {
	Lock()
	Unlock()
	RLock()
	RUnlock()
}

var _ Locker = (*sync.RWMutex)(nil)

// EmptyLock 空锁，不干任何事情
type EmptyLock struct {
}

func (l EmptyLock) Lock() {

}

func (l EmptyLock) Unlock() {

}

func (l EmptyLock) RLock() {

}

func (l EmptyLock) RUnlock() {

}
