package main

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"testing"
)

func TestTypeToString(t *testing.T) {
	tests := []struct {
		name     string
		input    any
		expected string
	}{
		{
			name:     "int",
			input:    42,
			expected: "Тип переменной: int; Значение переменной: 42\n",
		},
		{
			name:     "float64",
			input:    3.14,
			expected: "Тип переменной: float64; Значение переменной: 3.14\n",
		},
		{
			name:     "string",
			input:    "Golang",
			expected: "Тип переменной: string; Значение переменной: Golang\n",
		},
		{
			name:     "bool",
			input:    true,
			expected: "Тип переменной: bool; Значение переменной: true\n",
		},
		{
			name:     "complex64",
			input:    complex64(1 + 2i),
			expected: "Тип переменной: complex64; Значение переменной: (1+2i)\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := typeToString(tt.input)
			if result != tt.expected {
				t.Errorf("typeToString() = %v, want %v", result, tt.expected)
			}
		})
	}
}

func TestAllVariablesToString(t *testing.T) {
	tests := []struct {
		name        string
		input       []any
		expected    string
		expectError bool
	}{
		{
			name: "all supported types",
			input: []any{
				42,
				3.14,
				"Golang",
				true,
				complex64(1 + 2i),
			},
			expected:    "423.14Golangtrue(1+2i)",
			expectError: false,
		},
		{
			name: "only integers",
			input: []any{
				10,
				20,
				30,
			},
			expected:    "102030",
			expectError: false,
		},
		{
			name: "unsupported type",
			input: []any{
				42,
				struct{}{},
			},
			expected:    "",
			expectError: true,
		},
		{
			name:        "empty slice",
			input:       []any{},
			expected:    "",
			expectError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := allVariablesToString(tt.input)

			if tt.expectError {
				if err == nil {
					t.Error("Expected error, got nil")
				}
				if !errors.Is(err, ErrUnsupportedType) {
					t.Errorf("Expected ErrUnsupportedType, got %v", err)
				}
			} else {
				if err != nil {
					t.Errorf("Unexpected error: %v", err)
				}
				if result != tt.expected {
					t.Errorf("allVariablesToString() = %v, want %v", result, tt.expected)
				}
			}
		})
	}
}

func TestHash(t *testing.T) {
	originalSalt := Salt
	defer func() { Salt = originalSalt }()

	tests := []struct {
		name        string
		input       []rune
		salt        string
		expectError bool
		errorType   error
		expected    string
	}{
		{
			name:        "simple string",
			input:       []rune("test"),
			salt:        "test-salt",
			expectError: false,
			expected:    "",
		},
		{
			name:        "empty string",
			input:       []rune(""),
			salt:        "go-2024",
			expectError: true,
			errorType:   ErrEmptyData,
		},
		{
			name:        "unicode string",
			input:       []rune("тест"),
			salt:        "go-2024",
			expectError: false,
			expected:    "",
		},
		{
			name:        "odd length string",
			input:       []rune("12345"),
			salt:        "middle",
			expectError: false,
			expected:    "",
		},
		{
			name:        "nil slice",
			input:       nil,
			salt:        "go-2024",
			expectError: true,
			errorType:   ErrEmptyData,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Salt = tt.salt

			result, err := hash(tt.input)

			if tt.expectError {
				if err == nil {
					t.Error("Expected error, got nil")
					return
				}

				if tt.errorType != nil {
					if !errors.Is(err, tt.errorType) {
						t.Errorf("Expected error type %v, got %v", tt.errorType, err)
					}
				}

				return
			}

			if err != nil {
				t.Errorf("hash() returned unexpected error: %v", err)
				return
			}

			if len(result) != sha256.Size {
				t.Errorf("hash() returned %d bytes, expected %d", len(result), sha256.Size)
			}

			if tt.expected != "" {
				actual := fmt.Sprintf("%x", result)
				if actual != tt.expected {
					t.Errorf("hash() = %s, want %s", actual, tt.expected)
				}
			}
		})
	}
}
