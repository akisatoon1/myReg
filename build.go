// https://deniskyashif.com/2019/02/17/implementing-a-regular-expression-engine/

package main

type State struct {
	isFinal            bool
	transition         map[string]*State
	epsilonTransitions []*State
}

type NFA struct {
	start *State
	end   *State
}

// build automaton for match string
func buildNFA(parsedRegExp string) *NFA {
	if parsedRegExp == "" {
		return turnIsFinalToTrue(fromEpsilon())
	}

	stack := &Stack{}
	for i := 0; i < len(parsedRegExp); i++ {
		switch symbol := string(parsedRegExp[i]); symbol {

		case SymbolClosure:
			stack.push(closure(stack.pop().(*NFA)))

		case SymbolConcat:
			right, left := stack.pop().(*NFA), stack.pop().(*NFA)
			stack.push(concat(left, right))

		case SymbolUnion:
			right, left := stack.pop().(*NFA), stack.pop().(*NFA)
			stack.push(union(left, right))

		default:
			stack.push(fromSymbol(symbol))
		}
	}

	return turnIsFinalToTrue(stack.pop().(*NFA))
}

//
// create NFA
//

func closure(nfa *NFA) *NFA {
	start, end := initialState(), initialState()

	start.epsilonTransitions = []*State{nfa.start, end}
	nfa.end.epsilonTransitions = []*State{nfa.start, end}

	return &NFA{start, end}
}

func concat(leftNfa *NFA, rightNfa *NFA) *NFA {
	*leftNfa.end = *rightNfa.start
	return &NFA{leftNfa.start, rightNfa.end}
}

func union(leftNfa *NFA, rightNfa *NFA) *NFA {
	start, end := initialState(), initialState()

	start.epsilonTransitions = []*State{leftNfa.start, rightNfa.start}
	leftNfa.end.epsilonTransitions = []*State{end}
	rightNfa.end.epsilonTransitions = []*State{end}

	return &NFA{start, end}
}

func fromSymbol(symbol string) *NFA {

	start, end := initialState(), initialState()
	start.transition[symbol] = end

	return &NFA{start, end}
}

func fromEpsilon() *NFA {

	start, end := initialState(), initialState()
	start.epsilonTransitions = []*State{end}

	return &NFA{start, end}
}

// initial state transition map to use
func initialState() *State {
	return &State{transition: make(map[string]*State)}
}

// only last state is final state
func turnIsFinalToTrue(nfa *NFA) *NFA {
	nfa.end.isFinal = true
	return nfa
}
