package avltree

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_height(t *testing.T) {
	type args struct {
		t *Tree
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "one node",
			args: args{t: &Tree{
				parent: nil,
				left:   nil,
				right:  nil,
				key:    "somekey",
				val:    []byte("someval"),
			}},
			want: 0,
		},
		{
			name: "height 2 tree",
			args: args{t: constructTreeHeight2Balance1()},
			want: 2,
		},
		// TODO: Add test cases.
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, height(tt.args.t))
		})
	}
}

func TestTree_balance(t1 *testing.T) {
	type fields struct {
		parent *Tree
		left   *Tree
		right  *Tree
		key    string
		val    []byte
	}
	tests := []struct {
		name string
		tree *Tree
		want int
	}{
		{
			name: "nil tree",
			tree: nil,
			want: 0},
		{
			name: "one node",
			tree: &Tree{
				parent: nil,
				left:   nil,
				right:  nil,
				key:    "some_key",
				val:    []byte("some_val"),
			},
			want: 0,
		},
		{
			name: "right_heavy",
			tree: constructTreeHeight2Balance1(),
			want: 1,
		},
		{
			name: "balanced",
			tree: constructTreeHeight2Balance0(),
			want: 0,
		},
		{
			name: "left_heavy",
			tree: constructTreeHeight2BalanceNeg1(),
			want: -1,
		},
	}
	for _, tt := range tests {
		t1.Run(tt.name, func(t1 *testing.T) {
			assert.Equalf(t1, tt.want, tt.tree.balance(), "balance()")
		})
	}
}

func constructTreeHeight2Balance1() *Tree {
	root := &Tree{
		parent: nil,
		left:   nil,
		right:  nil,
		key:    "root_key",
		val:    []byte("root_val"),
	}

	rootLeft := &Tree{
		parent: root,
		left:   nil,
		right:  nil,
		key:    "rootLeft_key",
		val:    []byte("rootLeft_val"),
	}

	rootRight := &Tree{
		parent: root,
		left:   nil,
		right:  nil,
		key:    "rootRight_key",
		val:    []byte("rootRight_val"),
	}
	root.right = rootRight
	root.left = rootLeft

	rootRightRight := &Tree{
		parent: rootRight,
		left:   nil,
		right:  nil,
		key:    "rootRightRight_key",
		val:    []byte("rootRightRight_val"),
	}

	rootRightLeft := &Tree{
		parent: rootRight,
		left:   nil,
		right:  nil,
		key:    "rootRightLeft_key",
		val:    []byte("rootRightLeft_val"),
	}

	rootRight.right = rootRightRight
	rootRight.left = rootRightLeft

	return root
}

func constructTreeHeight2BalanceNeg1() *Tree {
	root := &Tree{
		parent: nil,
		left:   nil,
		right:  nil,
		key:    "root_key",
		val:    []byte("root_val"),
	}

	rootLeft := &Tree{
		parent: root,
		left:   nil,
		right:  nil,
		key:    "rootLeft_key",
		val:    []byte("rootLeft_val"),
	}

	rootRight := &Tree{
		parent: root,
		left:   nil,
		right:  nil,
		key:    "rootRight_key",
		val:    []byte("rootRight_val"),
	}
	root.right = rootRight
	root.left = rootLeft

	rootLeftRight := &Tree{
		parent: rootLeft,
		left:   nil,
		right:  nil,
		key:    "rootLeftRight_key",
		val:    []byte("rootLeftRight_val"),
	}

	rootLeftLeft := &Tree{
		parent: rootLeft,
		left:   nil,
		right:  nil,
		key:    "rootLeftLeft_key",
		val:    []byte("rootLeftLeft_val"),
	}

	rootLeft.right = rootLeftRight
	rootLeft.left = rootLeftLeft

	return root
}

func constructTreeHeight2Balance0() *Tree {
	root := &Tree{
		parent: nil,
		left:   nil,
		right:  nil,
		key:    "root_key",
		val:    []byte("root_val"),
	}

	rootLeft := &Tree{
		parent: root,
		left:   nil,
		right:  nil,
		key:    "rootLeft_key",
		val:    []byte("rootLeft_val"),
	}

	rootRight := &Tree{
		parent: root,
		left:   nil,
		right:  nil,
		key:    "rootRight_key",
		val:    []byte("rootRight_val"),
	}
	root.right = rootRight
	root.left = rootLeft

	rootLeftRight := &Tree{
		parent: rootLeft,
		left:   nil,
		right:  nil,
		key:    "rootLeftRight_key",
		val:    []byte("rootLeftRight_val"),
	}

	rootLeftLeft := &Tree{
		parent: rootLeft,
		left:   nil,
		right:  nil,
		key:    "rootLeftLeft_key",
		val:    []byte("rootLeftLeft_val"),
	}

	rootLeft.right = rootLeftRight
	rootLeft.left = rootLeftLeft

	rootRightRight := &Tree{
		parent: rootRight,
		left:   nil,
		right:  nil,
		key:    "rootRightRight_key",
		val:    []byte("rootRightRight_val"),
	}

	rootRightLeft := &Tree{
		parent: rootRight,
		left:   nil,
		right:  nil,
		key:    "rootRightLeft_key",
		val:    []byte("rootRightLeft_val"),
	}

	rootRight.right = rootRightRight
	rootRight.left = rootRightLeft

	return root
}
