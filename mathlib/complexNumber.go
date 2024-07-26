package main

import "math"

type Complx struct{
	pair[2] float64
}

func (num1 *Complx) conjugate() [2]float64{
	conjugate := [2]float64{num1.pair[0], -1 * num1.pair[1]}
	return conjugate
}

func (num1 *Complx) conjugateProduct() float64{
	return num1.pair[0] * num1.pair[0] + num1.pair[1] * num1.pair[1]
}

func (num1 *Complx) abs() float64{
	return math.Sqrt(num1.pair[0] * num1.pair[0] + num1.pair[1] * num1.pair[1])
}

func product(num1 *Complx, num2 *Complx) *Complx{
	prod := Complx{pair: [2]float64	{num1.pair[0] * num2.pair[0] - num1.pair[1] * num2.pair[1],
								num1.pair[0] * num2.pair[1] + num1.pair[1] * num2.pair[0]}}
	return &prod
}

func add(num1 *Complx, num2 *Complx) *Complx{
	res := Complx{ pair: [2] float64 {num1.pair[0] + num2.pair[0], num2.pair[1] + num2.pair[1]}}
	return &res
}

func sub(num1 *Complx, num2 *Complx) *Complx{
	res := Complx{ pair: [2] float64 {num1.pair[0] - num2.pair[0], num2.pair[1] - num2.pair[1]}}
	return &res
}


