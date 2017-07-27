package nils

func CountNils(maybeNils []interface{}) int {
	var count int
	for _, v := range maybeNils {
		if v == nil {
			count++
		}
	}
	return count
}

func RemoveNils(someNils []interface{}, count int) []interface{} {
	noNils := make([]interface{}, len(someNils)-count)
	j := 0
	for _, v := range someNils {
		if v != nil {
			noNils[j] = v
			j++
		}
	}
	return noNils
}

func RemoveNilsIfNeeded(maybeNils []interface{}) []interface{} {
	numNils := CountNils(maybeNils)
	if numNils > 0 {
		return RemoveNils(maybeNils, numNils)
	}
	return maybeNils
}
