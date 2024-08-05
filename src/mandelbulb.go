package main

import(
	"math"
    "mathlib"
)

type mandelbulb struct{
	quat mathlib.Quat
    c float64

}

func createBulb(comp, n1, x1, y1, z1 float64) *mandelbulb{
    q := mathlib.Quat{X: x1, Y: y1, Z: z1, R: n1}
    
    bulb := mandelbulb{	quat: q, c: comp}
	return &bulb
}

func (man *mandelbulb) phi() float64{
	temp := man.quat
	return math.Atan(temp.Y / temp.X) 
}

func (man *mandelbulb) radius() float64{
	temp := man.quat
	return math.Sqrt(math.Pow(temp.X, 2) + math.Pow(temp.Y, 2) + math.Pow(temp.Z, 2)) 
}

func (man *mandelbulb) theta() float64{
	temp := man.quat
	return math.Atan(math.Sqrt(math.Pow(temp.X, 2) + math.Pow(temp.Y, 2)) / temp.Z)
}

func (man *mandelbulb) calcVector() *mandelbulb{
    vec := man.quat
    rn := math.Pow(man.radius(), vec.R)

    vec.X = rn * math.Sin(vec.R * man.theta()) * math.Cos(vec.R * man.phi())
    vec.Y = rn * math.Sin(vec.R * man.theta()) * math.Sin(vec.R * man.phi()) 
    vec.Z = rn * math.Cos(vec.R * man.theta())

    return man
}


