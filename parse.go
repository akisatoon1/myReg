// https://blog.cernera.me/converting-regular-expressions-to-postfix-notation-with-the-shunting-yard-algorithm/

package main

// handle operator symbol
// more priority, smaller priority(int)
type Operator struct {
	priority int
	char     string
}

//
// parse
//

// parse regexp string to postfix notation for build NFA
func parseRegExp(regularExp string) string {
	closure := &Operator{priority: 1, char: SymbolClosure}
	concat := &Operator{priority: 2, char: SymbolConcat}
	union := &Operator{priority: 3, char: SymbolUnion}
	leftParen := &Operator{priority: 4, char: "("}

	stack := &Stack{}
	queue := ""
	for i := 0; i < len(regularExp); i++ {
		switch symbol := string(regularExp[i]); symbol {
		case SymbolClosure:
			handleInputOperator(closure, stack, &queue)
		case SymbolConcat:
			handleInputOperator(concat, stack, &queue)
		case SymbolUnion:
			handleInputOperator(union, stack, &queue)
		case "(":
			handleLeftParen(leftParen, stack)
		case ")":
			handleRightParen(stack, &queue)
		default:
			appendSymbolToQueue(symbol, &queue)
		}
	}
	appendRestOfStackToQueue(stack, &queue)

	return queue
}

func appendSymbolToQueue(symbol string, queue *string) {
	*queue += symbol
}

func handleInputOperator(inputOpe *Operator, stack *Stack, queue *string) {
	for ope := stack.lookTop(); ope != nil && ope.(*Operator).priority <= inputOpe.priority; ope = stack.lookTop() {
		appendSymbolToQueue(stack.pop().(*Operator).char, queue)
	}
	stack.push(inputOpe)
}

func handleLeftParen(op *Operator, stack *Stack) {
	stack.push(op)
}

func handleRightParen(stack *Stack, queue *string) {
	for ope := stack.pop(); ope != nil && ope.(*Operator).char != "("; ope = stack.pop() {
		appendSymbolToQueue(ope.(*Operator).char, queue)
	}
}

func appendRestOfStackToQueue(stack *Stack, queue *string) {
	for ope := stack.pop(); ope != nil; ope = stack.pop() {
		appendSymbolToQueue(ope.(*Operator).char, queue)
	}
}
