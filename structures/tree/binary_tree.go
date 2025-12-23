package tree

import "cmp"

type BinaryNode[T cmp.Ordered] struct {
	value T
	left  *BinaryNode[T]
	right *BinaryNode[T]
}

type BinaryTree[T cmp.Ordered] struct {
	root *BinaryNode[T]
}
