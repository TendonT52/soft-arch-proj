package utils

func IsElementInStringArray(s string, arr *[]string) bool {
	for _, a := range *arr {
		if a == s {
			return true
		}
	}

	return false
}

func IsItemInArray(s int64, arr *[]int64) bool {
	for _, a := range *arr {
		if a == s {
			return true
		}
	}

	return false
}

func CheckArrayEqual(s1 *[]string, s2 *[]string) bool {
	if len(*s1) != len(*s2) {
		return false
	}

	for _, s := range *s1 {
		if IsElementInStringArray(s, s2) == false {
			return false
		}
	}

	return true
}
