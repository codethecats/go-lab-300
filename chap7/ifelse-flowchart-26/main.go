package main

func main() {
	if Condition1 {
		if Condition2 {
			Operation1
		} else if Condition3 {
			Operation2
		}
		Operation3
	} else if Condition4 {
		Operation4
		if Condition5 {
			Operation5
		} else {
			Operation6
		}
		Operation7
	}
	Operation8
}