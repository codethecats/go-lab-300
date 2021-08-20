package main

func main() {
	if Condition1 {
		Operation1
		if Condition2 {
			Operation2
		} else {
			Operation3
			Operation4
		}
	} else if Condition3 {
		if Condition4 {
			Operation5
		} else if Condition5 {
			Operation6
		} else {
			Operation7
		}
		Operation8
	}
	Operation9
}

