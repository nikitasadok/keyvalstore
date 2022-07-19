package memtable

import (
	"strings"
)

type Tree struct {
	left  *Tree
	right *Tree
	key   string
	val   []byte
}


func (t *Tree) Insert(key string, val []byte) {
	*t = *insertRec(t, key, val)
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
		t.left = insertRec(t.left, key, val)
	} else if strings.Compare(key, t.key) == 1 {
		t.right = insertRec(t.right, key, val)
	} else {
		t.val = val
		return t
	}

	balance := t.balance()

	if balance < -1 && strings.Compare(key, t.left.key) == -1 {
		// right rotate
		t = rightRotate(t)
	}

	if balance < -1 && strings.Compare(key, t.left.key) == 1 {
		// left, right rotate
		t = leftRightRotate(t)
	}

	if balance > 1 && strings.Compare(key, t.right.key) == 1 {
		// left rotate
		t = leftRotate(t)
	}

	if balance > 1 && strings.Compare(key, t.right.key) == -1 {
		//right, left rot
		t = rightLeftRotate(t)
	}

	return t
}

// insertWithoutRebuild - for testing purposes
func insertWithoutRebuild(t *Tree, key string, val []byte) *Tree {
	if t == nil {
		t = &Tree{
			key: key,
			val: val,
		}
		return t
	}

	if strings.Compare(key, t.key) == -1 {
		t.left = insertWithoutRebuild(t.left, key, val)
	} else {
		t.right = insertWithoutRebuild(t.right, key, val)
	}

	return t
}

func (t *Tree) Search(key string) []byte {
	node := search(t, key)
	if node == nil {
		return nil
	}

	return node.val
}

func search(t *Tree, key string) *Tree {
	if t == nil || t.key == key {
		return t
	}

	if strings.Compare(key, t.key) == -1 {
		return search(t.left, key)
	}

	return search(t.right, key)
}

// only for testing
func preorder(t *Tree, path *[]*Tree) {
	if t == nil {
		return
	}
	preorder(t.left, path)
	*path = append(*path, t)
	preorder(t.right, path)
}

func rightRotate(t *Tree) *Tree {
	var pivot, pivotChild *Tree
	pivot = t
	pivotChild = pivot.left

	pivot.left = pivotChild.right
	pivotChild.right = pivot

	return pivotChild
}

func leftRotate(t *Tree) *Tree {
	var pivot, pivotChild *Tree
	pivot = t
	pivotChild = pivot.right

	pivot.right = pivotChild.left
	pivotChild.left = pivot

	return pivotChild
}

func leftRightRotate(t *Tree) *Tree {
	var pivot, pivotLeft, pivotLeftRight *Tree
	pivot = t
	pivotLeft = t.left
	pivotLeftRight = t.left.right

	pivot.left = pivotLeftRight.right
	pivotLeft.right = pivotLeftRight.left
	pivotLeftRight.right = pivot
	pivotLeftRight.left = pivotLeft

	return pivotLeftRight
}

func rightLeftRotate(t *Tree) *Tree {
	var pivot, pivotRight, pivotRightLeft *Tree
	pivot = t
	pivotRight = t.right
	pivotRightLeft = t.right.left

	pivot.right = pivotRightLeft.left
	pivotRight.left = pivotRightLeft.right
	pivotRightLeft.left = pivot
	pivotRightLeft.right = pivotRight

	return pivotRightLeft
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
