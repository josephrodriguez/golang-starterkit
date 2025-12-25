package tree

import "cmp"

// node represents a single element in the binary search tree.
//
// Each node holds a value and pointers to its left and right child nodes.
type node[T cmp.Ordered] struct {
	value       T
	left, right *node[T]
}

// isLeaf reports whether the node has no children.
func (n *node[T]) isLeaf() bool {
	return n.left == nil && n.right == nil
}

func (n *node[T]) findMinimum(p *node[T]) (*node[T], *node[T]) {
	var current, parent *node[T]

	for current.left != nil {
		parent = current
		current = current.left
	}

	return current, parent
}

// BinaryTree represents a binary search tree (BST) of ordered elements.
//
// Elements are arranged such that values smaller than a node are stored
// in the left subtree, and values greater than a node are stored in the
// right subtree.
type BinaryTree[T cmp.Ordered] struct {
	root  *node[T]
	count int
}

// NewBinaryTree creates and returns a new empty BinaryTree.
//
// The returned tree is initialized and ready for use.
func NewBinaryTree[T cmp.Ordered]() *BinaryTree[T] {
	return &BinaryTree[T]{
		root: nil,
	}
}

// Inserts an element into the binary search tree.
//
// If the tree is empty, the element becomes the root node.
// If the element already exists in the tree, the operation
// has no effect and the tree remains unchanged.
func (t *BinaryTree[T]) Insert(element T) {
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

// Inserts multiple elements into the binary search tree.
//
// Each element is added individually using the Add method.
func (t *BinaryTree[T]) InsertRange(elements ...T) {
	for _, element := range elements {
		t.Insert(element)
	}
}

// Delete attempts to remove an element from the binary search tree.
//
// The current implementation always returns false and does not
// modify the tree.
func (t *BinaryTree[T]) Delete(element T) bool {
	if t.root == nil {
		return false
	}

	var current, parent *node[T] = t.root, nil
	for current != nil {
		switch {
		case element > current.value:
			parent = current
			current = current.right
		case element < current.value:
			parent = current
			current = current.left
		default:
			if current.left == nil {
				t.transplant(parent, current, current.right)

			} else if current.right == nil {
				t.transplant(parent, current, current.left)
			} else {
				minNode, minParent := current.right.findMinimum(current)

				if minParent != current {
					t.transplant(minParent, minNode, minNode.right)
					minNode.right = current.right
				}

				t.transplant(parent, current, minNode)
				minNode.left = current.left
			}

			t.count--
			return true
		}
	}

	return false
}

// Search reports whether the specified element exists in the tree.
//
// The search traverses the tree according to binary search tree rules.
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

// Min returns the minimum value stored in the tree.
//
// The second return value reports whether a value was found.
func (t *BinaryTree[T]) Min() (T, bool) {
	curr := t.root

	for curr.left != nil {
		curr = curr.left
	}

	return curr.value, curr != nil
}

// Max returns the maximum value stored in the tree.
//
// The second return value reports whether a value was found.
func (t *BinaryTree[T]) Max() (T, bool) {
	curr := t.root

	for curr.right != nil {
		curr = curr.right
	}

	return curr.value, curr != nil
}

// IsEmpty reports whether the tree contains no elements.
func (t *BinaryTree[T]) IsEmpty() bool {
	return t.count == 0
}

// Count returns the number of elements currently stored in the tree.
func (t *BinaryTree[T]) Count() int {
	return t.count
}

// Height returns the height of the binary tree.
//
// The current implementation always returns zero.
func (t *BinaryTree[T]) Height() int {
	return 0
}

func (t *BinaryTree[T]) InOrder(visit func(T) bool) {
	var traverse func(*node[T]) bool

	traverse = func(node *node[T]) bool {
		if node == nil {
			return false
		}

		if node.left != nil && !traverse(node.left) {
			return false
		}

		if !visit(node.value) {
			return false
		}

		if node.right != nil && !traverse(node.right) {
			return false
		}

		return true
	}

	traverse(t.root)
}

func (t *BinaryTree[T]) transplant(parent *node[T], u *node[T], v *node[T]) {
	if parent == nil {
		t.root = v
	} else if parent.left == u {
		parent.left = v
	} else {
		parent.right = v
	}
}
