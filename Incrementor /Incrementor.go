package Incrementor_

import (
	"log"
	"math"
)

type Incrementor struct {
	count    int64
	maxCount int64
}

// Used to add default values
func NewIncrementor() *Incrementor {
	incrementor := &Incrementor{count: 0, maxCount: math.MaxInt}
	return incrementor
}

// Returns the current value of the variable count
func (inc *Incrementor) getNumber() int64 {
	return inc.count
}

// Adds the value of count by 1. Before doing this, it checks if the value has reached the maximum value. If yes - count=0
func (inc *Incrementor) incrementNumber() {
	if inc.count == math.MaxUint64 {
		inc.count = 0
		return
	}
	inc.count++
}

// Changes the maximum value that can be for the count variable. Checks for negative numbers and 0
func (inc *Incrementor) setMaximumValue(maximumValue int64) {

	if maximumValue <= 0 {
		log.Print("failed to set a maximum value: the value must be more than 0")
		return
	}
	maximumValue = inc.maxCount
}
