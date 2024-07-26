package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Функция для превращения римских в арабские.
func romanToArabic(r string) int {
	// Мапа для соответствия римских чисел арабским

	romanNumerals := map[string]int{
		"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5,
		"VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10, "XL": 40,
		"L": 50, "C": 100}

	// Поиск значения в мапе
	val, exists := romanNumerals[r]
	if !exists {
		panic("Неверный формат числа") // Если значение не найдено, вызываем панику
	}
	// Возвращаем найденное значение
	return val
}

// функция для преобразования арабских чисел в римские

func arabicToRoman(arabicNumber int) string {
	// Определяем таблицу преобразования арабских чисел в римские
	arabicToRoman := []struct {
		value  int
		symbol string
	}{
		{100, "C"},
		{50, "L"},
		{40, "XL"},
		{10, "X"},
		{9, "IX"},
		{5, "V"},
		{4, "IV"},
		{1, "I"},
	}

	// Переменная для хранения результата
	result := ""

	// Проходим по таблице преобразования
	for _, pair := range arabicToRoman {
		// Пока число больше или равно текущему значению, добавляем символ и уменьшаем число
		for arabicNumber >= pair.value {
			result += pair.symbol
			arabicNumber -= pair.value
		}
	}
	return result
}

func calculate(a, b int, operator string) int {
	switch operator {
	case "+":
		return a + b

	case "-":
		return a - b

	case "*":
		return a * b

	case "/":
		if b == 0 {
			panic("Делить на ноль нельзя")
		}
		return a / b
	default:
		panic(" Оператор не найден")

	}
}

func main() {
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Println("Введите выражение (например, 3 + 2 или IV * III):")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		a, b, operator, isRoman := parseInput(input)
		result := calculate(a, b, operator)

		if isRoman {
			if result < 1 {
				panic("Римские числа не могут быть меньше 1")
			}
			fmt.Println("Результат:", arabicToRoman(result))
		} else {
			fmt.Println("Результат:", result)
		}
	}
}

func parseInput(input string) (int, int, string, bool) {
	var a, b int
	var operator string
	isRoman := false

	switch {
	case strings.Contains(input, "+"):
		operator = "+"
	case strings.Contains(input, "-"):
		operator = "-"
	case strings.Contains(input, "*"):
		operator = "*"
	case strings.Contains(input, "/"):
		operator = "/"
	default:
		panic("Оператор не найден!")

	}
	parts := strings.Split(input, operator)
	if len(parts) != 2 {
		panic("Не верный формат ввода!")
	}
	aStr, bStr := strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1])
	a, isRoman = ConvertingStr(aStr)
	b, _ = ConvertingStr(bStr)

	if a < 1 || a > 10 || b < 1 || b > 10 {
		panic("Числа должны быть от 1 до 10 включительно")
	}
	return a, b, operator, isRoman

}

//strings.Split возвращает срез (slice) строк, разделенных по оператору	                               //В нашем примере, strings.Split("3 + 5", "+") вернет ["3 ", " 5"].

func ConvertingStr(input string) (int, bool) {
	val, err := strconv.Atoi(input)
	if err == nil {
		return val, false
	}
	romanValue := romanToArabic(input)
	return romanValue, true
}
