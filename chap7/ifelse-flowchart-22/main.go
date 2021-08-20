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
	} else if Condition4 {
		if Condition5 {
			Operation4
			Operation5
		} else if Condition6 {
			Operation6
		} else {
			Operation7
		}
		Operation8
		Operation9
	} else {
		Operation10
		if Condition7 {
			Operation11
		} else {
			Operation12
			Operation13
		}
		Operation14
	}
	Operation15
}