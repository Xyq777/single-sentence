package stringsX

func RemoveDuplicates(list []string) []string {
	valueMap := make(map[string]struct{})
	result := make([]string, 0)
	for _, value := range list {
		if _, ok := valueMap[value]; !ok {
			valueMap[value] = struct{}{}
			result = append(result, value)
		}

	}
	return result

}
