package main

func main() {
	if Condition1 {
		Operation1
		Operation2
		if Condition2 {
			Operation3
		} else if Condition3 {
			Operation4
			Operation5
		}
		Operation6
		Operation7
	} else if Condition4 {
		if Condition5 {
			Operation8
			Operation9
		} else {
			if Condition6 {
				Operation10
			} else if Condition7 {
				Operation11
			}
		}
	}
	Operation12
	Operation13
}