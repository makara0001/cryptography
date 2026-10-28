package main

import (
	"fmt"
)
// func Add() int {
// 	x := 10
// 	x += 10
// 	return x
// }
// func Sub() int {
// 	x := 10
// 	x -= 5
// 	return x
// }
// func Time() int {
// 	x := 10
// 	x *= 5
// 	return x
// }
// func Divide() int {
// 	x := 10
// 	x /= 5
// 	return x
// }
// func per() int {
// 	x := 10
// 	x %= 3
// 	return x
// }


// func main() {
// 	Addition := Add()
// 	Subtract := Sub()
// 	Time := Time()
// 	Divide := Divide()
// 	percentage := per()


// 	fmt.Printf("the result: %d\n", Addition)
// 	fmt.Printf("the result: %d\n", Subtract)
// 	fmt.Printf("the result: %d\n", Time)
// 	fmt.Printf("the result: %d\n", Divide)
// 	fmt.Printf("the result: %d\n", percentage)


// }

func Operation() (int,int,int,int,int) {
	var num1 int 
	var num2 int

	fmt.Printf("Enter the first number :")
	fmt.Scanln(&num1)
	fmt.Printf("Enter the second number :")
	fmt.Scanln(&num2)

	add := num1 + num2
	sub := num1 - num2
	time := num1 * num2
	divide := num1 / num2
	per := num1 % num2
	return add, sub, time, divide,per
	
}

func main(){

	add, sub, time, divide,per := Operation()
	fmt.Printf("Addition: %d\n", add)
	fmt.Printf("Subtraction: %d\n", sub)
	fmt.Printf("Multiplication: %d\n", time)
	fmt.Printf("Division: %d\n", divide)
	fmt.Printf("Remainder: %d\n", per)
}