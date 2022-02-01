package main

import (
	"fmt"
	"math"
)

type Incrementor struct {
	count     int
	max_count int
}

func init_Incrementor() *Incrementor {
	struc := &Incrementor{count: 0, max_count: math.MaxInt}
	return struc
}

func (inc *Incrementor) getNumber() int {
	return inc.count
}

func (inc *Incrementor) incrementNumber() {
	if inc.count == inc.max_count {
		inc.count = 0
		return
	}
	inc.count++
}

func (inc *Incrementor) setMaximumValue(maximumValue int) {

	if maximumValue == 0 {
		fmt.Println("value cannot be zero")
		return
	}
}

func main() {

	// я не знаю как тут объявить объект структуры, потому что там вроде надо count вводить а он должен быть стандартный

}
