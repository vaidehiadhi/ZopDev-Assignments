package Solutions

func SliceToMap(slice []string) map[string]int {
	result := make(map[string]int)
	
	for index, element := range slice {
		result[element] = index
	}

	return result
}
