package time

import (
	"syscall"
	"testing"
	"time"
	"unsafe"
)

func systime(t *syscall.Time_t) (tt syscall.Time_t, err error) {
	r0, _, e1 := syscall.RawSyscall(syscall.SYS_TIME, uintptr(unsafe.Pointer(t)), 0, 0)
	if e1 != 0 {
		err = e1
	}
	tt = syscall.Time_t(r0)
	return
}

func BenchmarkSyscallTime(b *testing.B) {
	for ii := 0; ii < b.N; ii++ {
		systime(nil)
	}
}

func BenchmarkVDSOTime(b *testing.B) {
	for ii := 0; ii < b.N; ii++ {
		syscall.Time(nil)
	}
}

func gettimeofday(tv *syscall.Timeval) (err error) {
	_, _, e1 := syscall.RawSyscall(syscall.SYS_GETTIMEOFDAY, uintptr(unsafe.Pointer(tv)), 0, 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func BenchmarkSyscallGettimeofday(b *testing.B) {
	var tv syscall.Timeval
	for ii := 0; ii < b.N; ii++ {
		gettimeofday(&tv)
	}
}

func BenchmarkVDSOGettimeofday(b *testing.B) {
	var t syscall.Timeval
	for ii := 0; ii < b.N; ii++ {
		syscall.Gettimeofday(&t)
	}
}

const (
	// Identifier for system-wide realtime clock.
	CLOCK_REALTIME = 0
	// Monotonic system-wide clock.
	CLOCK_MONOTONIC = 1
	// High-resolution timer from the CPU.
	CLOCK_PROCESS_CPUTIME_ID = 2
	// Thread-specific CPU-time clock.
	CLOCK_THREAD_CPUTIME_ID = 3
	// Monotonic system-wide clock, not adjusted for frequency scaling.
	CLOCK_MONOTONIC_RAW = 4
	// Identifier for system-wide realtime clock, updated only on ticks.
	CLOCK_REALTIME_COARSE = 5
	// Monotonic system-wide clock, updated only on ticks.
	CLOCK_MONOTONIC_COARSE = 6
	// Monotonic system-wide clock that includes time spent in suspension.
	CLOCK_BOOTTIME = 7
)

func clock_gettime(clockid int, ts *syscall.Timespec) (err error) {
	_, _, e1 := syscall.RawSyscall(syscall.SYS_CLOCK_GETTIME, uintptr(clockid), uintptr(unsafe.Pointer(ts)), 0)
	if e1 != 0 {
		err = e1
	}
	return
}

func BenchmarkSyscallClockGettimeRealtime(b *testing.B) {
	var ts syscall.Timespec
	for ii := 0; ii < b.N; ii++ {
		clock_gettime(CLOCK_REALTIME, &ts)
	}
}

func BenchmarkSyscallClockGettimeMonotonic(b *testing.B) {
	var ts syscall.Timespec
	for ii := 0; ii < b.N; ii++ {
		clock_gettime(CLOCK_MONOTONIC, &ts)
	}
}

func BenchmarkTimeNow(b *testing.B) {
	for ii := 0; ii < b.N; ii++ {
		time.Now()
	}
}
