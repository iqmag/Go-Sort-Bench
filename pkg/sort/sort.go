package sort

func BubbleSort(ar []int) {
	n := len(ar)
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if ar[j] > ar[j+1] {
				ar[j], ar[j+1] = ar[j+1], ar[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
}

func SelectionSort(ar []int) {
	for i := 0; i < len(ar)-1; i++ {
		minIndex := i
		for j := i + 1; j < len(ar); j++ {
			if ar[j] < ar[minIndex] {
				minIndex = j
			}
		}
		ar[i], ar[minIndex] = ar[minIndex], ar[i]
	}
}

func InsertionSort(ar []int) {
	if len(ar) < 2 {
		return
	}
	for i := 1; i < len(ar); i++ {
		for j := i; j > 0 && ar[j-1] > ar[j]; j-- {
			ar[j-1], ar[j] = ar[j], ar[j-1]
		}
	}
}

func MergeSort(ar []int) []int {
	if len(ar) < 2 {
		return ar
	}

	middle := len(ar) / 2

	sortedAr := make([]int, 0, len(ar))
	left, right := MergeSort(ar[:middle]), MergeSort(ar[middle:])

	var i, j = 0, 0
	for i < len(left) && j < len(right) {
		if left[i] > right[j] {
			sortedAr = append(sortedAr, right[j])
			j++
		} else {
			sortedAr = append(sortedAr, left[i])
			i++
		}
	}

	sortedAr = append(sortedAr, left[i:]...)
	sortedAr = append(sortedAr, right[j:]...)

	return sortedAr
}

func QuickSort(ar []int) {
	if len(ar) < 2 {
		return
	}

	left, right := 0, len(ar)-1
	pivotIndex := 0

	ar[pivotIndex], ar[right] = ar[right], ar[pivotIndex]

	for i := 0; i < len(ar); i++ {
		if ar[i] < ar[right] {
			ar[i], ar[left] = ar[left], ar[i]
			left++
		}
	}

	ar[left], ar[right] = ar[right], ar[left]

	QuickSort(ar[:left])
	QuickSort(ar[left+1:])
}