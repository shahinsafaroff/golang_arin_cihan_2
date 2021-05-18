package main

import (
	"fmt"
	"strconv"
)

func CarExecute(c Carface){
	fmt.Println("\n")
	fmt.Println("Car Information: \n" + c.Information())
	fmt.Println("\n")

	msg := " "

	isRun := c.Run()
	if isRun {
		msg ="Is running"
	}else {
		msg = "Stopped"
	}
	fmt.Println("Car " + msg + ".")

	isStop := c.Stop()
	if isStop {
		msg = "Stopped"
	}else {
		msg = "Can't stop! Brakes are malfunctioning"
	}
	fmt.Println("Car " + msg + ".")
}

func main() {
	// Ferrari
	// Ferrari Brand information will be embedded into Ferrari's struct constructor NewFerrari()
	ferr := new(Ferrari)
	ferr.Brand = "Ferrari"
	ferr.Model = "F50"
	ferr.Color = "Red"
	ferr.Speed = 340
	ferr.Price = 1.4
	ferr.Special = true
	//fmt.Println(ferr.Information())

	merc := new(Mercedes)
	merc.Brand = "Mercedes"
	merc.Model = "CLK"
	merc.Color = "Black"
	merc.Speed = 290
	merc.Price = 0.4
	//fmt.Println(merc.Information())

	CarExecute(ferr)
	CarExecute(merc)
}
//Interface

type Carface interface {
	Run() 			bool
	Stop() 			bool
	Information() 	string
}

//Base Structs

type Car struct {
	Brand	string
	Model 	string
	Color 	string
	Speed	int
	Price	float64
}

type SpecialProduction struct {
	Special bool
}

//Ferrari
type Ferrari struct {
	Car //composition or inheritance
	SpecialProduction
}

func (_ Ferrari) Run() bool {
	return true
}

func (_ Ferrari) Stop() bool {
	return true
}

func (x *Ferrari) Information() string {
	ret := "\t" + x.Brand + " " + x.Model + "\n" + "\t" + x.Color + "\n" + "\t" + "Speed: " + strconv.Itoa(x.Speed) + "\n" + "\t" + "Price: " + strconv.FormatFloat(x.Price, 'g', -1, 64) + " Million"
	add := "Yes"
	if x.Special {
		ret += "\n" + "\t" + "Special: " + add
	}
	return ret
}

//Lamborghini
type Lamborghini struct {
	Car //composition or inheritance
	SpecialProduction
}

func (_ Lamborghini) Run() bool {
	return true
}

func (_ Lamborghini) Stop() bool {
	return true
}

func (x *Lamborghini) Information() string {
	ret := "\t" + x.Brand + " " + x.Model + "\n" + "\t" + x.Color + "\n" + "\t" + "Speed: " + strconv.Itoa(x.Speed) + "\n" + "\t" + "Price: " + strconv.FormatFloat(x.Price, 'g', -1, 64) + " Million"
	add := "Yes"
	if x.Special {
		ret += "\n" + "\t" + "Special: " + add
	}
	return ret
}

//Mercedes SLK
type Mercedes struct {
	Car //composition or inheritance
	SpecialProduction
}

func (_ Mercedes) Run() bool {
	return true
}

func (_ Mercedes) Stop() bool {
	return true
}

func (x *Mercedes) Information() string {
	ret := "\t" + x.Brand + " " + x.Model + "\n" + "\t" + x.Color + "\n" + "\t" + "Speed: " + strconv.Itoa(x.Speed) + "\n" + "\t" + "Price: " + strconv.FormatFloat(x.Price, 'g', -1, 64) + " Million"
	return ret
}

