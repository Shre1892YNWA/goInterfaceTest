package cars

type Cars interface {
	Grey() string
	FuelType() string
	Age() int
}

type car struct {
	colour  string
	mfgYear int
	fuel    string
}

func Newcar(c string, f string, y int) Cars {
	return &car{
		colour:  c,
		fuel:    f,
		mfgYear: y,
	}
}

func (c *car) Grey() string {
	if c.colour == "Grey" {
		return "Colour of car is grey"
	}
	return "Colour Grey is not available for this model"
}

func (c *car) Age() int {
	return 2022 - c.mfgYear
}

func (c *car) FuelType() string {
	return c.fuel
}
