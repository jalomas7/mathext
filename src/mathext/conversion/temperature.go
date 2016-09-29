/*
    Copyright (C) 2016  Jacob Massengill

    Temperature.go provides temperature conversion between Farenheit, 
    Celsius, and Kelvin temperatures
	
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

//returns the farenheit temperature of c
func CelsiusToFahrenheit(c float64) float64 {
	return c * (9/5) + 32
}

//returns the celsius temperature of f
func FahrenheitToCelsius(f float64) float64 {
	return f * (5/9) - 32
}

//returns the kelvin temperature of c
func CelsiusToKelvin(c float64) float64 {
	return c + 273.15
}

//returns the kelvin temperature of f
func FarenheitToKelvin(f float64) float64 {
	return CelsiusToKelvin(FahrenheitToCelsius(f))
}

//returns the celsius temperature of k
func KelvinToCelsius(k float64) float64 {
	return k - 273.15
}

//returns the Farenheit temperature of k
func KelvinToFarenheit(k float64) float64 {
	return CelsiusToFahrenheit(KelvinToCelsius(k))
}
