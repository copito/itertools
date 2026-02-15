package group

import (
	"fmt"
	"reflect"
	"slices"
	"testing"
)

func TestGroupBySimple(t *testing.T) {
	type testCase struct {
		name     string
		input    []string
		keyFunc  func(string) int
		expected map[int][]string
	}

	testCases := []testCase{
		{
			name:  "group by length",
			input: []string{"a", "b", "aa", "bb", "ccc"},
			keyFunc: func(s string) int {
				return len(s)
			},
			expected: map[int][]string{
				1: {"a", "b"},
				2: {"aa", "bb"},
				3: {"ccc"},
			},
		},
		{
			name:  "group by length (unordered)",
			input: []string{"a", "aa", "b", "bb", "ccc"},
			keyFunc: func(s string) int {
				return len(s)
			},
			expected: map[int][]string{
				1: {"b"},
				2: {"bb"},
				3: {"ccc"},
			},
		},
		{
			name:  "group by first letter",
			input: []string{"apple", "apricot", "banana", "berry", "cherry"},
			keyFunc: func(s string) int {
				return int(s[0]) // group by ASCII code of first letter
			},
			expected: map[int][]string{
				int('a'): {"apple", "apricot"},
				int('b'): {"banana", "berry"},
				int('c'): {"cherry"},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := make(map[int][]string)
			for key, group := range GroupBy(slices.Values(tc.input), tc.keyFunc) {
				var groupSlice []string
				for val := range group {
					groupSlice = append(groupSlice, val)
				}
				result[key] = groupSlice
			}

			if !reflect.DeepEqual(result, tc.expected) {
				t.Errorf("expected %v, got %v", tc.expected, result)
			}
		})
	}
}

func TestGroupByComplexValue(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	input := []Person{
		{"Alice", 30},
		{"Charlie", 30},
		{"Bob", 25},
		{"David", 25},
		{"Eve", 35},
	}

	expected := map[int][]Person{
		30: {
			{"Alice", 30},
			{"Charlie", 30},
		},
		25: {
			{"Bob", 25},
			{"David", 25},
		},
		35: {
			{"Eve", 35},
		},
	}

	result := make(map[int][]Person)
	for key, group := range GroupBy(slices.Values(input), func(p Person) int { return p.Age }) {
		var groupSlice []Person
		for val := range group {
			groupSlice = append(groupSlice, val)
		}
		result[key] = groupSlice
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}

func TestGroupByComplexKey(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	type ComplexKey struct {
		AgeGroup   string
		LastLetter string
	}

	input := []Person{
		{"Alice", 30},
		{"Charlie", 30},
		{"Bob", 25},
		{"David", 25},
		{"Eve", 35},
	}

	expected := map[ComplexKey][]Person{
		{"30s", "e"}: {
			{"Alice", 30},
			{"Charlie", 30},
		},
		{"20s", "b"}: {
			{"Bob", 25},
		},
		{"20s", "d"}: {
			{"David", 25},
		},
		{"30s", "e"}: {
			{"Eve", 35},
		},
	}

	result := make(map[ComplexKey][]Person)
	for key, group := range GroupBy(slices.Values(input), func(p Person) ComplexKey {
		return ComplexKey{
			AgeGroup:   fmt.Sprintf("%ds", p.Age/10*10),
			LastLetter: string(p.Name[len(p.Name)-1]),
		}
	}) {
		var groupSlice []Person
		for val := range group {
			groupSlice = append(groupSlice, val)
		}
		result[key] = groupSlice
	}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("expected %v, got %v", expected, result)
	}
}
