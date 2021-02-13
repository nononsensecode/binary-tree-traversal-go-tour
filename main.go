package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk traverse a binary tree in in-order fashion
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}

	// This is in-order traversal. If you want pre order place the `assignment to channel line`
	// before the first Walk, and if you want post order place that line after the second Walk
	Walk(t.Left, ch)
	ch <- t.Value
	Walk(t.Right, ch)
}

// Same matches two binary trees to check whether they are same
// This trees are of same length and are also sorted.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for i := 0; i < 10; i++ {
		if <-ch1 != <-ch2 {
			return false
		}
	}

	return true
}

func main() {
	t := tree.New(10)
	isSame := Same(t, t)
	fmt.Println(isSame)

	t1 := tree.New(5)
	t2 := tree.New(3)
	isSame = Same(t1, t2)
	fmt.Println(isSame)
}