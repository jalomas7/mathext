/*
    Copyright (C) 2016  Jacob Massengill

    This program is free software: you can redistribute it and/or modify
    it under the terms of the GNU General Public License as published by
    the Free Software Foundation, either version 3 of the License, or
    (at your option) any later version.

    This program is distributed in the hope that it will be useful,
    but WITHOUT ANY WARRANTY; without even the implied warranty of
    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
    GNU General Public License for more details.

    You should have received a copy of the GNU General Public License
    along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/

package conversion

/*IMPERIAL UNITS*/

func InchesToFeet(i float64) float64 {
	return i/12
}

func FeetToInches(f float64) float64 {
	return f*12
}

func FeetToMiles(f float64) float64 {
	return f/5280
}

func MilesToFeet(m float64) float64 {
	return m*5280
}

func InchesToMiles(i float64) float64 {
	return FeetToMiles(InchesToFeet(i))
}

func MilesToInches(m float64) float64 {
	return FeetToInches(MilesToFeet(m))
}

func FeetToYards(f float64) float64 {
	return f/3
}

func YardsToFeet(y float64) float64 {
	return y*3
}

func InchesToYards(i float64) float64 {
	return FeetToYards(InchesToFeet(i))
}

func YardsToInches(y float64) float64 {
	return FeetToInches(YardsToFeet(y))
}

func YardsToMiles(y float64) float64 {
	return FeetToMiles(YardsToFeet(y))
}

/*METRIC UNITS*/

func scaleDown(u float64, scale int) float64 {
	if scale <= 0 {
		return u
	}else {
		return scaleDown(u*10, scale-1)
	}
}


func scaleUp(u float64, scale int) float64 {
	if scale <= 0 {
		return u
	}else {
		return scaleUp(u/10, scale-1)
	}
}

func CentimeterToDecimeter(c float64) float64 {
	return scaleUp(c, 1)
}

func CentimeterToMeter(c float64) float64 {
	return scaleUp(c, 2)
}

func CentimeterToKilometer(c float64) float64 {
	return scaleUp(c, 4)
}

func DecimeterToCentimeter(d float64) float64 {
	return scaleDown(d, 1)
}

func DecimeterToMeter(d float64) float64 {
	return scaleUp(d, 1)
}

func DecimeterToKilometer(d float64) float64 {
	return scaleUp(d, 3)
}

func MeterToDecimeter(m float64) float64 {
	return scaleDown(m, 1)
}

func MeterToCentimeter(m float64) float64 {
	return scaleDown(m, 2)
}

func MeterToKilometer(m float64) float64 {
	return scaleUp(m, 2)
}

func KilometerToMeter(k float64) float64 {
	return scaleDown(k, 2)
}

func KilometerToDecimeter(k float64) float64 {
	return scaleDown(k, 3)
}

func KilometerToCentimeter(k float64) float64 {
	return scaleDown(k, 4)
}
