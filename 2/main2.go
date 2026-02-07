package main

import (
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"
)

const (
	defaultSliceLen = 10
)

var (
	ErrCopyMismatch = errors.New("number of copied elements doesn't match source length")
)

func main() {
	// 1. Создайте слайс целых чисел originalSlice, содержащий 10 произвольных значений, которые генерируются случайным
	// образом (при каждом запуске должны получаться новые значения)
	numbers := createNewSlice()
	fmt.Println(`Cозданный слайс со псевдослучайными числами:`, numbers)

	// 2. Напишите функцию sliceExample, которая принимает слайс и возвращает новый слайс,
	// содержащий только четные числа из исходного слайса.
	onlyEvenNumbers := sliceExample(numbers)
	fmt.Println(`Слайс содержащий только четные числа:`, onlyEvenNumbers)

	// 3. Напишите функцию addElements, которая принимает слайс и число.
	// Функция должна добавлять это число в конец слайса и возвращать новый слайс.
	newNumbers, err := addElements(numbers, 2)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(`Слайс созданный с добавленным элементов:`, newNumbers)

	// 4. Напишите функцию copySlice, которая принимает слайс и возвращает его копию.
	// Убедитесь, что изменения в оригинальном слайсе не влияют на его копию.
	copyNumbers, err := copySlice(numbers)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(`Скопированный слайс:`, newNumbers)
	fmt.Printf("Адресса изначального слайса. 0-ой элемент: %p. %d-ый элемент: %p\n",
		&numbers[0],
		len(numbers)-1,
		&numbers[len(numbers)-1])
	fmt.Printf("Адресса копии слайса. 0-ой элемент: %p. %d-ый элемент: %p\n",
		&copyNumbers[0],
		len(copyNumbers)-1,
		&copyNumbers[len(copyNumbers)-1])

	// 5. Напишите функцию removeElement, которая принимает слайс и индекс элемента, который нужно удалить.
	// Функция должна возвращать новый слайс без элемента по указанному индексу.
	newNumbers = removeElement(numbers, 2)
	fmt.Println("Новый слайс с удаленным вторым элементом из изначального слайса:", newNumbers)
}

func createNewSlice() []int {
	newNumbers := make([]int, defaultSliceLen)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < defaultSliceLen; i++ {
		newNumbers[i] = r.Int() % 100
	}

	return newNumbers
}

func sliceExample(numbers []int) []int {
	newNumbers := make([]int, 0, len(numbers))
	for _, number := range numbers {
		if number%2 == 0 {
			newNumbers = append(newNumbers, number)
		}
	}
	return newNumbers
}

func addElements(numbers []int, newNumber int) ([]int, error) {
	newNumbers := make([]int, len(numbers)+1)
	num := copy(newNumbers, numbers)
	if num != len(numbers) {
		return nil, ErrCopyMismatch
	}
	newNumbers = append(newNumbers, newNumber)
	return newNumbers, nil
}

func copySlice(numbers []int) ([]int, error) {
	newNumbers := make([]int, len(numbers))
	num := copy(newNumbers, numbers)
	if num != len(numbers) {
		return nil, ErrCopyMismatch
	}
	return newNumbers, nil
}

func removeElement(numbers []int, numberIndex int) []int {
	if len(numbers) == 0 {
		return []int{}
	}

	newNumbers := make([]int, 0, len(numbers)-1)
	newNumbers = append(newNumbers, numbers[:numberIndex]...)
	newNumbers = append(newNumbers, numbers[numberIndex+1:]...)
	return newNumbers
}
