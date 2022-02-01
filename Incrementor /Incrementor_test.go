package Incrementor

import (
	"github.com/stretchr/testify/assert"
	"math"
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
	assert.Equal(t, 10, incrementor.count)
	assert.NotNil(t, incrementor.Error())
}

func TestIncrementorIntNegativeMaxValue(t *testing.T) {
	incrementor := NewIncrementor()
	incrementor.setMaximumValue(-10)

	assert.Equal(t, math.MaxInt64, incrementor.maxCount)
	assert.Equal(t, 0, incrementor.count)
	assert.NotNil(t, incrementor.Error())
}