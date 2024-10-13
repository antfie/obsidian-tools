package main

func isInArray(in string, array []string) bool {
	for _, item := range array {
		if item == in {
			return true
		}
	}

	return false
}
