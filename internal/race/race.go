// +build race

package race

import (
	"runtime"
	"unsafe"
)

const Enabled = true

func Acquire(ptr unsafe.Pointer) {
	runtime.RaceAcquire(ptr)
}

func Release(ptr unsafe.Pointer) {
	runtime.RaceRelease(ptr)
}
