package main

import (
	"fmt"
	. "github.com/Frank-Mayer/upordown"
)

func main() {
	switch UpOrDown() {
	case Up:
		fmt.Println("Stack grows upward")
	case Down:
		fmt.Println("Stack grows downward")
	}
}
