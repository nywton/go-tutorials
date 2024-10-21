package main

import "fmt"

type gasEngine struct {
	mpg     uint8
	gallons uint8
	owner
	int
}

type eletricEngine struct {
	mpkwh uint8
	kwh   uint8
}

type owner struct {
	name string
}

type engine interface {
	milesLeft() uint8
}

func (e gasEngine) milesLeft() uint8 {
	return e.gallons * e.mpg
}

func (e eletricEngine) milesLeft() uint8 {
	return e.kwh * e.mpkwh
}

func canMakeIt(e engine, miles uint8) bool {
	if miles <= e.milesLeft() {
		fmt.Println("We can make it!")
		return true
	} else {
		fmt.Println("We can't make it!")
		return false
	}
}
func main() {
	var myEngine = gasEngine{25, 40, owner{"Alex"}, 12}

	fmt.Println(myEngine.mpg, myEngine.gallons, myEngine.name)
	fmt.Println("Total miles left: ", myEngine.milesLeft())

	var myEngine2 = eletricEngine{25, 20}

	fmt.Println(myEngine2.mpkwh, myEngine2.kwh)
	fmt.Println("Total miles left: ", myEngine2.milesLeft())
}
