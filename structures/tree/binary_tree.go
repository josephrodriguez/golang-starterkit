package tree

import "cmp"

type node[T cmp.Ordered] struct {
	value T
	left  *node[T]
	right *node[T]
}

type BinaryTree[T cmp.Ordered] struct {
	root *node[T]
}

func NewBinaryTree[T cmp.Ordered]() *BinaryTree[T] {
	return &BinaryTree[T]{
		root: nil,
	}
}

func (t *BinaryTree[T]) Add(element T) {

	if t.root == nil {
		t.root = &node[T]{value: element}
		return
	}

	for current := t.root; ; {
		switch {
		case element > current.value:
			if current.right == nil {
				current.right = &node[T]{value: element}
				return
			}
			current = current.right
		case element < current.value:
			if current.left == nil {
				current.left = &node[T]{value: element}
				return
			}
			current = current.left
		case element == current.value:
			return
		}
	}
}
