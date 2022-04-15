package main

import (
	"log"
	"math/big"
)

func main() {
	x := new(big.Int)
	x.SetString("10000000000000000000000000", 10)
	y := new(big.Int)
	y.SetString("1000000000000000000000000", 10)

	tmp := &big.Int{}
	tmp = tmp.Add(x, y)
	log.Println("Sum = ", tmp)
	tmp = tmp.Mul(x, y)
	log.Println("Sub = ", tmp)
	tmp = tmp.Sub(x, y)
	log.Println("Mul = ", tmp)
	tmp = tmp.Div(x, y)
	log.Println("Div = ", tmp)

}
