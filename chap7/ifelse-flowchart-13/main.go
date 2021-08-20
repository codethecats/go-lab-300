package main

func main() {
	if Condition1 {
		if Condition2 {
			Operation1
		} else {
			Operation2
		}
	} else {
		Operation3
		if Condition3 {
			Operation4
		} else if Condition4 {
			Operation5
		} else {
			Operation6
		}
	}
	Operation7
}