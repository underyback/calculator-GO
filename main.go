package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
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
func sub(a number, b number) error {
	if a.IfArabic == false && a.Number <= b.Number {
		err := errors.New("результат выходит за границы римской системы счисления")
		return err
	} else {
		c := a.Number - b.Number
		fmt.Print("Результат вычитания: ")
		if a.IfArabic {
			fmt.Println(c)
		} else {
			fmt.Println(ArabToRome(c))
		}
	}
	return nil
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
func div(a number, b number) error {

	if a.IfArabic == false && a.Number <= b.Number {
		err := errors.New("результат выходит за границы римской системы счисления")
		return err
	} else {
		c := a.Number / b.Number
		fmt.Print("Результат деления: ")

		if a.IfArabic {
			fmt.Println(c)
		} else {
			fmt.Println(ArabToRome(c))
		}
	}
	return nil
}

func main() {
	var a int
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
	reader := bufio.NewReader(os.Stdin)
	line, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	line = strings.ToUpper(line)
	data := strings.Fields(line)

	if len(data) > 3 {
		err := errors.New("превышено количество операций")
		fmt.Print(err)
		log.Fatal(err)
	}

	value, ok := IntMap[data[0]]
	if ok {
		ai = value
	} else {
		err := errors.New("введенно число вне допустимого диапазона")
		log.Fatal(err)
	}
	value, ok = IntMap[data[2]]
	if ok {
		bi = value
	} else {
		err := errors.New("введенно число вне допустимого диапазона")
		log.Fatal(err)
	}

	if ai.IfArabic != bi.IfArabic {
		err := errors.New("числа из разных систем счисления")
		log.Fatal(err)
	}

	switch data[1] {
	case "+":
		add(ai, bi)
	case "-":
		err = sub(ai, bi)
		if err != nil {
			log.Fatal(err)
		}
	case "*":
		mult(ai, bi)
	case "/":
		err = div(ai, bi)
		if err != nil {
			log.Fatal(err)
		}
	default:
		err := errors.New("введен неверный оператор")
		log.Fatal(err)
	}
	fmt.Print("Введите любой символ чтобы выйти из программы ")
	fmt.Scan(&a)
}
