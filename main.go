package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Введите выражение:")
	str := strings.Split(scan(), " ")
	switch {

	case len(str) < 3:
		{
			fmt.Println("Строка не является математической операцией")
			return
		}
	case len(str) > 3:
		{
			fmt.Println("Формат математической операции не удовлетворяет заданию — два операнда и один оператор")
			return
		}
	case romanToInt(str[0]) != 0 && romanToInt(str[2]) == 0 ||
		romanToInt(str[0]) == 0 && romanToInt(str[2]) != 0:
		{
			fmt.Println("Используются одновременно разные системы счисления")
			return
		}
	case romanToInt(str[0]) != 0 && romanToInt(str[2]) != 0:
		{
			if !testRoman(str[0]) || !testRoman(str[2]) {
				fmt.Println("Некорректный ввод аргументов")
				return
			}

			result, err := calc(romanToInt(str[0]), str[1], romanToInt(str[2]))
			if err == nil {
				if result <= 0 {
					fmt.Println("В римской системе нет отрицательных чисел")
				}
				fmt.Println(roman(result))
			} else {
				fmt.Println(err)
			}
		}
	case romanToInt(str[0]) == 0 && romanToInt(str[2]) == 0:
		{
			x, err1 := strconv.Atoi(str[0])
			y, err2 := strconv.Atoi(str[2])
			if err1 != nil || err2 != nil {
				fmt.Println("Некорректный ввод аргументов")
				return
			}
			result, err := calc(x, str[1], y)
			if err == nil {
				fmt.Println(result)
			} else {
				fmt.Println(err)
			}

		}
	}
}

func scan() string {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	if sc.Err() != nil {
		fmt.Println("Ошибка ввода. Err: ", sc.Err())
	}
	if sc.Text() == "" {
		fmt.Println("Ошибка ввода. Вы ничего не ввели")
	}
	return strings.ToUpper(sc.Text())
}

func roman(number int) string {
	conversions := []struct {
		value int
		digit string
	}{
		{1000, "M"},
		{900, "CM"},
		{500, "D"},
		{400, "CD"},
		{100, "C"},
		{90, "XC"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	roman := ""
	for _, conversion := range conversions {
		for number >= conversion.value {
			roman += conversion.digit
			number -= conversion.value
		}
	}
	return roman
}

func testRoman(s string) bool {
	arrRoman := []string{"I", "V", "X", "L", "C", "D", "M"}
	str := strings.Split(s, "")
	u := 0
	for i := 0; i < len(str); i++ {
		for j := 0; j < len(arrRoman); j++ {
			if strings.Contains(str[i], arrRoman[j]) {
				u++
			}
		}
	}
	if u != len(s) {
		return false
	}
	return true
}

func romanToInt(s string) int {
	rMap := map[string]int{"I": 1, "V": 5, "X": 10, "L": 50, "C": 100, "D": 500, "M": 1000}

	result := 0

	for k := range s {
		if k < len(s)-1 && rMap[s[k:k+1]] < rMap[s[k+1:k+2]] {
			result -= rMap[s[k:k+1]]
		} else {
			result += rMap[s[k:k+1]]
		}
	}
	return result
}

func calc(x int, operation string, y int) (int, error) {
	if x > 10 || x <= 0 || y > 10 || y <= 0 {
		return 0, errors.New("Калькулятор должен принимать на вход числа от 1(I) до 10(X) включительно, не более ")
	}
	switch operation {
	case "+":
		return x + y, nil
	case "-":
		return x - y, nil
	case "*":
		return x * y, nil
	case "/":
		return x / y, nil
	default:
		return 0, errors.New("Недопустимая операция")
	}
}
