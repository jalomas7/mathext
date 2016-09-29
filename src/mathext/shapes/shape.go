/*
    Copyright (C) 2016  Jacob Massengill

    shapes.go is a definition of a generic shape interface and a
    generic implementation for it.

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

package shapes

type GeoShape interface {
  Area() float64
  Perimeter() float64
}

type Shape struct{
  Sides []float64
  Angles []float64
}

func (s Shape) Area() float64 {
  return 1.0
}

func (s Shape) Perimeter() float64 {
  sum := 0
  for _, i := range Sides{
    sum+=i
  }
  return sum
}
