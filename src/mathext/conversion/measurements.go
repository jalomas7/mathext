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
