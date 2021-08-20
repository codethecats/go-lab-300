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
			} else {
				Operation6
			}
		}
		Operation7
	}
	Operation8
}