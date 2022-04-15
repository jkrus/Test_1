package main

import (
	"log"
	"math/big"
)

type (
	Point struct {
		x, y *big.Float
	}
)

func NewPoint(x, y float64) *Point {
	return &Point{x: big.NewFloat(x), y: big.NewFloat(y)}
}

func (p *Point) GetX() *big.Float {
	return p.x
}

func (p *Point) GetY() *big.Float {
	return p.y
}

func main() {

	d := dist(NewPoint(1, 2), NewPoint(0, 0))

	log.Println(d)
}

func dist(p1, p2 *Point) *big.Float {
	tmp := &big.Float{}
	xd := &big.Float{}
	xd = xd.Add(p1.GetX(), tmp.Neg(p2.GetX()))
	xd = xd.Mul(xd, xd)

	yd := &big.Float{}
	yd = yd.Add(p1.GetY(), tmp.Neg(p2.GetY()))
	yd = yd.Mul(yd, yd)

	sum := &big.Float{}
	sum = sum.Add(xd, yd)

	return tmp.Sqrt(sum)
}
