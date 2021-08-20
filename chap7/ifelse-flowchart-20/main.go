package main

func main() {
	if Condition1 {
		Operation1
		if Condition2 {
			Operation2
		}
	} else if Condition3 {
		if Condition4 {
			Operation3
		} else {
			Operation4
		}
		Operation5
		Operation6
	} else {
		Operation7
		if Condition5 {
			Operation8
			Operation9
		} else {
			Operation10
		}
	}
	Operation11
}