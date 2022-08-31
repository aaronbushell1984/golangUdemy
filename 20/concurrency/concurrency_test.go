package concurrency

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
)

func ExampleGetArchitecture() {
	fmt.Println(GetArchitecture())
	// Output:
	// amd64
}

func ExampleGetOs() {
	fmt.Println(GetOs())
	// Output:
	// windows
}

func ExampleGetNumberOfCpus() {
	fmt.Println(GetNumberOfCpus())
	// Output:
	// 8
}

func ExampleGetNumberOfGoRoutines() {
	fmt.Printf("%T", GetNumberOfGoRoutines())
	// Output:
	// int
}

func ExampleCreateOneSimpleGoRoutine() {
	CreateOneSimpleGoRoutine()
	fmt.Printf("%T", GetNumberOfGoRoutines())
	// Output:
	// int
}

func TestCreaOneSimpleGoRoutine(t *testing.T) {
	initial := runtime.NumGoroutine()
	CreateOneSimpleGoRoutine()
	updated := runtime.NumGoroutine()
	got := updated - initial
	want := 1
	if got != want {
		t.Errorf("CreateOneSimpleGoRoutine() got: %q, want: %q", got, want)
	}
}

func ExampleWaitGroupCount_GetCount() {
	wgc := WaitGroupCount{
		sync.WaitGroup{},
		0,
	}
	fmt.Println(wgc.GetCount())
	// Output:
	// 0
}

func ExampleWaitGroupCount_Add() {
	wgc := WaitGroupCount{
		sync.WaitGroup{},
		0,
	}
	wgc.Add(1)
	wgc.Add(1)
	fmt.Println(wgc.GetCount())
	// Output:
	// 2
}

func ExampleWaitGroupCount_Done() {
	wgc := WaitGroupCount{
		sync.WaitGroup{},
		0,
	}
	wgc.Add(1)
	wgc.Add(1)
	wgc.Add(1)
	fmt.Println(wgc.GetCount())
	wgc.Done()
	fmt.Println(wgc.GetCount())
	// Output:
	// 3
	// 2
}

func ExampleGetStatus() {
	fmt.Println(GetStatus(&status))
	// Output:
	// running
}

func ExampleSetStatus() {
	SetStatus(&status, "completed")
	fmt.Println(GetStatus(&status))
	// Output:
	// completed
}

func ExampleWaitForStatusUpdate() {
	WaitForStatusUpdate()
	fmt.Println(GetStatus(&status))
	// Output:
	// completed
}

func ExampleDoNotWaitForStatusUpdate() {
	running := "running"
	SetStatus(&status, running)
	DoNotWaitForStatusUpdate()
	fmt.Println(GetStatus(&status))
	// Output:
	// running
}
