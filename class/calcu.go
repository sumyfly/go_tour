package main

import "errors"

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

func (c *Calculator) div(b float64) (r float64, err error) {
	defer func() {
		if p := recover(); p != nil {
			str, ok := p.(string)
			if ok {
				err = errors.New(str)
			} else {
				err = errors.New("panic")
			}
			r = -9999999
		}
	}()

	if b == 0 {
		panic("Can not divide zero!")
	}

	println("????")
	// r = c.Value / b
	return r, err
}

func mul(c *Calculator, b float64) float64 {
	return c.Value * b
}
