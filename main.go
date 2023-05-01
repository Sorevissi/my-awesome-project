package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var romanNumberals = map[string]int{
	"I": 1,
	"V": 5,
	"X": 10,
	"L": 50,
	"C": 100,
	"D": 500,
	"M": 1000}

const (
	err1 = "Вывод ошибки, так как строка не является математической операцией"
	err2 = "Вывод ошибки, так как формат математической операции не удовлетворяет заданию - два операнда и один оператор (+, -, /, *)."
	err3 = "Вывод ошибки, так как используются дновременно разные системы счисления."
	err4 = "Вывод ошибки, так как в римской системе нет отрицательных чисел."
	err5 = "Вывод ошибки, так как в римской системе нет числа 0."
	err6 = "Калькулятор умеет работать только с арабскими целыми числами или римскими цифрами от 1 до 10 включительно"
)

func main() {
	var numberOperation string
	in := bufio.NewScanner(os.Stdin)
	in.Scan()
	numberOperation = in.Text()

	operation := strings.Split(numberOperation, " ")

	if len(operation) != 3 {
		fmt.Print(err1)
	} else if operation[1] == "+" || operation[1] == "-" || operation[1] == "*" || operation[1] == "/" {
		enteringAndValidatingData(operation)
	} else {
		fmt.Print(err2)
	}
}

func enteringAndValidatingData(operation []string) {
	if chekArabikNumber(operation[0]) && chekArabikNumber(operation[2]) {
		var x, y int
		x, _ = strconv.Atoi(operation[0])
		y, _ = strconv.Atoi(operation[2])
		fmt.Print(calculate(x, y, operation[1]))
	} else if chekRomanNumber(operation[0]) && chekRomanNumber(operation[2]) {
		var x, y int
		x = RomanToArabic(operation[0])
		y = RomanToArabic(operation[2])
		if zna := calculate(x, y, operation[1]); zna < 0 {
			fmt.Print(err4)
		} else if zna == 0 {
			fmt.Print(err5)
		} else {
			fmt.Print(ArabicToRoman(zna))
		}
	} else {
		if (chekRomanNumber(operation[0]) && chekArabikNumber(operation[2])) || (chekArabikNumber(operation[0]) && chekRomanNumber(operation[2])) {
			fmt.Print(err3)
		} else {
			fmt.Print(err1)
		}
	}
}

func RomanToArabic(roman string) int {
	result := 0
	for i := 0; i < len(roman); i++ {
		if i+1 < len(roman) && romanNumberals[string(roman[i])] < romanNumberals[string(roman[i+1])] {
			result -= romanNumberals[string(roman[i])]
		} else {
			result += romanNumberals[string(roman[i])]
		}
	}
	return result
}

func ArabicToRoman(arabyc int) string {
	result := ""
	var arabycOfRim = map[int]string{
		1000: "M", 900: "CM", 500: "D",
		400: "CD", 100: "C", 90: "XC",
		50: "L", 40: "XL", 10: "X",
		9: "IX", 5: "V", 4: "IV",
		1: "I"}

	arabycNumber := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}

	for _, v := range arabycNumber {
		for i := 0; i < arabyc/v; i++ {
			result += arabycOfRim[v]
		}
		arabyc %= v
	}
	return result
}

func chekRomanNumber(number string) bool {
	numb := true
	for rim := range number {
		if _, ok := romanNumberals[string(number[rim])]; ok {
		} else {
			numb = false
			return numb
		}
	}
	return numb
}

func chekArabikNumber(number string) bool {
	if _, err := strconv.Atoi(number); err != nil {
		return false
	} else {
		return true
	}
}

func calculate(x, y int, operator string) int {
	sum := func(x, y int) int { return x + y }
	subtraction := func(x, y int) int { return x - y }
	productOfNumbers := func(x, y int) int { return x * y }
	quotientOfNumbers := func(x, y int) int { return x / y }
	if (1 <= x && x <= 10) && (1 <= y && y <= 10) {
		switch operator {
		case "+":
			return sum(x, y)
		case "-":
			return subtraction(x, y)
		case "*":
			return productOfNumbers(x, y)
		case "/":
			return quotientOfNumbers(x, y)
		default:
			os.Exit(0)
			return 0
		}
	} else {
		fmt.Print(err6)
		os.Exit(0)
		return 0
	}
}
