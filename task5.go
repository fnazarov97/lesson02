package main

//min(int, min(int, min(int, int)))
func recur(a int) string {
	var str string
	if a == 2 {
		str = "min(int, int)"
		return str
	}
	str = recur(a - 1)
	result := "min(int, " + str + ")"
	return result
}
