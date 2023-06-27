package main

import (
	"fmt"
	"strings"
)

type TransformFunc func(s string) string

func Uppercase(s string) string {
	return strings.ToUpper(s)
}

// composing the decorator function
// wow nice concept to avoid hard coding
// TransformFunc only have one parameter but to add prefixer we made TransformFunc as return type and do necessary transform and return according to its structure
// should learn this pattern more.. it will be usefull

func Prefixer(prefix string) TransformFunc {
	return func(s string) string {
		return prefix + s
	}
}

func tranformString(s string, fn TransformFunc) string {
	return fn(s)
}

func main() {
	fmt.Println((tranformString("World", Uppercase)))
	fmt.Println((tranformString("Hello", Prefixer("FOO_"))))
	fmt.Println((tranformString("Hello", Prefixer("BAR_"))))
}
