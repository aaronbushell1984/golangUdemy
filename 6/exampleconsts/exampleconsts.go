// Package exampleconsts demonstrates the use of consts in go by setting strings for door lock state.
//
// If variables are assigned to the const values:
//	const doorLocked, doorUnlocked string = "locked", "unlocked"
//
// A single change in the const declaration will update all uses.
//	const doorLocked, doorUnlocked string = "Locked", "Unlocked"
//
// Attempting to assign a new string to the status will warn developer with compilation error.
//	doorUnlocked = "Unlocked"
package exampleconsts

import "fmt"

const doorLocked, doorUnlocked string = "locked", "unlocked"

// UnlockDoor sets door VAR to the unchangeable CONST "unlocked" string
func UnlockDoor(door string) string {
	door = doorUnlocked
	return fmt.Sprint(door)
}

// LockDoor sets door VAR to the unchangeable CONST "locked" string
func LockDoor(door string) string {
	door = doorLocked
	return fmt.Sprint(door)
}
