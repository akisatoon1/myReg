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
	fmt.Println(buildNFA(parseRegExp("a?((b+c)*?a)?b")).match("abbcbbccab") == true)
	fmt.Println(buildNFA(parseRegExp("a?((b+c)*?a)?b")).match("aab") == true)
	fmt.Println(buildNFA(parseRegExp("a?((b+c)*?a)?b")).match("abbcbbccabc") == false)
	fmt.Println(buildNFA(parseRegExp("a?((b+c)*?a)?b")).match("abbcbbccb") == false)
}
