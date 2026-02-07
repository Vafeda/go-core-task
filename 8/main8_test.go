package main

import (
	"fmt"
	"testing"
)

func TestHashTable_Add(t *testing.T) {
	tests := []struct {
		name          string
		arraySize     uint
		key           string
		value         int
		expectedValue int
	}{
		{
			name:          "Добавление нового элемента",
			arraySize:     10,
			key:           "testKey",
			value:         42,
			expectedValue: 42,
		},
		{
			name:          "Обновление существующего элемента",
			arraySize:     10,
			key:           "testKey",
			value:         100,
			expectedValue: 100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ht := New(tt.arraySize)

			// Добавляем элемент
			ht.Add(tt.key, tt.value)

			// Проверяем, что элемент добавлен корректно
			value, exists := ht.Get(tt.key)
			if !exists {
				t.Errorf("Add() элемент не добавлен")
			}

			if value != tt.expectedValue {
				t.Errorf("Add() = %v, ожидалось %v", value, tt.expectedValue)
			}
		})
	}
}

func TestHashTable_Remove(t *testing.T) {
	tests := []struct {
		name        string
		arraySize   uint
		key         string
		value       int
		removeKey   string
		shouldExist bool
	}{
		{
			name:        "Удаление существующего элемента",
			arraySize:   10,
			key:         "key1",
			value:       10,
			removeKey:   "key1",
			shouldExist: false,
		},
		{
			name:        "Удаление несуществующего элемента",
			arraySize:   10,
			key:         "key1",
			value:       10,
			removeKey:   "key2",
			shouldExist: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ht := New(tt.arraySize)

			// Добавляем элемент
			ht.Add(tt.key, tt.value)

			// Удаляем элемент
			ht.Remove(tt.removeKey)

			// Проверяем результат
			exists := ht.Exists(tt.removeKey)
			if exists != tt.shouldExist {
				t.Errorf("Remove() = %v, ожидалось %v для ключа %s",
					exists, tt.shouldExist, tt.removeKey)
			}
		})
	}
}

func TestHashTable_Copy(t *testing.T) {
	tests := []struct {
		name      string
		arraySize uint
		elements  map[string]int
	}{
		{
			name:      "Пустая таблица",
			arraySize: 10,
			elements:  map[string]int{},
		},
		{
			name:      "Один элемент",
			arraySize: 10,
			elements:  map[string]int{"key1": 1},
		},
		{
			name:      "Несколько элементов",
			arraySize: 10,
			elements:  map[string]int{"key1": 1, "key2": 2, "key3": 3},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ht := New(tt.arraySize)

			// Добавляем элементы
			for key, value := range tt.elements {
				ht.Add(key, value)
			}

			// Копируем
			copyMap := ht.Copy()

			// Проверяем размер
			if len(copyMap) != len(tt.elements) {
				t.Errorf("Copy() размер = %d, ожидалось %d",
					len(copyMap), len(tt.elements))
			}

			// Проверяем все элементы
			for key, expectedValue := range tt.elements {
				actualValue, exists := copyMap[key]
				if !exists {
					t.Errorf("Copy() ключ %s отсутствует", key)
				}

				if actualValue != expectedValue {
					t.Errorf("Copy() для ключа %s = %d, ожидалось %d",
						key, actualValue, expectedValue)
				}
			}

			// Проверяем, что это действительно копия
			if len(tt.elements) > 0 {
				firstKey := ""
				for k := range tt.elements {
					firstKey = k
					break
				}

				// Изменяем оригинальную таблицу
				ht.Add(firstKey, 999)

				// Проверяем, что копия не изменилась
				if copyMap[firstKey] == 999 {
					t.Errorf("Copy() возвращает ссылку, а не копию")
				}
			}
		})
	}
}

func TestHashTable_Exists(t *testing.T) {
	tests := []struct {
		name        string
		arraySize   uint
		addKey      string
		addValue    int
		checkKey    string
		shouldExist bool
	}{
		{
			name:        "Ключ существует",
			arraySize:   10,
			addKey:      "existingKey",
			addValue:    42,
			checkKey:    "existingKey",
			shouldExist: true,
		},
		{
			name:        "Ключ не существует",
			arraySize:   10,
			addKey:      "existingKey",
			addValue:    42,
			checkKey:    "nonExistingKey",
			shouldExist: false,
		},
		{
			name:        "Пустая таблица",
			arraySize:   10,
			addKey:      "",
			addValue:    0,
			checkKey:    "anyKey",
			shouldExist: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ht := New(tt.arraySize)

			if tt.addKey != "" {
				ht.Add(tt.addKey, tt.addValue)
			}

			exists := ht.Exists(tt.checkKey)

			if exists != tt.shouldExist {
				t.Errorf("Exists() для ключа %s = %v, ожидалось %v",
					tt.checkKey, exists, tt.shouldExist)
			}
		})
	}
}

func TestHashTable_Get(t *testing.T) {
	tests := []struct {
		name          string
		arraySize     uint
		addKey        string
		addValue      int
		getKey        string
		expectedValue int
		shouldExist   bool
	}{
		{
			name:          "Получение существующего элемента",
			arraySize:     10,
			addKey:        "key1",
			addValue:      100,
			getKey:        "key1",
			expectedValue: 100,
			shouldExist:   true,
		},
		{
			name:          "Получение несуществующего элемента",
			arraySize:     10,
			addKey:        "key1",
			addValue:      100,
			getKey:        "key2",
			expectedValue: 0,
			shouldExist:   false,
		},
		{
			name:          "Обновление и получение элемента",
			arraySize:     10,
			addKey:        "key1",
			addValue:      200,
			getKey:        "key1",
			expectedValue: 200,
			shouldExist:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ht := New(tt.arraySize)

			// Добавляем элемент
			ht.Add(tt.addKey, tt.addValue)

			// Пробуем получить элемент
			value, exists := ht.Get(tt.getKey)

			// Проверяем флаг существования
			if exists != tt.shouldExist {
				t.Errorf("Get() exists = %v, ожидалось %v",
					exists, tt.shouldExist)
			}

			// Проверяем значение
			if exists && value != tt.expectedValue {
				t.Errorf("Get() value = %d, ожидалось %d",
					value, tt.expectedValue)
			}

			// Проверяем, что для несуществующего элемента возвращается 0
			if !exists && value != 0 {
				t.Errorf("Get() для несуществующего элемента вернул %d вместо 0",
					value)
			}
		})
	}
}

func TestHashTable_EdgeCases(t *testing.T) {
	t.Run("Коллизии хэшей", func(t *testing.T) {
		ht := New(1) // Маленький размер для гарантированных коллизий

		// Добавляем несколько элементов, которые попадут в один бакет
		ht.Add("a", 1) // 'a' = 97
		ht.Add("b", 2) // 'b' = 98

		// Проверяем, что оба элемента доступны
		if !ht.Exists("a") {
			t.Error("Exists() не находит элемент 'a' при коллизии")
		}

		if !ht.Exists("b") {
			t.Error("Exists() не находит элемент 'b' при коллизии")
		}

		valA, _ := ht.Get("a")
		if valA != 1 {
			t.Errorf("Get() для 'a' = %d, ожидалось 1", valA)
		}

		valB, _ := ht.Get("b")
		if valB != 2 {
			t.Errorf("Get() для 'b' = %d, ожидалось 2", valB)
		}
	})

	t.Run("Удаление из середины цепочки", func(t *testing.T) {
		// Создаем таблицу с маленьким размером для тестирования цепочек
		ht := New(1)

		// Добавляем больше элементов, чем BucketSize
		for i := 0; i < BucketSize*2; i++ {
			key := fmt.Sprintf("key%d", i)
			ht.Add(key, i)
		}

		// Удаляем элемент из середины
		ht.Remove("key5")

		// Проверяем, что элемент удален
		if ht.Exists("key5") {
			t.Error("Remove() не удалил элемент из середины цепочки")
		}

		// Проверяем, что другие элементы все еще доступны
		if !ht.Exists("key0") {
			t.Error("Remove() удалил не тот элемент")
		}

		if !ht.Exists("key10") {
			t.Error("Remove() повредил цепочку")
		}
	})

	t.Run("Полная копия с цепочками", func(t *testing.T) {
		ht := New(1)

		// Создаем длинную цепочку
		elements := make(map[string]int)
		for i := 0; i < BucketSize*3; i++ {
			key := fmt.Sprintf("k%d", i)
			ht.Add(key, i*10)
			elements[key] = i * 10
		}

		// Копируем
		copyMap := ht.Copy()

		// Проверяем все элементы
		for key, expectedValue := range elements {
			actualValue, exists := copyMap[key]
			if !exists {
				t.Errorf("Copy() не скопировал ключ %s", key)
			}
			if actualValue != expectedValue {
				t.Errorf("Copy() для ключа %s = %d, ожидалось %d",
					key, actualValue, expectedValue)
			}
		}
	})
}
