package avltree

import "strings"

type Tree struct {
	parent *Tree
	left   *Tree
	right  *Tree
	key    string
	val    []byte
}

func (t *Tree) Insert(key string, val []byte) {
	if key == t.key {
		t.val = val
		return
	}

}

func insert(t *Tree, key string, val []byte) {
	t = insertRec(t, key, val)

}
func insertRec(t *Tree, key string, val []byte) *Tree {
	if t == nil {
		t = &Tree{
			key: key,
			val: val,
		}
		return t
	}

	if strings.Compare(key, t.key) == -1 {
		lChild := insertRec(t.left, key, val)
		t.left = lChild
		t.left.parent = t
	} else if strings.Compare(key, t.key) == 1 {
		rChild := insertRec(t.right, key, val)
		t.right = rChild
		t.right.parent = t
	} else {
		t.val = val
		return t
	}

	balance := t.balance()

	if balance < -1 && strings.Compare(key, t.left.key) == -1 {
		// right rotate
	}

	if balance < -1 && strings.Compare(key, t.left.key) == 1 {
		// left, right rotate
	}

	if balance > 1 && strings.Compare(key, t.right.key) == 1 {
		// left rotate
	}

	if balance > 1 && strings.Compare(key, t.right.key) == -1 {
		//right, left rot
	}

	return t
}

func rightRotate(t *Tree) *Tree {
	var pivot, pivotChild *Tree
	pivotChild = pivot.left
	pivot = t

	pivot.left = pivotChild.right
	pivotChild.right = pivot

	return pivotChild
}

func leftRotate(t *Tree) *Tree {
	var pivot, pivotChild *Tree
	pivotChild = pivot.right
	pivot = t

	pivot.right = pivotChild.left
	pivotChild.left = pivotChild

	return pivotChild
}

func (t *Tree) balance() int {
	return height(t.right) - height(t.left)
}

func height(t *Tree) int {
	if t == nil {
		return -1
	}
	return 1 + max(height(t.left), height(t.right))
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
