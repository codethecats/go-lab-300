package main

func main() {
	if Condition1 {
		Operation1
		if Condition2 {
			Operation2
			Operation3
		} else if Condition3 {
			Operation4
			Operation5
		} else {
			Operation6
			Operation7
		}
	} else if Condition4 {
		if Condition5 {
			Operation8
		} else {
			Operation9
		}
		Operation10
	} else {
		Operation11
		if Condition6 {
			Operation12
		} else {
			Operation13
			Operation14
		}
		Operation15
	}
	Operation16
	Operation17
}