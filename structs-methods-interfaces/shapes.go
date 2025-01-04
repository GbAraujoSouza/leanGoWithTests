package structsmethodsinterfaces

import "math"

type Shape interface {
  Area() float64
}

type Rectangle struct {
  Width float64
  Height float64
}

type Triangle struct {
  Base float64
  Height float64
}

type Circle struct {
  Radius float64
}


// Perimeter calculates the perimeter of a rectangle
// given the width and the height
func (r Rectangle) Perimeter() float64 {
  return 2 * (r.Height + r.Width)
}

// Area calculates the area of a rectangle given
// the width and the height
func (r Rectangle) Area() float64 {
  return r.Width * r.Height
}

func (t Triangle) Area() float64 {
  return (t.Base * t.Height) / 2
}

func (c Circle) Area() float64 {
  return math.Pi * math.Pow(c.Radius, 2)
}
