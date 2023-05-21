package binarytree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestINTERNALNuevoNodo(t *testing.T) {
	nodo := NewBinaryNode(5, nil, nil)
	assert.Equal(t, 5, nodo.GetData())
}

//      4
//     / \
//    2   5
//   / \
//  1   3

func TestINTERNALSize(t *testing.T) {
	nodo1 := NewBinaryNode(1, nil, nil)
	nodo3 := NewBinaryNode(3, nil, nil)
	nodo2 := NewBinaryNode(2, nodo1, nodo3)
	nodo5 := NewBinaryNode(5, nil, nil)
	nodo4 := NewBinaryNode(4, nodo2, nodo5)
	raiz := nodo4
	assert.Equal(t, 5, raiz.Size())
}

func TestINTERNALHeightOnEmptyTree(t *testing.T) {
	nodo := NewBinaryNode(1, nil, nil)
	assert.Equal(t, 0, nodo.Height())
}

func TestINTERNALHeight(t *testing.T) {
	nodo1 := NewBinaryNode(1, nil, nil)
	nodo3 := NewBinaryNode(3, nil, nil)
	nodo2 := NewBinaryNode(2, nodo1, nodo3)
	nodo5 := NewBinaryNode(5, nil, nil)
	nodo4 := NewBinaryNode(4, nodo2, nodo5)
	raiz := nodo4
	assert.Equal(t, 2, raiz.Height())
}

func TestINTERNALChildren(t *testing.T) {
	izq := NewBinaryNode(1, nil, nil)
	der := NewBinaryNode(3, nil, nil)
	raiz := NewBinaryNode(2, izq, der)

	assert.Equal(t, izq, raiz.GetLeft())
	assert.Equal(t, der, raiz.GetRight())
}

func TestIsAVLEmptyTree(t *testing.T) {
	var n *BinaryNode[int]
	assert.True(t, n.IsAVL())
}

func TestIsAVLBalancedTree(t *testing.T) {
	left := NewBinaryNode(1, nil, nil)
	right := NewBinaryNode(3, nil, nil)
	root := NewBinaryNode(2, left, right)
	assert.True(t, root.IsAVL())
}

func TestIsAVLUnbalancedLeftHeavy(t *testing.T) {
	// left-left heavy: 1 -> left -> 2 -> left -> 3
	n3 := NewBinaryNode(3, nil, nil)
	n2 := NewBinaryNode(2, n3, nil)
	root := NewBinaryNode(1, n2, nil)
	assert.False(t, root.IsAVL())
}

func TestIsAVLUnbalancedRightHeavy(t *testing.T) {
	// right-right heavy: 1 -> right -> 2 -> right -> 3
	n3 := NewBinaryNode(3, nil, nil)
	n2 := NewBinaryNode(2, nil, n3)
	root := NewBinaryNode(1, nil, n2)
	assert.False(t, root.IsAVL())
}
