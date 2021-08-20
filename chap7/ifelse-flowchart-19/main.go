package main

func main() {
	if Condition1 {
		if Condition2 {
			Operation1
		} else {
			Operation2
		}
	} else if Condition3 {
		if Condition4 {
			Operation3
		} else {
			Operation4
			Operation5
		}
		Operation6
	} else {
		Operation7
		if Condition5 {
			Operation8
		}
		Operation9
	}
	Operation10
}