// Package concurrency demonstrates the use of code which could run parallelly in go
//
// This package contains a custom struct and implementation of sync.WaitGroup to provide a count for testing
//	type WaitGroupCount struct {
//		sync.WaitGroup
//		count int64
// 	}
// 	func (wg *WaitGroupCount) Add(delta int) {
//		atomic.AddInt64(&wg.count, int64(delta))
//		wg.WaitGroup.Add(delta)
// 	}
// 	func (wg *WaitGroupCount) Done() {
//		atomic.AddInt64(&wg.count, -1)
//		wg.WaitGroup.Done()
// 	}
// 	func (wg *WaitGroupCount) GetCount() int {
//		return int(atomic.LoadInt64(&wg.count))
// 	}
package concurrency

import (
	"runtime"
	"sync"
	"sync/atomic"
)

type WaitGroupCount struct {
    sync.WaitGroup
    count int64
}

var (wgc = WaitGroupCount{
	sync.WaitGroup{},
	0,
})

var status string = "running"

// WaitGroupCount_Add adds one delta to sync.Waitgroup and customer counter
func (wg *WaitGroupCount) Add(delta int) {
    atomic.AddInt64(&wg.count, int64(delta))
    wg.WaitGroup.Add(delta)
}

// WaitGroupCount_Done completes one delta on sync.Waitgroup and decreases customer counter by one
func (wg *WaitGroupCount) Done() {
    atomic.AddInt64(&wg.count, -1)
    wg.WaitGroup.Done()
}

// WaitGroupCount_GetCount returns the custom count on sync.Waitgroup
func (wg *WaitGroupCount) GetCount() int {
    return int(atomic.LoadInt64(&wg.count))
}

// GetArchitecture returns the running program's architecture
func GetArchitecture() string {
	return runtime.GOARCH
}

// GetOs returns running program's operating system
func GetOs() string {
	return runtime.GOOS
}

// GetNumberOfCpus returns the number of logical CPUs usable by the current process
func GetNumberOfCpus() int {
	return runtime.NumCPU()
}

// GetNumberOfGoRoutines returns the number of goroutines that currently exist
func GetNumberOfGoRoutines() int {
	return runtime.NumGoroutine()
}

// CreateOneSimpleGoRoutine creates one go routine which does nothing but demonstrate increment of the number of go routines
func CreateOneSimpleGoRoutine() {
	go GetNumberOfGoRoutines()
}

// GetStatus returns the current value from the pointer to status
func GetStatus(status *string) string {
	return *status
}

// UpdateStatus updates the status with the provided new status
func UpdateStatus(status *string, newStatus string) {
	*status = newStatus
}

// UpdateStatusWithDone updates the status with the provided new status
//
// Completes a worker with:
//	wgc.Done()
func UpdateStatusWithDone(status *string, newStatus string) {
	*status = newStatus
	wgc.Done()
}

// DoNotWaitForStatusUpdate runs a go routine to UpdateStatus:
//	go UpdateStatus(&status, "completed")
// This does not update the status as the function completes before routine does 
func DoNotWaitForStatusUpdate() {
	go UpdateStatus(&status, "completed")
}

// WaitForStatusUpdate runs a go routine to UpdateStatusWithDone() - Compare with DoNotWaitForStatusUpdate()
//
// First waitgroup is added:
//	wgc.Add(1)
// UpdateStatusWithDone is called as go routine:
//	go UpdateStatusWithDone(&status, "completed")
// UpdateStatusWithDone contains a completion of waitgroup:
//	*status = newStatus
//	wgc.Done()
// Wait is called to hold completion of function until go routine completes:
//	wgc.Wait()
func WaitForStatusUpdate() {
	wgc.Add(1)
	go UpdateStatusWithDone(&status, "completed")
	wgc.Wait()
}