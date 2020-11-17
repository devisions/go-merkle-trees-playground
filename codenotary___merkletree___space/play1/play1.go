package main

import (
	"fmt"

	"devisions.org/go-merkle-trees-playground/codenotary___merkletree___space"
	"github.com/codenotary/merkletree"
)

func main() {

	s := codenotary___merkletree___space.NewLocalMemStore()
	fmt.Printf(" initially, tree's width is %d.\n", s.Width())

	merkletree.Append(s, []byte("d0"))
	merkletree.Append(s, []byte("d1"))
	merkletree.Append(s, []byte("d2"))

	fmt.Printf("\n--- printing the tree ---\n\n")
	if lms, ok := s.(*codenotary___merkletree___space.LocalMemStore); ok {
		lms.Print()
	}
}
