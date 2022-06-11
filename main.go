package main

import (
	"fmt"
	"pracint/cars"
)

func main() {
	ci := cars.Newcar("Grey", "Disel", 2014)
	ciColor := ci.Grey()
	ciFuel := ci.FuelType()
	ciAge := ci.Age()

	fmt.Println(ciColor, "and fuel type is", ciFuel, "age is", ciAge, "years")
}
