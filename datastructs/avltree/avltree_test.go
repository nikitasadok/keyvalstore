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
				left:  nil,
				right: nil,
				key:   "somekey",
				val:   []byte("someval"),
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
		left  *Tree
		right *Tree
		key   string
		val   []byte
	}
	tests := []struct {
		name string
		tree *Tree
		want int
	}{
		{
			name: "one node",
			tree: &Tree{
				left:  nil,
				right: nil,
				key:   "some_key",
				val:   []byte("some_val"),
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

func Test_search(t *testing.T) {
	type args struct {
		t   *Tree
		key string
	}
	tests := []struct {
		name string
		args args
		want *Tree
	}{
		{
			name: "nil tree",
			args: args{t: nil},
			want: nil},
		{
			name: "one node key exists",
			args: args{t: &Tree{
				left:  nil,
				right: nil,
				key:   "some_key",
				val:   []byte("some_val"),
			}, key: "some_key"},
			want: &Tree{
				left:  nil,
				right: nil,
				key:   "some_key",
				val:   []byte("some_val"),
			}},
		{
			name: "one node key doesn't exist",
			args: args{t: &Tree{
				left:  nil,
				right: nil,
				key:   "some_key",
				val:   []byte("some_val"),
			}, key: "some_key_which_doesnt exist"},
			want: nil,
		},
		{
			name: "bigger tree key exists",
			args: args{t: constructTreeHeight2Balance1(), key: "u_rootRightRight_key"},
			want: &Tree{
				left:  nil,
				right: nil,
				key:   "u_rootRightRight_key",
				val:   []byte("rootRightRight_val"),
			},
		},
		{
			name: "bigger tree key doesnt exist",
			args: args{t: constructTreeHeight2Balance1(), key: "some_random_key"},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equalf(t, tt.want, search(tt.args.t, tt.args.key), "search(%v, %v)", tt.args.t, tt.args.key)
		})
	}
}

func Test_preorder(t *testing.T) {
	type args struct {
		t    *Tree
		path *[]*Tree
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "nil tree",
			args: args{
				t:    nil,
				path: nil,
			},
		},
		{
			name: "one node",
			args: args{
				t: &Tree{
					left:  nil,
					right: nil,
					key:   "some_key",
					val:   []byte("some_val"),
				},
				path: &[]*Tree{&Tree{
					key: "some_key",
					val: []byte("some_val"),
				}},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			preorder(tt.args.t, tt.args.path)
		})
	}
}

func constructTreeHeight2Balance1() *Tree {
	root := &Tree{
		left:  nil,
		right: nil,
		key:   "root_key",
		val:   []byte("root_val"),
	}

	rootLeft := &Tree{
		left:  nil,
		right: nil,
		key:   "k_rootLeft_key",
		val:   []byte("rootLeft_val"),
	}

	rootRight := &Tree{
		left:  nil,
		right: nil,
		key:   "t_rootRight_key",
		val:   []byte("rootRight_val"),
	}
	root.right = rootRight
	root.left = rootLeft

	rootRightRight := &Tree{
		left:  nil,
		right: nil,
		key:   "u_rootRightRight_key",
		val:   []byte("rootRightRight_val"),
	}

	rootRightLeft := &Tree{
		left:  nil,
		right: nil,
		key:   "s_rootRightLeft_key",
		val:   []byte("rootRightLeft_val"),
	}

	rootRight.right = rootRightRight
	rootRight.left = rootRightLeft

	return root
}

func constructTreeHeight2BalanceNeg1() *Tree {
	root := &Tree{
		left:  nil,
		right: nil,
		key:   "root_key",
		val:   []byte("root_val"),
	}

	rootLeft := &Tree{
		left:  nil,
		right: nil,
		key:   "k_rootLeft_key",
		val:   []byte("rootLeft_val"),
	}

	rootRight := &Tree{
		left:  nil,
		right: nil,
		key:   "s_rootRight_key",
		val:   []byte("rootRight_val"),
	}
	root.right = rootRight
	root.left = rootLeft

	rootLeftRight := &Tree{
		left:  nil,
		right: nil,
		key:   "l_rootLeftRight_key",
		val:   []byte("rootLeftRight_val"),
	}

	rootLeftLeft := &Tree{
		left:  nil,
		right: nil,
		key:   "a_rootLeftLeft_key",
		val:   []byte("rootLeftLeft_val"),
	}

	rootLeft.right = rootLeftRight
	rootLeft.left = rootLeftLeft

	return root
}

func constructTreeHeight2Balance0() *Tree {
	root := &Tree{
		left:  nil,
		right: nil,
		key:   "root_key",
		val:   []byte("root_val"),
	}

	rootLeft := &Tree{
		left:  nil,
		right: nil,
		key:   "k_rootLeft_key",
		val:   []byte("rootLeft_val"),
	}

	rootRight := &Tree{
		left:  nil,
		right: nil,
		key:   "v_rootRight_key",
		val:   []byte("rootRight_val"),
	}
	root.right = rootRight
	root.left = rootLeft

	rootLeftRight := &Tree{
		left:  nil,
		right: nil,
		key:   "l_rootLeftRight_key",
		val:   []byte("rootLeftRight_val"),
	}

	rootLeftLeft := &Tree{
		left:  nil,
		right: nil,
		key:   "a_rootLeftLeft_key",
		val:   []byte("rootLeftLeft_val"),
	}

	rootLeft.right = rootLeftRight
	rootLeft.left = rootLeftLeft

	rootRightRight := &Tree{
		left:  nil,
		right: nil,
		key:   "w_rootRightRight_key",
		val:   []byte("rootRightRight_val"),
	}

	rootRightLeft := &Tree{
		left:  nil,
		right: nil,
		key:   "u_rootRightLeft_key",
		val:   []byte("rootRightLeft_val"),
	}

	rootRight.right = rootRightRight
	rootRight.left = rootRightLeft

	return root
}
