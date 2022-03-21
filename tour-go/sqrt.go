package main

import (
	"math"
)


func sqrt(x float64) float64{
	z := 1.0
	lastz:=0.0

	for  round(z) != round(lastz)  {
		lastz = z
		z -= (z*z - x) / (2*z)
	}


	return z
}

func round(x float64) float64 {
	return math.Round(x*100000)/100000
}