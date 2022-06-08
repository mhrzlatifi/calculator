package main

import (
	"errors"
	"fmt"
)

func calc(num1, num2 int, operator string) (res float32) {

	switch operator {
	case "/":
		res = float32(num1) / float32(num2)
	case "*":
		res = float32(num1) * float32(num2)
	case "-":
		res = float32(num1) - float32(num2)
	case "+":
		res = float32(num1) + float32(num2)
	case "%":
		res = float32(num1 % num2)
	}
	return
}

func parseInputs(num1, num2 int, operator string) (res float32, err error) {

	operators := map[string]string{
		"/": "devide",
		"*": "multiply",
		"-": "Subtract",
		"+": "add",
		"%": "remainder",
	}

	_, ok := operators[operator]
	if !ok {
		res, err = 0, errors.New("Wrong operator")
		return
	}

	res, err = calc(num1, num2, operator), nil
	return
}

func main() {

	var a, b int
	var c string
	fmt.Print("Enter the first number: ")
	fmt.Scanf("%d", &a)
	fmt.Print("Enter the second number: ")
	fmt.Scanf("%d", &b)
	fmt.Print("Enter the operator: ")
	fmt.Scanf("%s", &c)

	res, err := parseInputs(a, b, c)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(res)
	}
}
