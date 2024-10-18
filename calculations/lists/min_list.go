package lists

// Erwartet eine Liste von Zahlen und gibt das Minimum zurück.
// Wenn die Liste leer ist, wird 0 zurückgegeben.
func MinList(nums []int) int {
	// TODO

	if len(nums) == 0 {
		return 0
	}

	m := nums[0]

	for i := 0; i < len(nums); i++ {
		if nums[i] < m {
			m = nums[i]
		}
	}

	return m

}
