package Solutions

func SumValuesByKey(m map[string][]int) map[string]int {

	result := make(map[string]int)
	for key, values := range m {
		sum := 0
		for _, value := range values {
			sum += value
		}
		result[key] = sum
	}

	return result
}
