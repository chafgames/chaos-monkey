package server

import "log"

func getFreeMonkey() (int, bool) {
	for index, monkeyState := range myState.Monkeys {
		if monkeyState.Active == false {
			log.Printf("%d is the first dead monkey", index)

			return index, true
		}
		log.Printf("leaving active monkey %d alone", index)
	}
	return -1, false
}
