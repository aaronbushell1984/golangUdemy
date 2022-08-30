// Package dog allows us to more fully understand dogs.
package dog

import "fmt"

// Years converts human years to dog years.
func Years(n int) (int, error) {
	if n < 0 {
		return 0, fmt.Errorf("human years cannot be negative")
	}
	return n * 7, nil
}

// YearsTwo converts human years to dog years.
func YearsTwo(n int) (int, error) {
	if n < 0 {
		return 0, fmt.Errorf("human years cannot be negative")
	}
	count := 0
	for i := 0; i < n; i++ {
		count += 7
	}
	return count, nil
}
