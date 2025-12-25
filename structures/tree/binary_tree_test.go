package tree

import (
	"math/rand"
	"testing"
)

func TestAddSingleElement(t *testing.T) {

	tree := NewBinaryTree[int]()
	tree.Insert(10)

	if tree.root == nil {
		t.Fatal("Expected tree root should be initialized")
	}

	if tree.Count() != 1 {
		t.Fatal("Expected count for the tree should be 1")
	}
}

func TestAddDuplicateIgnored(t *testing.T) {
	tree := NewBinaryTree[int]()

	tree.InsertRange(10, 10, 10)

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

	tree.InsertRange(567, 785, 563, 123, 342, 3434, 0, 34, 10003, 389, 345, 232)
	count := tree.Count()

	if count != 12 {
		t.Fatal("Unexpected count")
	}
}

func TestDeleteNode(t *testing.T) {
	tests := []struct {
		name     string
		elements []int
		delete   int
		expected bool
	}{
		{"Should return ok for existing element", []int{5, 3, 7, 6, 8, 4, 2}, 2, true},
		{"Should return fail for missing element", []int{5, 3, 7, 6, 8, 4, 2}, 10, false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tree := NewBinaryTree[int]()
			tree.InsertRange(test.elements...)
			ok := tree.Delete(test.delete)

			if ok != test.expected {
				t.Fatal("Deletion should be expected to be sucessful")
			}
		})
	}
}

func TestDeleteRightLeafNodes(t *testing.T) {
	tests := []struct {
		name     string
		elements []int
		delete   []int
		expected bool
		count    int
	}{
		{"Should return ok for existing element", []int{5, 3, 7, 6, 8, 4, 2}, []int{2, 4, 6, 8}, true, 3},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			tree := NewBinaryTree[int]()
			tree.InsertRange(test.elements...)

			for _, element := range test.delete {
				ok := tree.Delete(element)

				if ok != test.expected {
					t.Fatal("Deletion should be expected to be sucessful")
				}
			}

			if tree.Count() != test.count {
				t.Fatal("Unexpected tree count value:", tree.Count())
			}
		})
	}
}

func BenchmarkAddRandom(b *testing.B) {

	b.ResetTimer()

	tree := NewBinaryTree[int]()
	for i := 0; b.Loop(); i++ {
		tree.Insert(rand.Int())
	}
}

func BenchmarkAddSorted(b *testing.B) {
	tree := NewBinaryTree[int]()

	b.ResetTimer()

	for i := 0; b.Loop(); i++ {
		tree.Insert(i)
	}
}
