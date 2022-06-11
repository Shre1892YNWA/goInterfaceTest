package main

import (
	"fmt"
	"pracint/cars"
)

func main() {
	ci := cars.Newcar("Yellow", "Disel", 2014)
	ciColor := ci.Grey()
	ciFuel := ci.FuelType()
	ciAge := ci.Age()

	fmt.Printf("%s , and fuel type is %s, age is %d years\n", ciColor, ciFuel, ciAge)
	//fmt.Println(ciColor, "and fuel type is", ciFuel, "age is", ciAge, "years")
}
