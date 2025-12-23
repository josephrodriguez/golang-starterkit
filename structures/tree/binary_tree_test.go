package tree

import (
	"math/rand"
	"testing"
)

func TestAddSingleElement(t *testing.T) {

	tree := NewBinaryTree[int]()
	tree.Add(10)

	if tree.root == nil {
		t.Fatal("Expected tree root should be initialized")
	}

	if tree.Count() != 1 {
		t.Fatal("Expected count for the tree should be 1")
	}
}

func TestAddDuplicateIgnored(t *testing.T) {
	tree := NewBinaryTree[int]()

	tree.Add(10)
	tree.Add(10)
	tree.Add(10)

	if tree.root == nil {
		t.Fatal("root should not be nil")
	}

	if tree.root.left != nil || tree.root.right != nil {
		t.Fatal("duplicate insert should not create child nodes")
	}

	if tree.count != 1 {
		t.Fatal("Expected count should be 1")
	}
}

func TestAddRangeElements(t *testing.T) {

	tree := NewBinaryTree[int]()

	tree.AddRange(567, 785, 563, 123, 342, 3434, 0, 34, 10003, 389, 345, 232)
	count := tree.Count()

	if count != 12 {
		t.Fatal("Unexpected count")
	}
}

func BenchmarkAddRandom(b *testing.B) {

	b.ResetTimer()

	tree := NewBinaryTree[int]()
	for i := 0; b.Loop(); i++ {
		tree.Add(rand.Int())
	}
}

func BenchmarkAddSorted(b *testing.B) {
	tree := NewBinaryTree[int]()

	b.ResetTimer()

	for i := 0; b.Loop(); i++ {
		tree.Add(i)
	}
}
