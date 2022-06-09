package main

import (
	"fmt"
)

type Calc struct {
	Num1, Num2 int
}

func (c *Calc) Add() float32 {
	return float32(c.Num1 + c.Num2)
}

func (c *Calc) Sub() float32 {
	return float32(c.Num1 - c.Num2)
}

func (c *Calc) Dev() float32 {
	return float32(c.Num1 / c.Num2)
}

func (c *Calc) Multiply() float32 {
	return float32(c.Num1 * c.Num2)
}

func (c *Calc) Rem() float32 {
	return float32(c.Num1 % c.Num2)
}

func main() {
	var op string

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("wrong operator!!! plz enter operator again")
		}
	}()

	c := Calc{}

	fmt.Print("Enter the first number: ")
	fmt.Scanf("%d", &c.Num1)
	fmt.Print("Enter the second number: ")
	fmt.Scanf("%d", &c.Num2)
	fmt.Print("Enter the operator: ")
	fmt.Scanf("%s", &op)

	operators := map[string]func() float32{
		"/": c.Dev,
		"*": c.Multiply,
		"-": c.Sub,
		"+": c.Add,
		"%": c.Rem,
	}

	res, ok := operators[op]
	if !ok {
		fmt.Println("wrong operator!!!")
		return
	}

	fmt.Println(res())

}
