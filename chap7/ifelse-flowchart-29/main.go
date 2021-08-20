package main

func main() {
	if Condition1 {
		if Condition2 {
			Operation1
			Operation2
			if Condition3 {
				Operation3
			}
		} else {
			Operation4
			if Condition4 {
				Operation5
			}
		}
		Operation6
	} else {
		Operation7
		Operation8
		if Condition5 {
			if Condition6 {
				Operation9
			} else {
				Operation10
			}
			Operation11
		} else {
			Operation12
			if Condition7 {
				Operation13
			}
		}
	}
	Operation14
}