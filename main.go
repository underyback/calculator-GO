package main

import (
	"fmt"
)

type number struct {
	Number   int
	IfArabic bool
}

func ArabToRome(a int) string {
	RomeMap := map[int]string{
		1:  "I",
		2:  "II",
		3:  "III",
		4:  "IV",
		5:  "V",
		6:  "VI",
		7:  "VII",
		8:  "VIII",
		9:  "IX",
		10: "X",
		20: "XX",
		30: "XXX",
		40: "XL",
		50: "L",
		60: "LX",
		70: "LXX",
		80: "LXXX",
		90: "XC",
	}
	if a == 100 {
		return "C"
	} else {
		b := a / 10
		b *= 10
		c := a % 10
		return RomeMap[b] + RomeMap[c]
	}
}

func add(a number, b number) {
	c := a.Number + b.Number
	fmt.Print("Результат сложения: ")
	if a.IfArabic {
		fmt.Println(c)
	} else {
		fmt.Println(ArabToRome(c))
	}

}
func sub(a number, b number) {
	if a.IfArabic == false && a.Number <= b.Number {
		fmt.Println("Результат выходит за границы римской системы счисления")
	} else {
		c := a.Number - b.Number
		fmt.Print("Результат вычитания: ")
		if a.IfArabic {
			fmt.Println(c)
		} else {
			fmt.Println(ArabToRome(c))
		}
	}
}
func mult(a number, b number) {
	c := a.Number * b.Number
	fmt.Print("Результат умножения: ")
	if a.IfArabic {
		fmt.Println(c)
	} else {
		fmt.Println(ArabToRome(c))
	}
}
func div(a number, b number) {
	c := a.Number / b.Number
	fmt.Print("Результат деления: ")
	if a.IfArabic {
		fmt.Println(c)
	} else {
		fmt.Println(ArabToRome(c))
	}
}

func main() {
	var a, b, s string
	var ai, bi number
	IntMap := map[string]number{
		"1":    {1, true},
		"2":    {2, true},
		"3":    {3, true},
		"4":    {4, true},
		"5":    {5, true},
		"6":    {6, true},
		"7":    {7, true},
		"8":    {8, true},
		"9":    {9, true},
		"10":   {10, true},
		"I":    {1, false},
		"II":   {2, false},
		"III":  {3, false},
		"IV":   {4, false},
		"V":    {5, false},
		"VI":   {6, false},
		"VII":  {7, false},
		"VIII": {8, false},
		"IX":   {9, false},
		"X":    {10, false},
	}
	fmt.Print("Введите операцию ")
	fmt.Scan(&a, &s, &b)
	value, ok := IntMap[a]
	if ok {
		ai = value
	} else {
		fmt.Println("Введенно число вне допустимого диапазона")
		return
	}
	value, ok = IntMap[b]
	if ok {
		bi = value
	} else {
		fmt.Println("Введенно число вне допустимого диапазона")
		return
	}

	if ai.IfArabic != bi.IfArabic {
		fmt.Println("Числа из разных систем счисления")
	}

	switch s {
	case "+":
		add(ai, bi)
	case "-":
		sub(ai, bi)
	case "*":
		mult(ai, bi)
	case "/":
		div(ai, bi)
	default:
		fmt.Println("Введен неверный оператор")
		return
	}
	fmt.Print("Введите любой символ чтобы выйти из программы ")
	fmt.Scan(&a)
}
