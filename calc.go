package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Введите математическое выражение,\n")
	var userInput string
	_, err := fmt.Scan(&userInput)
	if err != nil {
		panic("Ввод данных не корректен")
		return
	}
	userInput = strings.TrimSpace(userInput)
	var userInputLen, i int
	userInputLen = len(userInput)
	var matOperator, romeSimbol string
	var userInputOperand []string
	var operand1, operand2, itog int
	var isRomeNum1, isRomeNum2 bool
	inputChr := strings.Split(userInput, "")

	for i < userInputLen {
		switch inputChr[i] {
		case "+":
			matOperator = inputChr[i]
			userInputOperand = strings.Split(userInput, inputChr[i])
			operand1, isRomeNum1, romeSimbol = lineDigitization(userInputOperand[0])
			operand2, isRomeNum2, romeSimbol = lineDigitization(userInputOperand[1])
			operandConditions(operand1, operand2, isRomeNum1, isRomeNum2, userInputOperand)
			itog = operand1 + operand2

			romeSimbol = intToRome(itog)
		case "-":
			matOperator = inputChr[i]
			userInputOperand = strings.Split(userInput, inputChr[i])
			operand1, isRomeNum1, romeSimbol = lineDigitization(userInputOperand[0])
			operand2, isRomeNum2, romeSimbol = lineDigitization(userInputOperand[1])
			operandConditions(operand1, operand2, isRomeNum1, isRomeNum2, userInputOperand)
			itog = operand1 - operand2
			if isRomeNum1 == true && itog <= 0 {
				panic("Результат не соответсвует римской системе исчеслений,\n" +
					"полученное значение либо равно 0, либо меньше 0 \n" +
					"Результат на экран выведен не будет")
				return
			}
			romeSimbol = intToRome(itog)
		case "*":
			matOperator = inputChr[i]
			userInputOperand = strings.Split(userInput, inputChr[i])
			operand1, isRomeNum1, romeSimbol = lineDigitization(userInputOperand[0])
			operand2, isRomeNum2, romeSimbol = lineDigitization(userInputOperand[1])
			operandConditions(operand1, operand2, isRomeNum1, isRomeNum2, userInputOperand)
			itog = operand1 * operand2
			romeSimbol = intToRome(itog)
		case "/":
			matOperator = inputChr[i]
			userInputOperand = strings.Split(userInput, inputChr[i])
			operand1, isRomeNum1, romeSimbol = lineDigitization(userInputOperand[0])
			operand2, isRomeNum2, romeSimbol = lineDigitization(userInputOperand[1])
			operandConditions(operand1, operand2, isRomeNum1, isRomeNum2, userInputOperand)
			itog = operand1 / operand2
			romeSimbol = intToRome(itog)

		}
		i++
	}
	if matOperator == "" {
		panic("Введен неизвестный (или вовсе отсутствует) математический оператор, \n" +
			"возможно вы использовали недопустимые символы или числа. \n" +
			"для вычисления необходимы два операнда и математический оператор \n" +
			"Разрешены целые числа от 1 до 10 (I - X в римских цифрах)" +
			"а так же знаки математических операций + - * /")
		return
	}
	if isRomeNum1 == true {
		fmt.Printf("\n\n%v=%v", userInput, romeSimbol)
	} else {
		fmt.Printf("\n\n%v=%v", userInput, itog)
	}
}

func lineDigitization(operandStr string) (operandInt int, validRomeSmb bool, romeSinbol string) {

	const maxRomanNumVal int = 101
	var validOperand bool
	operandInt, err := strconv.Atoi(operandStr)
	validOperand = true
	if err != nil {
		validOperand = false
		operandStr = strings.ToUpper(operandStr)
		validVariants := [maxRomanNumVal]string{"0", "I", "II", "III", "IV", "V", "VI",
			"VII", "VIII", "IX", "X", "XI", "XII", "XIII", "XIV", "XV", "XVI", "XVII", "XVIII",
			"XIX", "XX", "XXI", "XXII", "XXIII", "XXIV", "XXV", "XXVI", "XXVII", "XXVIII",
			"XXIX", "XXX", "XXXI", "XXXII", "XXXIII", "XXXIV", "XXXV", "XXXVI", "XXXVII",
			"XXXVIII", "XXXIX", "XL", "XLI", "XLII", "XLIII", "XLIV", "XLV", "XLVI", "XLVII",
			"XLVIII", "XLIX", "L", "LI", "LII", "LIII", "LIV", "LV", "LVI", "LVII", "LVIII",
			"LIX", "LX", "LXI", "LXII", "LXIII", "LXIV", "LXV", "LXVI", "LXVII", "LXVIII",
			"LXIX", "LXX", "LXXI", "LXXII", "LXXIII", "LXXIV", "LXXV", "LXXVI", "LXXVII",
			"LXXVIII", "LXXIX", "LXXX", "LXXXI", "LXXXII", "LXXXIII", "LXXXIV", "LXXXV", "LXXXVI",
			"LXXXVII", "LXXXVIII", "LXXXIX", "XC", "XCI", "XCII", "XCIII", "XCIV", "XCV", "XCVI",
			"XCVII", "XCVIII", "XCIX", "C"}

		for index, simbol := range validVariants {
			if operandStr == simbol {
				validOperand = true
				validRomeSmb = true
				operandInt = index
				if index == 0 {
					operandInt = index
					validRomeSmb = false
					validOperand = false
				}
			}
		}
	}

	if validOperand == false {
		panic("Недопустимые символы или числа. " +
			"Разрешены целые числа от 1 до 10 (I - X в римских цифрах)" +
			"а так же знаки математических операций + - * /")
		return
	}
	return
}

func intToRome(inpunint int) (output string) {
	const maxRomanNumVal int = 100
	if inpunint < 0 || inpunint > maxRomanNumVal {
		return
	}
	validVariants := [maxRomanNumVal]string{"0", "I", "II", "III", "IV", "V", "VI",
		"VII", "VIII", "IX", "X", "XI", "XII", "XIII", "XIV", "XV", "XVI", "XVII", "XVIII",
		"XIX", "XX", "XXI", "XXII", "XXIII", "XXIV", "XXV", "XXVI", "XXVII", "XXVIII",
		"XXIX", "XXX", "XXXI", "XXXII", "XXXIII", "XXXIV", "XXXV", "XXXVI", "XXXVII",
		"XXXVIII", "XXXIX", "XL", "XLI", "XLII", "XLIII", "XLIV", "XLV", "XLVI", "XLVII",
		"XLVIII", "XLIX", "L", "LI", "LII", "LIII", "LIV", "LV", "LVI", "LVII", "LVIII",
		"LIX", "LX", "LXI", "LXII", "LXIII", "LXIV", "LXV", "LXVI", "LXVII", "LXVIII",
		"LXIX", "LXX", "LXXI", "LXXII", "LXXIII", "LXXIV", "LXXV", "LXXVI", "LXXVII",
		"LXXVIII", "LXXIX", "LXXX", "LXXXI", "LXXXII", "LXXXIII", "LXXXIV", "LXXXV", "LXXXVI",
		"LXXXVII", "LXXXVIII", "LXXXIX", "XC", "XCI", "XCII", "XCIII", "XCIV", "XCV", "XCVI",
		"XCVII", "XCVIII", "XCIX"}
	output = validVariants[inpunint]
	return
}

func operandConditions(operand1 int, operand2 int, isRomeNum1 bool, isRomeNum2 bool, userInputOperand []string) {
	if isRomeNum1 != isRomeNum2 {
		panic("Допускается ввод чисел в одном типе представления чисел")
		return
	}
	if operand1 == 0 || operand2 == 0 {
		panic("0 Недопустимое число для ввода")
		return
	}
	if operand1 > 10 || operand2 > 10 {
		panic("Допускается ввод числа не более 10 (X)")
		return
	}
	if len(userInputOperand) > 2 {
		panic("Слишком много операндов\n" +
			"Допускается только 2 операнда (одна арифмитическая операция)")
		return
	}
}
