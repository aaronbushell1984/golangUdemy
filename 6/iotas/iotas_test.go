package iotas

import (
	"fmt"
	"testing"
)

func TestReturnFourIotaConsts(t *testing.T) {
	want := "0, 1, 2, 3"
	if got := ReturnFourIotaConsts(); got != want {
		t.Errorf("ReturnFourIotaConsts() got: %q, wanted: %q", got, want)
	}
}

func ExampleReturnBits() {
	fmt.Println(ReturnBits("kb"))
	fmt.Println(ReturnBits("mb"))
	fmt.Println(ReturnBits("gb"))
	fmt.Println(ReturnBits("Nonsense"))
	// Output:
	// 1024
	// 1048576
	// 1073741824
	// Not supported. Must be kb, mb or gb.
}
