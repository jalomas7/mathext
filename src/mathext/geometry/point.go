/*
    Copyright (C) 2016  Jacob Massengill

	point.go defines a generic two-dimmensional point
	and some various functions on points

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

type Point struct {
	x, y float64
}

func (p Point) GetPoint() (float64, float64) {
	return p.x,p.y
}

func (p *Point) SetPoint(x,y float64) {
	p.x = x
	p.y = y
}

//returns the distance between p and p2
func (p Point) Distance(p2 Point) float64 {
	x1, y1 := p.getPoint()
	x2, y2 := p2.getPoint()
	return math.Sqrt( math.Pow(x2-x1,2) + math.Pow(y2-y1,2)  )
}
