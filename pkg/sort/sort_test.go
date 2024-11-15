package sort_test

import (
	"math/rand"
	"testing"
	"go-sort-bench/pkg/sort"
)

func BenchmarkBubbleSort(b *testing.B) {
	b.Run("small arrays", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10, 10)
			b.StartTimer()
			sort.BubbleSort(ar)
			b.StopTimer()
		}
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(100, 1000)
			b.StartTimer()
			sort.BubbleSort(ar)
			b.StopTimer()
		}
	})

	b.Run("big arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice(10000, 100000)
			b.StartTimer()
			sort.BubbleSort(ar)
			b.StopTimer()
		}
	})
}

func generateSlice(max, size int) []int {
	ar := make([]int, size)
	for i := range ar {
		ar[i] = rand.Intn(max*2) - max
	}

	return ar
}


func BenchmarkSelectionSort(b *testing.B) {
	b.Run("small arrays", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice2(10, 10)
			b.StartTimer()
			sort.SelectionSort(ar)
			b.StopTimer()
		}
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice2(100, 1000)
			b.StartTimer()
			sort.SelectionSort(ar)
			b.StopTimer()
		}
	})

	b.Run("big arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice2(10000, 100000)
			b.StartTimer()
			sort.SelectionSort(ar)
			b.StopTimer()
		}
	})
}

func generateSlice2(max, size int) []int {
	ar := make([]int, size)
	for i := range ar {
		ar[i] = rand.Intn(max*2) - max
	}

	return ar
}


func BenchmarkInsertionSort(b *testing.B) {
	b.Run("small arrays", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice3(10, 10)
			b.StartTimer()
			sort.InsertionSort(ar)
			b.StopTimer()
		}
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice3(100, 1000)
			b.StartTimer()
			sort.InsertionSort(ar)
			b.StopTimer()
		}
	})

	b.Run("big arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice3(10000, 100000)
			b.StartTimer()
			sort.InsertionSort(ar)
			b.StopTimer()
		}
	})
}

func generateSlice3(max, size int) []int {
	ar := make([]int, size)
	for i := range ar {
		ar[i] = rand.Intn(max*2) - max
	}

	return ar
}


func BenchmarkMergeSort(b *testing.B) {
	b.Run("small arrays", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice4(10, 10)
			b.StartTimer()
			sort.MergeSort(ar)
			b.StopTimer()
		}
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice4(100, 1000)
			b.StartTimer()
			sort.MergeSort(ar)
			b.StopTimer()
		}
	})

	b.Run("big arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice4(10000, 100000)
			b.StartTimer()
			sort.MergeSort(ar)
			b.StopTimer()
		}
	})
}

func generateSlice4(max, size int) []int {
	ar := make([]int, size)
	for i := range ar {
		ar[i] = rand.Intn(max*2) - max
	}

	return ar
}


func BenchmarkQuickSort(b *testing.B) {
	b.Run("small arrays", func(b *testing.B) {
		b.ReportAllocs()
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice5(10, 10)
			b.StartTimer()
			sort.QuickSort(ar)
			b.StopTimer()
		}
	})

	b.Run("middle arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice5(100, 1000)
			b.StartTimer()
			sort.QuickSort(ar)
			b.StopTimer()
		}
	})

	b.Run("big arrays", func(b *testing.B) {
		b.StopTimer()
		for i := 0; i < b.N; i++ {
			ar := generateSlice5(10000, 100000)
			b.StartTimer()
			sort.QuickSort(ar)
			b.StopTimer()
		}
	})
}

func generateSlice5(max, size int) []int {
	ar := make([]int, size)
	for i := range ar {
		ar[i] = rand.Intn(max*2) - max
	}

	return ar
}