package main

import "fmt"

type Incrementor struct {
	count int
}

func (inc *Incrementor) getNumber() int {
	return inc.count
}

func (inc *Incrementor) incrementNumber() {
	inc.count++
}

func (inc *Incrementor) setMaximumValue(maximumValue int) {

	if maximumValue == 0 {
		fmt.Println("value cannot be zero")
		return
	}
}

func main() {

	incrementor := Incrementor{
		count: 10,
	}
	incrementor.setMaximumValue(0)

	incrementor.getNumber()

	incrementor.incrementNumber()

}

