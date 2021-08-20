package main

func main() {
	if Condition1 {
		if Condition2 {
			Operation1
		} else if Condition3 {
			Operation2
		} else {
			Operation3
		}
		Operation4
	} else {
		Operation5
		if Condition4 {
			Operation6
			Operation7
		} else {
			Operation8
		}
	}
	Operation9
}