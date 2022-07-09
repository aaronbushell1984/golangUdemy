package exampleconsts

import "testing"

func TestUnlockDoor(t *testing.T) {
	var door string = doorUnlocked
	want := "unlocked"
	if got := UnlockDoor(door); got != want {
		t.Errorf("UnlockDoor = %q, want %q", got, want)
	}
}

func TestLockDoor(t *testing.T) {
	var door string = doorLocked
	want := "locked"
	if got := LockDoor(door); got != want {
		t.Errorf("LockDoor = %q, want %q", got, want)
	}
}
