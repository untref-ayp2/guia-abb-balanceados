package avl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAVLTreeEmptyTree(t *testing.T) {
	tr := NewAVLTree[int]()
	assert.True(t, tr.IsEmpty())
	assert.Equal(t, 0, tr.Size())

	_, err1 := tr.FindMin()
	_, err2 := tr.FindMax()
	assert.Error(t, err1)
	assert.Error(t, err2)
}

func TestAVLTreeInsertAndSearch(t *testing.T) {
	tr := NewAVLTree[int]()
	vals := []int{50, 20, 70, 10, 30, 60, 80}
	for _, v := range vals {
		tr.Insert(v)
	}

	for _, v := range vals {
		assert.True(t, tr.Search(v))
	}
	assert.False(t, tr.Search(999))

	assert.False(t, tr.IsEmpty())
	assert.Equal(t, len(vals), tr.Size())
}

func TestAVLTreeFindMinMax(t *testing.T) {
	tr := NewAVLTree[int]()
	for _, v := range []int{5, 3, 8, 1, 4, 7, 9} {
		tr.Insert(v)
	}

	min, errMin := tr.FindMin()
	max, errMax := tr.FindMax()
	assert.NoError(t, errMin)
	assert.NoError(t, errMax)
	assert.Equal(t, 1, min)
	assert.Equal(t, 9, max)
}

func TestAVLTreeRemove(t *testing.T) {
	tr := NewAVLTree[int]()
	for _, v := range []int{5, 3, 7, 2, 4} {
		tr.Insert(v)
	}

	tr.Remove(3)
	assert.False(t, tr.Search(3))
	assert.Equal(t, 4, tr.Size())

	tr.Remove(5)
	assert.False(t, tr.Search(5))
	assert.Equal(t, 3, tr.Size())
}

func TestAVLTreeClear(t *testing.T) {
	tr := NewAVLTree[int]()
	for _, v := range []int{1, 2, 3} {
		tr.Insert(v)
	}
	tr.Clear()
	assert.True(t, tr.IsEmpty())
	assert.Equal(t, 0, tr.Size())
}

func TestAVLTreeHeightAndBalance(t *testing.T) {
	tr := NewAVLTree[int]()
	tr.Insert(2)
	tr.Insert(1)
	tr.Insert(3)
	// after inserting 2,1,3 the tree is perfectly balanced
	assert.Equal(t, 1, tr.GetHeight())
	assert.Equal(t, 0, tr.GetBalance())
}

func TestAVLTreeString(t *testing.T) {
	tr := NewAVLTree[int]()
	tr.Insert(42)
	// String() prints only root.data via node.string()
	assert.Equal(t, "42", tr.String())
}
