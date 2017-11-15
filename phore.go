package sema

import (
	"sync/atomic"
	"unsafe"

	"github.com/betawaffle/sema/internal/race"
)

const (
	Size = unsafe.Sizeof(Phore{})
)

type Phore struct {
	count uint32
	_     [60]byte
}

// Acquire blocks until a resource is available.
func (s *Phore) Acquire() {
	semacquire(&s.count, false, 1)

	if race.Enabled {
		race.Acquire(unsafe.Pointer(s))
	}
}

// AcquireSpin spins until a resource is available.
func (s *Phore) AcquireSpin() {
	for !atomic.CompareAndSwapUint32(&s.count, 1, 0) {
		// spin
	}
}

// Count returns the number of resources available.
func (s *Phore) Count() uint32 {
	return atomic.LoadUint32(&s.count)
}

// SetCount can be used to set the initial number of resources available.
// SetCount is NOT safe to call concurrently with any other methods.
func (s *Phore) SetCount(n uint32) {
	s.count = n
}

// Release unblocks a waiting Acquire or increments the available count.
func (s *Phore) Release() {
	if race.Enabled {
		race.Release(unsafe.Pointer(s))
	}

	semrelease(&s.count, false)
}

// ReleaseHandoff is an alternative to Release that is faster in some cases.
func (s *Phore) ReleaseHandoff() {
	if race.Enabled {
		race.Release(unsafe.Pointer(s))
	}

	semrelease(&s.count, true)
}

//go:linkname semacquire runtime.semacquire1
func semacquire(count *uint32, lifo bool, flags int)

//go:linkname semrelease runtime.semrelease1
func semrelease(count *uint32, handoff bool)
