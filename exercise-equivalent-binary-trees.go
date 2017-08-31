package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk calls the helper function once, ensuring
// the channel gets closed appropriately.
func Walk(t *tree.Tree, ch chan int) {
	helper(t, ch)
	close(ch)
}

// helper walks the tree t sending all values
// from the tree to the channel ch.
func helper(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	helper(t.Left, ch)
	ch <- t.Value
	helper(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for l := range ch1 {
		if r := <-ch2; l != r {
			return false
		}
	}
	return true
}

func main() {
	tree1 := tree.New(121)
	tree2 := tree.New(12)
	fmt.Println(Same(tree1, tree2))
}
