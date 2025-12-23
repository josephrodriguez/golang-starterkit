package tree

import "cmp"

type node[T cmp.Ordered] struct {
	value T
	left  *node[T]
	right *node[T]
}

type BinaryTree[T cmp.Ordered] struct {
	root  *node[T]
	count int
}

func NewBinaryTree[T cmp.Ordered]() *BinaryTree[T] {
	return &BinaryTree[T]{
		root: nil,
	}
}

func (t *BinaryTree[T]) Add(element T) {

	if t.root == nil {
		t.root = &node[T]{value: element}
		t.count++
		return
	}

	for current := t.root; ; {
		switch {
		case element > current.value:
			if current.right == nil {
				current.right = &node[T]{value: element}
				t.count++
				return
			}
			current = current.right
		case element < current.value:
			if current.left == nil {
				current.left = &node[T]{value: element}
				t.count++
				return
			}
			current = current.left
		case element == current.value:
			return
		}
	}
}

func (t *BinaryTree[T]) AddRange(elements ...T) {

	for _, element := range elements {
		t.Add(element)
	}
}

func (t *BinaryTree[T]) Delete(element T) bool {

	if t.root == nil {
		return false
	}

	for current := t.root; current != nil; {
		return false
	}

	return false
}

func (t *BinaryTree[T]) Search(element T) bool {

	if t.root == nil {
		return false
	}

	for current := t.root; current != nil; {
		switch {
		case element > current.value:
			current = current.right
		case element < current.value:
			current = current.left
		default:
			return true
		}
	}

	return false
}

func (t *BinaryTree[T]) Min() (T, bool) {

	current := t.root

	for current.left != nil {
		current = current.left
	}

	return current.value, current != nil
}

func (t *BinaryTree[T]) Max() (T, bool) {

	current := t.root

	for current.right != nil {
		current = current.right
	}

	return current.value, current != nil
}

func (t *BinaryTree[T]) IsEmpty() bool {
	return t.count == 0
}

func (t *BinaryTree[T]) Count() int {
	return t.count
}

func (t *BinaryTree[T]) Height() int {
	return 0
}
