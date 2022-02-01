package Incrementor

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIncrementorIntChangeMaxValue(t *testing.T) {
	incrementor := NewIncrementor()
	incrementor.setMaximumValue(10)
	for i := 0; i < 10; i++ {
		incrementor.incrementNumber()
	}
	incrementor.setMaximumValue(5)
	incrementor.incrementNumber()
	assert.Equal(t, 5, incrementor.maxCount)
	assert.Equal(t, 0, incrementor.count)
}
