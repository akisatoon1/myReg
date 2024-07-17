package main

import "fmt"

//
// *: Kleene Closure	↑
// ?: Concatenate		| priority
// +: Union (or)		|
// (): ...
//

func main() {
	fmt.Println("Integration Test:")
	fmt.Println(matchString("a?((b+c)*?a)?b", "abbcbbccab") == true)
	fmt.Println(matchString("a?((b+c)*?a)?b", "aab") == true)
	fmt.Println(matchString("a?((b+c)*?a)?b", "abbcbbccabc") == false)
	fmt.Println(matchString("a?((b+c)*?a)?b", "abbcbbccb") == false)
}

func matchString(regexp string, str string) bool {
	return compile(regexp).match(str)
}

func compile(regexp string) *NFA {
	return buildNFA(parseRegExp(regexp))
}
