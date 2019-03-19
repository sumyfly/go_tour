package main

// Calculator class
type Calculator struct {
	Value float64
}

func (c *Calculator) add(b float64) float64 {
	return c.Value + b
}

func (c *Calculator) del(b float64) float64 {
	c.Value = 10
	return c.Value - b
}

func (c Calculator) div(b float64) float64 {
	c.Value = 1000
	return c.Value / b
}

func mul(c *Calculator, b float64) float64 {
	return c.Value * b
}
