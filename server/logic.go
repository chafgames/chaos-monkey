package server

func getFreeMonkey() (int, bool) {
	for index, monkeyState := range myState.Monkeys {
		if monkeyState.Active == false {
			return index, true
		}
	}
	return -1, false
}
