// https://deniskyashif.com/2019/02/17/implementing-a-regular-expression-engine/

package main

//
// implement StateS
//

type StateS []*State

func (stateS StateS) isIncludeFinal() bool {
	for _, state := range stateS {
		if state.isFinal {
			return true
		}
	}
	return false
}

func (stateS StateS) isInclude(state *State) bool {
	for _, st := range stateS {
		if st == state {
			return true
		}
	}
	return false
}

// match given string using automaton
func (nfa *NFA) match(str string) bool {
	currentStateS := firstStateS(nfa)
	for i := 0; i < len(str); i++ {
		currentStateS = nextStateS(string(str[i]), currentStateS)
	}
	return currentStateS.isIncludeFinal()
}

func firstStateS(nfa *NFA) StateS {
	stateS := StateS{}
	addNextState(nfa.start, &stateS)
	return stateS
}

func nextStateS(symbol string, currentStateS StateS) StateS {
	stateS := StateS{}
	for _, state := range currentStateS {
		if transitted := state.transition[symbol]; transitted != nil {
			addNextState(transitted, &stateS) // stateSでの被りは無視
		}
	}
	return stateS
}

func addNextState(state *State, next *StateS) {
	visited := StateS{}
	if len(state.epsilonTransitions) == 0 {
		*next = append(*next, state)
	} else {
		for _, st := range state.epsilonTransitions {
			if !visited.isInclude(st) {
				visited = append(visited, st)
				addNextState(st, next)
			}
		}
	}
}
