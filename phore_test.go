package sema

import "testing"

func BenchmarkAcquireFast(b *testing.B) {
	var sem Phore
	sem.SetCount(1<<32 - 1)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			sem.Acquire()
		}
	})
}

func BenchmarkAcquireSlow(b *testing.B) {
	var sem Phore

	go func(n int) {
		for i := 0; i < n; i++ {
			sem.Release()
		}
	}(b.N)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			sem.Acquire()
		}
	})
}

func BenchmarkBroadcast(b *testing.B) {
	var sem Phore
	sem.SetCount(1)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			sem.Acquire()
			sem.Release()
		}
	})
}

func BenchmarkReleaseFast(b *testing.B) {
	var sem Phore
	sem.SetCount(1)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			sem.Release()
		}
	})
}

func BenchmarkReleaseSlow(b *testing.B) {
	var sem Phore

	go func(n int) {
		for i := 0; i < n; i++ {
			sem.Acquire()
		}
	}(b.N)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			sem.Release()
		}
	})
}
