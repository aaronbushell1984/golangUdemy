// Package iotas demonstrates the use of iotas in go
package iotas

import "fmt"

// ReturnFourIotaConsts returns a string demonstrating incrementing four iotas
func ReturnFourIotaConsts() string {
	const (
		a = iota
		b = iota
		c = iota
		d = iota
	)
	return fmt.Sprintf("%v, %v, %v, %v", a, b, c, d)
}

// ReturnBits takes kb, mb or gb and returns amount of bits as a string
//
// Bit shifting in increments of 10 using iota is used to obtain the values:
//		const (
//			_  = iota
//			kb = 1 << (iota * 10)
//			mb = 1 << (iota * 10)
//			gb = 1 << (iota * 10)
//		)
func ReturnBits(code string) string {
	const (
		_  = iota
		kb = 1 << (iota * 10)
		mb = 1 << (iota * 10)
		gb = 1 << (iota * 10)
	)
	switch {
	case code == "kb":
		return fmt.Sprint(kb)
	case code == "mb":
		return fmt.Sprint(mb)
	case code == "gb":
		return fmt.Sprint(gb)
	default:
		return "Not supported. Must be kb, mb or gb."
	}
}
