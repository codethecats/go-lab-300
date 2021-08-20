package main

func main() {
	if Condition1 {
		Operation1
		if Condition2 {
			Operation2
		} else {
			Operation3
		}
		Operation4
		Operation5
	} else if Condition3 {
		Operation6
		Operation7
		if Condition4 {
			Operation8
		} else if Condition5 {
			Operation9
		}
		Operation10
		Operation11
	}
	Operation12
}