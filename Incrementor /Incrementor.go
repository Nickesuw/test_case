package Incrementor

import (
	"log"
	"math"
)

type Incrementor interface {
	getNumber() int
	incrementNumber()
	setMaximumValue(maximumValue int)
}

type IncrementorInt struct {
	count    int
	maxCount int
	err      error
}

// Used to add default values
func NewIncrementorInt() *IncrementorInt {
	Incrementor := &IncrementorInt{count: 0, maxCount: math.MaxInt}
	return Incrementor
}

// As the methods of interface Incrementor have no errors to return, we have to create a custom method
func (inc *IncrementorInt) GetError() error {
	return inc.err
}

// Returns the current value of the variable count
func (inc *IncrementorInt) getNumber() int {
	return inc.count
}

// Adds the value of count by 1. Before doing this, it checks if the value has reached the maximum value. If yes - count=0
func (inc *IncrementorInt) incrementNumber() {

	if inc.count == math.MaxInt {
		log.Print("failed to increment: int overflow")
		return
	}
	if inc.count >= inc.maxCount {
		inc.count = 0
		return
	}
	inc.count++
}

// Changes the maximum value that can be for the count variable. Checks for negative numbers and 0
func (inc *IncrementorInt) setMaximumValue(maximumValue int) {
	if maximumValue <= 0 {
		log.Print("failed to set a maximum value: the value must be more than 0")
		return
	}
	inc.maxCount = maximumValue
}
