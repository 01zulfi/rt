package utils

func Belongs[T comparable](item T, array []T) bool {
	for _, element := range array {
		if element == item {
			return true
		}
	}
	return false
}
