package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Конвертация из римских в арабские
func romanToAr(num string) int {
	romanmap := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}
	return romanmap[num]
}

// Конвертация из арабских в римсские
func arToRoman(num int) string {
	vals := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}
	romann := []string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	var res string
	for i := 0; i < len(vals); i++ {
		for num >= vals[i] {
			res += romann[i]
			num -= vals[i]
		}
	}
	return res
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Введите выражение, согласно формату: <число 1> <оператор> <число 2>")
	fmt.Println("------------------------------------------------------------------")
	fmt.Println("----------------------------ВАЖНО---------------------------------")
	fmt.Println("Числа не могут быть меньше 1 и больше 10!")
	fmt.Println("Доступные операторы: + - * /")

	input, err := reader.ReadString('\n')
	if err != nil {
		panic(fmt.Sprintf("Ввод некорректен!"))
	}

	parts := strings.Fields(input)
	if len(parts) != 3 {
		panic(fmt.Sprintf("Это не математическая операция!"))
	}

	var mode string
	var num1, num2 int

	// Проверка римских чисел на корректность ввода
	if romanToAr(parts[0]) > 0 {
		if romanToAr(parts[2]) > 0 {
			num1, num2 = romanToAr(parts[0]), romanToAr(parts[2])
			mode = "rom"
		} else {
			panic(fmt.Sprintf("Второе число - арабское, а первое - римское!"))
		}
	} else {
		if romanToAr(parts[2]) > 0 {
			panic(fmt.Sprintf("Первое число - арабское, а первое - римское!"))
		}
	}

	// Проверка арабских чисел на корректность ввода обоих чисел
	if mode != "rom" {
		num1, err = strconv.Atoi(parts[0])
		if err != nil {
			panic(fmt.Sprintf("Первый операнд - не число!"))
		}

		num2, err = strconv.Atoi(parts[2])
		if err != nil {
			panic(fmt.Sprintf("Второй операнд - не число!"))
		} else {
			mode = "ar"
		}
	}

	// Проверка арабских чисел на соответствие условиям
	if mode == "ar" {
		if num1 > 10 {
			panic(fmt.Sprintf("Первое число больше 10!"))
		}
		if num2 > 10 {
			panic(fmt.Sprintf("Второе число больше 10!"))
		}
		if num1 <= 0 {
			panic(fmt.Sprintf("Первое число меньше 1!"))
		}
		if num2 <= 0 {
			panic(fmt.Sprintf("Второе число меньше 1!"))
		}
	}

	var (
		result  int
		resultR string
	)

	// Проверка корректности ввода операторов + вычисление и вывод
	switch parts[1] {
	case "+":
		if mode == "rom" {
			resultR = arToRoman(num1 + num2)
		} else {
			result = num1 + num2
		}
	case "-":
		if mode == "rom" {
			if num1-num2 > 0 {
				resultR = arToRoman(num1 - num2)
			} else {
				panic(fmt.Sprintf("В результате получилось отрицательное римское число!"))
			}
		} else {
			result = num1 - num2
		}
	case "*":
		if mode == "rom" {
			resultR = arToRoman(num1 * num2)
		} else {
			result = num1 * num2
		}
	case "/":
		if num2 == 0 {
			panic(fmt.Sprintf("На ноль делить нельзя!"))
		}
		if mode == "rom" {
			resultR = arToRoman(num1 / num2)
		} else {
			result = num1 / num2
		}
	default:
		panic(fmt.Sprintf("Неверный оператор!"))
	}

	if mode == "rom" {
		fmt.Println("Результат: ", resultR)
	} else {
		fmt.Println("Результат: ", result)
	}
}
