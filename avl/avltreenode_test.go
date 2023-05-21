package avl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// helper to build a tree by inserting values in order
func buildTree(vals ...int) *AVLNode[int] {
	var root *AVLNode[int]
	for _, v := range vals {
		if root == nil {
			root = NewAVLNode(v, nil, nil)
		} else {
			root = root.insert(v)
		}
	}
	return root
}

func TestInsertAndSearch(t *testing.T) {
	vals := []int{50, 20, 70, 10, 30, 60, 80}
	root := buildTree(vals...)
	for _, v := range vals {
		assert.True(t, root.search(v))
	}
	assert.False(t, root.search(999))
}

func TestFindMinMax(t *testing.T) {
	root := buildTree(5, 3, 8, 1, 4, 7, 9)
	assert.Equal(t, 1, root.findMin().GetData())
	assert.Equal(t, 9, root.findMax().GetData())
}

func TestSize(t *testing.T) {
	vals := []int{15, 10, 20, 8, 12, 17, 25}
	root := buildTree(vals...)
	assert.Equal(t, len(vals), root.Size())
}

func TestRemove(t *testing.T) {
	root := buildTree(5, 3, 7, 2, 4)
	root = root.remove(3)
	assert.False(t, root.search(3))
	assert.Equal(t, 4, root.Size())

	root = root.remove(5)
	assert.False(t, root.search(5))
	assert.Equal(t, 3, root.Size())
}

func TestHeightAndBalance(t *testing.T) {
	// manually link a small balanced tree:   2
	//                                     /   \
	//                                    1     3
	left := NewAVLNode(1, nil, nil)
	right := NewAVLNode(3, nil, nil)
	root := NewAVLNode(2, left, right)

	root.updateHeight()
	assert.Equal(t, 1, root.GetHeight())
	assert.Equal(t, 0, root.GetBalance())
}
