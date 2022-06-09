package main

import (
	"fmt"
)

type Calc struct {
	Num1, Num2 int
	op         string
}

func (c Calc) Add() float32 {
	return float32(c.Num1 + c.Num2)
}

func (c Calc) Sub() float32 {
	return float32(c.Num1 - c.Num2)
}

func (c Calc) Dev() float32 {
	return float32(c.Num1 / c.Num2)
}

func (c Calc) Multiply() float32 {
	return float32(c.Num1 * c.Num2)
}

func (c Calc) Rem() float32 {
	return float32(c.Num1 % c.Num2)
}

func calc() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
			calc()
		}
	}()

	c := Calc{}

	fmt.Print("Enter the first number: ")
	_, err := fmt.Scanf("%d", &c.Num1)
	if err != nil {
		panic(err)
	}

	fmt.Print("Enter the second number: ")
	_, err = fmt.Scanf("%d", &c.Num2)
	if err != nil {
		panic(err)
	}

	fmt.Print("Enter the operator: ")
	_, err = fmt.Scanf("%s", &c.op)
	if err != nil {
		panic(err)
	}

	operators := map[string]func() float32{
		"/": c.Dev,
		"*": c.Multiply,
		"-": c.Sub,
		"+": c.Add,
		"%": c.Rem,
	}

	res := operators[c.op]()
	fmt.Println(res)
}

func main() {
	calc()
}
