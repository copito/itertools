package sequence

import (
	"slices"
	"testing"
)

func TestCount(t *testing.T) {
	t.Run("integer count", func(t *testing.T) {
		var result []int
		i := 0
		for val := range Count(10, 1) {
			if i >= 5 {
				break
			}
			result = append(result, val)
			i++
		}

		expected := []int{10, 11, 12, 13, 14}
		if !slices.Equal(result, expected) {
			t.Errorf("expected %v, got %v", expected, result)
		}
	})

	t.Run("float count", func(t *testing.T) {
		var result []float64
		i := 0
		for val := range Count(2.5, 0.5) {
			if i >= 4 {
				break
			}
			result = append(result, val)
			i++
		}

		expected := []float64{2.5, 3.0, 3.5, 4.0}
		if !slices.Equal(result, expected) {
			t.Errorf("expected %v, got %v", expected, result)
		}
	})

	t.Run("negative step", func(t *testing.T) {
		var result []int
		i := 0
		for val := range Count(10, -2) {
			if i >= 5 {
				break
			}
			result = append(result, val)
			i++
		}

		expected := []int{10, 8, 6, 4, 2}
		if !slices.Equal(result, expected) {
			t.Errorf("expected %v, got %v", expected, result)
		}
	})
}

func TestCycle(t *testing.T) {
	t.Run("cycle basic", func(t *testing.T) {
		data := []string{"A", "B", "C"}
		var result []string
		i := 0
		for val := range Cycle(slices.Values(data)) {
			if i >= 8 {
				break
			}
			result = append(result, val)
			i++
		}

		expected := []string{"A", "B", "C", "A", "B", "C", "A", "B"}
		if !slices.Equal(result, expected) {
			t.Errorf("expected %v, got %v", expected, result)
		}
	})

	t.Run("cycle single element", func(t *testing.T) {
		data := []int{42}
		var result []int
		i := 0
		for val := range Cycle(slices.Values(data)) {
			if i >= 5 {
				break
			}
			result = append(result, val)
			i++
		}

		expected := []int{42, 42, 42, 42, 42}
		if !slices.Equal(result, expected) {
			t.Errorf("expected %v, got %v", expected, result)
		}
	})

	t.Run("cycle empty", func(t *testing.T) {
		data := []int{}
		var result []int
		for val := range Cycle(slices.Values(data)) {
			result = append(result, val)
			// Should not execute
		}

		if len(result) != 0 {
			t.Errorf("expected empty result, got %v", result)
		}
	})
}

func TestRepeat(t *testing.T) {
	t.Run("repeat with count", func(t *testing.T) {
		var result []int
		for val := range Repeat(10, 3) {
			result = append(result, val)
		}

		expected := []int{10, 10, 10}
		if !slices.Equal(result, expected) {
			t.Errorf("expected %v, got %v", expected, result)
		}
	})

	t.Run("repeat zero times", func(t *testing.T) {
		var result []string
		for val := range Repeat("hello", 0) {
			result = append(result, val)
		}

		if len(result) != 0 {
			t.Errorf("expected empty result, got %v", result)
		}
	})

	t.Run("repeat infinite", func(t *testing.T) {
		var result []string
		i := 0
		for val := range Repeat("x") {
			if i >= 5 {
				break
			}
			result = append(result, val)
			i++
		}

		expected := []string{"x", "x", "x", "x", "x"}
		if !slices.Equal(result, expected) {
			t.Errorf("expected %v, got %v", expected, result)
		}
	})

	t.Run("repeat struct", func(t *testing.T) {
		type Point struct {
			X, Y int
		}
		point := Point{X: 5, Y: 10}

		var result []Point
		for val := range Repeat(point, 3) {
			result = append(result, val)
		}

		expected := []Point{{5, 10}, {5, 10}, {5, 10}}
		if !slices.Equal(result, expected) {
			t.Errorf("expected %v, got %v", expected, result)
		}
	})
}
