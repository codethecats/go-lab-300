package main

func main() {
	if Condition1 {
		if Condition2 {
			Operation1
		}
	} else if Condition3 {
		if Condition4 {
			Operation2
		}
	} else {
		if Condition5 {
			Operation3
		}
		Operation4
	}
	Operation5
}