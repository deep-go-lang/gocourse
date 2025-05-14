package advanced

import (
	"math/rand"
	"testing"
)

func GenerateRandomSlice(n int) []int {
	slice := make([]int, n)
	for i := 0; i < n; i++ {
		slice[i] = rand.Intn(1000)
	}
	return slice
}

func SumSlice(slice []int) int {
	sum := 0
	for _, v := range slice {
		sum += v
	}
	return sum
}

func TestGenerateRandomSlice(t *testing.T) {
	slice := GenerateRandomSlice(100)
	if len(slice) != 100 {
		t.Errorf("Expected slice length 100, got %d", len(slice))
	}
}

func BenchmarkGenerateRandomSlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GenerateRandomSlice(1000)
	}
}

// func TestSumSlice(t *testing.T) {
// 	slice := GenerateRandomSlice(100)
// 	sum := SumSlice(slice)
// 	if sum != 499500 {
// 		t.Errorf("Expected sum 499500, got %d", sum)
// 	}
// }

func BenchmarkSumSlice(b *testing.B) {
	slice := GenerateRandomSlice(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		SumSlice(slice)
	}
}

// func Add(a, b int) int {
// 	return a + b
// }

// func BenchmarkAddSmallInputs(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		Add(1, 2)
// 	}
// }

// func BenchmarkAddLargeInputs(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		Add(1000000, 2000000)
// 	}
// }

// func BenchmarkAddMediumInputs(b *testing.B) {
// 	for i := 0; i < b.N; i++ {
// 		Add(1000, 2000)
// 	}
// }

// func TestAddSubTest(t *testing.T) {
// 	tests := []struct {
// 		a, b, expected int
// 	}{
// 		{1, 2, 3},
// 		{-1, 0, 0},
// 		{0, 0, 0},
// 		{-1, -1, -2},
// 	}

// 	for _, test := range tests {
// 		t.Run(fmt.Sprintf("Add(%d, %d)", test.a, test.b), func(t *testing.T) {
// 			result := Add(test.a, test.b)
// 			if result != test.expected {
// 				t.Errorf("Add(%d, %d) = %d; want %d", test.a, test.b, result, test.expected)
// 			}
// 		})
// 	}
// }

// func TestAddTableDriven(t *testing.T) {
// 	tests := []struct {
// 		a, b, expected int
// 	}{
// 		{1, 2, 3},
// 		{-1, 1, 0},
// 		{0, 0, 2},
// 		{-1, -1, -2},
// 	}

// 	for _, test := range tests {
// 		result := Add(test.a, test.b)
// 		if result != test.expected {
// 			t.Errorf("Add(%d, %d) = %d; want %d", test.a, test.b, result, test.expected)
// 		}
// 	}
// }

// func TestAdd(t *testing.T) {
// 	result := Add(3, 2)
// 	expected := 3
// 	if result != expected {
// 		t.Errorf("Add(1, 2) = %d; want %d", result, expected)
// 	}
// }

// func TestAdd(t *testing.T) {
// 	result := Add(3, 2)
// 	expected := 3
// 	if result != expected {
// 		t.Errorf("Add(1, 2) = %d; want %d", result, expected)
// 	}
// }

