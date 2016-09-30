/*
    Copyright (C) 2016  Jacob Massengill

	circle.go is the definition of a generic circle
	using the shape interface from shape.go

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
import "math"

type Circle struct {
	radius float64	
}

func (c Circle) Area() float64 {
	return math.Pi*math.Pow(c.radius,2)
}

func (c Circle) Perimeter() float64 {
	return 2*math.Pi*c.radius
}

func (c Circle) Circumference() float64 {
	return c.Perimeter()
}

func (c Circle) GetRadius() float64 {
	return c.radius
}
