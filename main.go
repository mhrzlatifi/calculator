package main

import (
	"errors"
	"fmt"
	"log"
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

func calc(num1, num2 int, op string) (float32, error) {

	if op == "/" && num2 == 0 {
		return 0, errors.New("division by zero")
	}

	c := Calc{num1, num2, op}

	operators := map[string]func() float32{
		"/": c.Dev,
		"*": c.Multiply,
		"-": c.Sub,
		"+": c.Add,
		"%": c.Rem,
	}

	res, ok := operators[c.op]
	if !ok {
		return 0, errors.New("operator is not supported")
	}

	return res(), nil
}

func main() {

	var num1, num2 int
	var op string

	fmt.Print("Enter the first number: ")
	_, err := fmt.Scanf("%d", &num1)
	if err != nil {
		log.Println(err)
		return
	}

	fmt.Print("Enter the second number: ")
	_, err = fmt.Scanf("%d", &num2)
	if err != nil {
		log.Println("second number should be of type int")
		return
	}

	fmt.Print("Enter the operator: ")
	_, err = fmt.Scanf("%s", &op)
	if err != nil {
		log.Println("operator should be of type string")
		return
	}

	res, err := calc(num1, num2, op)
	if err != nil {
		log.Println(err.Error())
		return
	}

	fmt.Println(res)
}
