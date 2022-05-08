package homework

func reverse(input []int64) (result []int64) {
	reverse := make([]int64, 0, len(input))
	for i := range input {
		reverse = append(reverse, input[len(input)-i-1])
	}
	return reverse
}
