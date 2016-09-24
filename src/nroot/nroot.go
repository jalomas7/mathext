package nroot

import "math"


//returns the rooth root of n
func Nroot(root , n float64) float64{
	return math.Pow(10, math.Log10(n)/root)
}
