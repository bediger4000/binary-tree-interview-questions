package tree

import "fmt"

// Balanced decides whether or not its argument is height-balanced. A
// height-balanced binary tree can be defined as one in which the
// heights of the two subtrees of any node never differ by more than
// one.
// This has to be part of package tree so it can use leftChild(),
// rightChild(), isNil() methods.
func Balanced(node Node) bool {
	if node.isNil() {
		return true
	}

	leftDepth := FindDepth(node.leftChild(), 0)
	fmt.Printf("Left depth %d\n", leftDepth)
	rightDepth := FindDepth(node.rightChild(), 0)
	fmt.Printf("Right depth %d\n", rightDepth)

	depthDiff := leftDepth - rightDepth

	if depthDiff >= -1 && depthDiff <= 1 {
		return true
	}

	return false
}