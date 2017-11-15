// +build !race

package race

import "unsafe"

const Enabled = false

func Acquire(ptr unsafe.Pointer) {}
func Release(ptr unsafe.Pointer) {}
