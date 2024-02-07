package upordown

import (
	"unsafe"
)

// Stack growth direction
type StackGrowthDirection uint8

// Stack growth direction
const (
	// Stack grows upward
	Up = iota
	// Stack grows downward
	Down
)

// upordown determines the direction of stack growth
func UpOrDown() StackGrowthDirection {
	var x uint8
	return func(x *uint8) StackGrowthDirection {
		var y uint8
		xPtr := uintptr(unsafe.Pointer(x))
		yPtr := uintptr(unsafe.Pointer(&y))
		if xPtr < yPtr {
			return Up
		} else {
			return Down
		}
	}(&x)
}
