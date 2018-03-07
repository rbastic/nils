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

func CountEmptyStrings(maybeEmpty []string) int {
	var count int
	for _, v := range maybeEmpty {
		if v == "" {
			count++
		}
	}
	return count
}

func RemoveEmptyStrings(someEmpty []string, count int) []string {
	noEmpty := make([]string, len(someEmpty)-count)
	j := 0
	for _, v := range someEmpty {
		if v != "" {
			noEmpty[j] = v
			j++
		}
	}
	return noEmpty
}

func RemoveEmptyStringsIfNeeded(maybeEmpty []string) []string {
	numEmpty := CountEmptyStrings(maybeEmpty)
	if numEmpty > 0 {
		return RemoveEmptyStrings(maybeEmpty, numEmpty)
	}
	return maybeEmpty
}
