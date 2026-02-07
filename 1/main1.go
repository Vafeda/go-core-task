package main

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"strconv"
)

var (
	Salt = "go-2024"
)

var (
	ErrUnsupportedType = errors.New("unsupported type")
	ErrEmptyData       = errors.New("cannot hash empty data")
	ErrHashWriteFailed = errors.New("ошибка записи в хеш-функцию")
)

func main() {
	Task()
}

func Task() {
	// 1. Создает несколько переменных различных типов данных:
	// ```
	// int (три числа в десятичной, восьмеричной и шеснадцатиричной системах)
	// float64
	// string
	// bool
	// complex64
	// ```
	var (
		numDecimal     int       = 42
		numOctal       int       = 052
		numHexadecimal int       = 0x2A
		pi             float64   = 3.14
		name           string    = "Golang"
		isActive       bool      = true
		complexNum     complex64 = 1 + 2i
	)

	allVariable := []any{
		numDecimal,
		numOctal,
		numHexadecimal,
		pi,
		name,
		isActive,
		complexNum,
	}

	// 2. Определяет тип каждой переменной и выводит его на экран.
	for _, variable := range allVariable {
		fmt.Print(typeToString(variable))
	}

	// 3. Преобразует все переменные в строковый тип и объединяет их в одну строку.
	variablesString, err := allVariablesToString(allVariable)
	if err != nil {
		return
	}

	// 4. Преобразовать эту строку в срез рун.
	variableStringRune := []rune(variablesString)

	// 5. Захэшировать этот срез рун SHA256, добавив в середину соль "go-2024" и вывести результат.
	hash, err := hash(variableStringRune)
	if err != nil {
		return
	}
	fmt.Printf("Hash: %x\n", hash)
}

func typeToString(variable any) string {
	return fmt.Sprintf("Тип переменной: %T; Значение переменной: %v\n", variable, variable)
}

func allVariablesToString(variables []any) (string, error) {
	var variablesString string
	for _, variable := range variables {
		switch v := variable.(type) {
		case int:
			variablesString += strconv.Itoa(v)
		case float64:
			variablesString += strconv.FormatFloat(v, 'f', -1, 64)
		case string:
			variablesString += v
		case bool:
			variablesString += strconv.FormatBool(v)
		case complex64:
			variablesString += strconv.FormatComplex(complex128(v), 'f', -1, 64)
		default:
			return "", ErrUnsupportedType
		}
	}

	return variablesString, nil
}

func hash(data []rune) ([]byte, error) {
	if len(data) == 0 {
		return []byte{}, ErrEmptyData
	}

	dataStr := string(data)
	saltedData := dataStr[:len(dataStr)/2] + Salt + dataStr[len(dataStr)/2+1:]
	dataBytes := []byte(saltedData)

	h := sha256.New()

	n, err := h.Write(dataBytes)
	if err != nil {
		return nil, err
	}
	if n != len(dataBytes) {
		return nil, ErrHashWriteFailed
	}

	return h.Sum(nil), nil
}
