package main

func main() {
	if Condition1 {
		if Condition2 {
			Operation1
		} else if Condition3 {
			Operation2
		}
	} else {
		if Condition4 {
			Operation3
		} else if Condition5 {
			Operation4
		}
		Operation5
	}
	Operation6
}