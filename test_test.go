package main

import (
	"runtime"
	"testing"
)

func testLog(b bool, t *testing.T) {
	if !b {
		_, _, line, _ := runtime.Caller(1)
		t.Errorf("line: %v\n", line)
	}
}

func TestStack(t *testing.T) {
	op0 := &Operator{priority: 0, char: "0"}
	op1 := &Operator{priority: 0, char: "1"}
	op2 := &Operator{priority: 0, char: "2"}

	stack := Stack{
		items: []any{},
		len:   0,
	}

	stack.push(op0)
	stack.push(op1)
	testLog(stack.pop().(*Operator).char == "1", t)
	testLog(stack.lookTop().(*Operator).char == "0", t)
	stack.push(op2)
	testLog(stack.lookTop().(*Operator).char == "2", t)
	testLog(stack.pop().(*Operator).char == "2", t)
	testLog(stack.pop().(*Operator).char == "0", t)
	testLog(stack.pop() == nil, t)
	testLog(stack.lookTop() == nil, t)
}

func TestParseRegExp(t *testing.T) {
	testLog(parseRegExp("a?(a+b)*?b") == "aab+*?b?", t)
	testLog(parseRegExp("a*?(a+(b+c)?a*)") == "a*abc+a*?+?", t)
	testLog(parseRegExp("a?((a+b)*+c*?b)?(b+c)*") == "aab+*c*b?+?bc+*?", t)
}

func TestBuildNFA(t *testing.T) {
	testLog(buildNFA("").start.epsilonTransitions[0].isFinal == true, t)
	testLog(buildNFA("ab?").start.transition["a"].transition["b"].isFinal == true, t)

	nfa := buildNFA("ab+")
	testLog(nfa.start.epsilonTransitions[0].transition["a"].epsilonTransitions[0].isFinal == true, t)
	testLog(nfa.start.epsilonTransitions[1].transition["b"].epsilonTransitions[0].isFinal == true, t)
	testLog(nfa.start.epsilonTransitions[0].transition["a"].isFinal == false, t)
	testLog(nfa.start.epsilonTransitions[0].isFinal == false, t)

	nfa = buildNFA("a*")
	testLog(nfa.start.epsilonTransitions[1].isFinal == true, t)
	testLog(nfa.start.epsilonTransitions[0].transition["a"].epsilonTransitions[1].isFinal == true, t)
	testLog(nfa.start.epsilonTransitions[0].transition["a"].epsilonTransitions[0].transition["a"].epsilonTransitions[1].isFinal == true, t)
	testLog(nfa.start.epsilonTransitions[0].transition["a"].isFinal == false, t)
	testLog(nfa.start.epsilonTransitions[0].isFinal == false, t)
}

func TestMatch(t *testing.T) {
	testLog(buildNFA("ab?").match("ab") == true, t)
	testLog(buildNFA("ab?").match("aa") == false, t)
	testLog(buildNFA("ab?").match("abc") == false, t)
	testLog(buildNFA("ab+").match("a") == true, t)
	testLog(buildNFA("ab+").match("b") == true, t)
	testLog(buildNFA("ab+").match("bb") == false, t)
	testLog(buildNFA("ab+").match("c") == false, t)
	testLog(buildNFA("a*").match("") == true, t)
	testLog(buildNFA("a*").match("aaa") == true, t)
	testLog(buildNFA("a*").match("aaaabaaaa") == false, t)
}
