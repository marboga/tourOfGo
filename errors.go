package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprint("cannot Sqrt negative number: ", float64(e))
}

func Sqrt(x ErrNegativeSqrt) (ErrNegativeSqrt, error) {
	if x >= 0 {
		return 0, nil
	} else {
		return 0, ErrNegativeSqrt(x)
	}
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
