package lexer

import (
	"reflect"
	"testing"
)

func TestIsSkippable(t *testing.T) {
	tests := []struct {
		char     rune
		expected bool
	}{
		{'a', false},
		{' ', true},
		{'#', false},
		{'$', false},
	}

	for _, test := range tests {
		result := IsSkippable(test.char)
		if result != test.expected {
			t.Errorf("Expected IsSkippable('%c') to be %v, but got %v", test.char, test.expected, result)
		}
	}
}

func TestNextChar(t *testing.T) {
	tests := []struct {
		src          []rune
		expectedChar string
		expectedSrc  []rune
	}{
		{[]rune{'a', 'b', 'c'}, "a", []rune{'b', 'c'}},
		{[]rune{'1', '2', '3'}, "1", []rune{'2', '3'}},
		{[]rune{}, "", []rune{}},
	}

	for _, test := range tests {
		resultChar := NextChar(&test.src)
		resultSrc := test.src

		if resultChar != test.expectedChar {
			t.Errorf("Expected NextChar(%v) to return '%s', but got '%s'", test.src, test.expectedChar, resultChar)
		}

		if !reflect.DeepEqual(resultSrc, test.expectedSrc) {
			t.Errorf("Expected source to be %v, but got %v", test.expectedSrc, resultSrc)
		}
	}
}

func TestExtractNum(t *testing.T) {
	tests := []struct {
		src         []rune
		expectedNum string
		expectedSrc []rune
	}{
		{[]rune{'1', '2', '3', 'a', '4', '5'}, "123", []rune{'a', '4', '5'}},
		{[]rune{'-', '5', '.', '6'}, "", []rune{'-', '5', '.', '6'}},
		{[]rune{}, "", []rune{}},
	}

	for _, test := range tests {
		resultNum, resultSrc := ExtractNum(test.src)

		if resultNum != test.expectedNum {
			t.Errorf("Expected ExtractNum(%v) to return '%s', but got '%s'", test.src, test.expectedNum, resultNum)
		}

		if !reflect.DeepEqual(resultSrc, test.expectedSrc) {
			t.Errorf("Expected source to be %v, but got %v", test.expectedSrc, resultSrc)
		}
	}
}

func TestExtractIdentifier(t *testing.T) {
	tests := []struct {
		src         []rune
		expectedStr string
		expectedSrc []rune
	}{
		{[]rune{'a', 'b', 'c', '1', '2', '3'}, "abc123", []rune{}},
		{[]rune{'1', 'a', 'b', 'c', '1', '2', '3'}, "", []rune{'1', 'a', 'b', 'c', '1', '2', '3'}},
		{[]rune{'_', 'x', 'y', 'z'}, "_xyz", []rune{}},
		{[]rune{}, "", []rune{}},
	}

	for _, test := range tests {
		resultStr, resultSrc := ExtractIdentifier(test.src)

		if resultStr != test.expectedStr {
			t.Errorf("Expected ExtractIdentifier(%v) to return '%s', but got '%s'", test.src, test.expectedStr, resultStr)
		}

		if !reflect.DeepEqual(resultSrc, test.expectedSrc) {
			t.Errorf("Expected source to be %v, but got %v", test.expectedSrc, resultSrc)
		}
	}
}

func TestIsIsAlpha(t *testing.T) {
	tests := []struct {
		char     rune
		expected bool
	}{
		{'a', true},
		{'_', true},
		{'1', false},
		{'$', false},
	}

	for _, test := range tests {
		result := IsAlpha(test.char)
		if result != test.expected {
			t.Errorf("Expected IsIdentifier('%c') to be %v, but got %v", test.char, test.expected, result)
		}
	}
}

func TestIsKeyword(t *testing.T) {
	tests := []struct {
		src        string
		expectedOK bool
	}{
		{"if", true},
		{"else", true},
		{"foo", false},
	}

	for _, test := range tests {
		_, resultOK := IsKeyword(test.src)

		if resultOK != test.expectedOK {
			t.Errorf("Expected IsKeyword('%s') to return OK as %v, but got %v", test.src, test.expectedOK, resultOK)
		}
	}
}

func TestIsInt(t *testing.T) {
	tests := []struct {
		char     rune
		expected bool
	}{
		{'5', true},
		{'-', false},
		{'.', false},
		{'a', false},
	}

	for _, test := range tests {
		result := IsInt(test.char)
		if result != test.expected {
			t.Errorf("Expected IsInt('%c') to be %v, but got %v", test.char, test.expected, result)
		}
	}
}
