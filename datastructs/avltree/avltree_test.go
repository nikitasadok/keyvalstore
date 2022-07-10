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
		name     string
		args     args
		wantPath *[]*Tree
	}{
		{
			name: "nil tree",
			args: args{
				t:    nil,
				path: nil,
			},
			wantPath: nil,
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
				path: &[]*Tree{},
			},
			wantPath: &[]*Tree{
				{
					left:  nil,
					right: nil,
					key:   "some_key",
					val:   []byte("some_val"),
				},
			},
		},
		{
			name: "height 2 tree",
			args: args{
				t:    constructTreeHeight2Balance1(),
				path: &[]*Tree{},
			},
			wantPath: treeHeight2Balance1PreorderTraversal(),
		},
		{
			name: "fully balanced height 2 tree",
			args: args{
				t:    constructTreeHeight2Balance0(),
				path: &[]*Tree{},
			},
			wantPath: treeHeight2Balance0PreorderTraversal(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			preorder(tt.args.t, tt.args.path)
			assert.Equal(t, tt.wantPath, tt.args.path)
		})
	}
}

func Test_rightRotate(t *testing.T) {
	type args struct {
		t *Tree
	}
	tests := []struct {
		name string
		args args
		want *Tree
	}{
		{
			name: "left-heavy tree",
			args: args{t: constructTreeHeight3BalanceNeg2()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			upd := rightRotate(tt.args.t)
			pre := &[]*Tree{}
			preorder(upd, pre)
			assert.Equal(t, treeHeight3BalanceNeg2PreorderTraversalAfterRightRotate(), pre)
			// assert.Equalf(t, tt.want, rightRotate(tt.args.t), "rightRotate(%v)", tt.args.t)
		})
	}
}

func Test_leftRotate(t *testing.T) {
	type args struct {
		t *Tree
	}
	tests := []struct {
		name string
		args args
		want *Tree
	}{
		{
			name: "right-heavy tree",
			args: args{t: constructTreeHeight3Balance2()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			upd := leftRotate(tt.args.t)
			pre := &[]*Tree{}
			preorder(upd, pre)
			assert.Equal(t, treeHeight3Balance2PreorderTraversalAfterLeftRotate(), pre)
		})
	}
}

/*
           root
         /      \
   rootLeft    rootRight
              /         \
       rootRightLeft  rootRightRight
*/
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

/*
        root
      /      \
   rootL    rootR
           /     \
        rootRL  rootRR
                /     \
            rootRRL  rootRRR
*/
func constructTreeHeight3Balance2() *Tree {
	root := &Tree{
		key: "root_key",
		val: []byte("root_val"),
	}

	rootL := &Tree{
		key: "k_rootLeft_key",
		val: []byte("rootLeft_val"),
	}

	rootR := &Tree{
		key: "t_rootRight_key",
		val: []byte("rootRight_val"),
	}
	root.right = rootR
	root.left = rootL

	rootRL := &Tree{
		key: "s_rootRightLeft_key",
		val: []byte("rootRightLeft_val"),
	}

	rootRR := &Tree{
		key: "v_rootRightRight_key",
		val: []byte("rootRightRight_val"),
	}

	rootR.left = rootRL
	rootR.right = rootRR

	rootRRL := &Tree{
		key: "u_rootRRL_key",
		val: []byte("u_rootRRL_val"),
	}

	rootRRR := &Tree{
		key: "z_rootRRR_key",
		val: []byte("z_rootRRR_val"),
	}

	rootRR.left = rootRRL
	rootRR.right = rootRRR

	return root
}

/*
                rootR
             /         \
        root           rootRR
       /    \         /      \
   rootL   rootRL   rootRRL  rootRRR
see constructTreeHeight3Balance2 for previous tree configuration and node names
*/
func treeHeight3Balance2PreorderTraversalAfterLeftRotate() *[]*Tree {
	rootR := &Tree{
		key: "t_rootRight_key",
		val: []byte("rootRight_val"),
	}

	root := &Tree{
		key: "root_key",
		val: []byte("root_val"),
	}

	rootRR := &Tree{
		key: "v_rootRightRight_key",
		val: []byte("rootRightRight_val"),
	}
	rootR.right = rootRR
	rootR.left = root

	rootL := &Tree{
		key: "k_rootLeft_key",
		val: []byte("rootLeft_val"),
	}

	rootRL := &Tree{
		key: "s_rootRightLeft_key",
		val: []byte("rootRightLeft_val"),
	}

	root.right = rootRL
	root.left = rootL

	rootRRL := &Tree{
		key: "u_rootRRL_key",
		val: []byte("u_rootRRL_val"),
	}

	rootRRR := &Tree{
		key: "z_rootRRR_key",
		val: []byte("z_rootRRR_val"),
	}

	rootRR.right = rootRRR
	rootRR.left = rootRRL

	return &[]*Tree{rootL, root, rootRL, rootR, rootRRL, rootRR, rootRRR}
}

func treeHeight2Balance1PreorderTraversal() *[]*Tree {
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

	return &[]*Tree{rootLeft, root, rootRightLeft, rootRight, rootRightRight}
}

/*
                 root
               /      \
         rootLeft    rootRight
        /        \
rootLeftLeft    rootLeftRight
*/
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

/*
                root(-2)
             /        \
        rootL(-1)    rootR(0)
       /     \
   rootLL(0) rootLR(0)
  /         \
rootLLL(0)  rootLLR(0)
*/
func constructTreeHeight3BalanceNeg2() *Tree {
	root := &Tree{
		left:  nil,
		right: nil,
		key:   "root_key",
		val:   []byte("root_val"),
	}

	rootL := &Tree{
		left:  nil,
		right: nil,
		key:   "k_rootLeft_key",
		val:   []byte("rootLeft_val"),
	}

	rootR := &Tree{
		left:  nil,
		right: nil,
		key:   "s_rootRight_key",
		val:   []byte("rootRight_val"),
	}
	root.right = rootR
	root.left = rootL

	rootLR := &Tree{
		left:  nil,
		right: nil,
		key:   "l_rootLeftRight_key",
		val:   []byte("rootLeftRight_val"),
	}

	rootLL := &Tree{
		left:  nil,
		right: nil,
		key:   "b_rootLeftLeft_key",
		val:   []byte("rootLeftLeft_val"),
	}

	rootL.right = rootLR
	rootL.left = rootLL

	rootLLL := &Tree{
		left:  nil,
		right: nil,
		key:   "a_rootLLL_key",
		val:   []byte("rootLLL_val"),
	}

	rootLLR := &Tree{
		left:  nil,
		right: nil,
		key:   "c_rootLLR_key",
		val:   []byte("rootLLR_key"),
	}

	rootLL.left = rootLLL
	rootLL.right = rootLLR

	return root
}

/*
                rootL
             /         \
        rootLL           root
       /     \          /     \
   rootLLL   rootLLR   rootLR  rootR
see constructTreeHeight3BalanceNeg2 for previous tree configuration and node names
*/
func treeHeight3BalanceNeg2PreorderTraversalAfterRightRotate() *[]*Tree {
	rootL := &Tree{
		left:  nil,
		right: nil,
		key:   "k_rootLeft_key",
		val:   []byte("rootLeft_val"),
	}

	rootLL := &Tree{
		left:  nil,
		right: nil,
		key:   "b_rootLeftLeft_key",
		val:   []byte("rootLeftLeft_val"),
	}

	root := &Tree{
		left:  nil,
		right: nil,
		key:   "root_key",
		val:   []byte("root_val"),
	}
	rootL.right = root
	rootL.left = rootLL

	rootLLL := &Tree{
		left:  nil,
		right: nil,
		key:   "a_rootLLL_key",
		val:   []byte("rootLLL_val"),
	}

	rootLLR := &Tree{
		left:  nil,
		right: nil,
		key:   "c_rootLLR_key",
		val:   []byte("rootLLR_key"),
	}

	rootLL.right = rootLLR
	rootLL.left = rootLLL

	rootLR := &Tree{
		left:  nil,
		right: nil,
		key:   "l_rootLeftRight_key",
		val:   []byte("rootLeftRight_val"),
	}

	rootR := &Tree{
		left:  nil,
		right: nil,
		key:   "s_rootRight_key",
		val:   []byte("rootRight_val"),
	}

	root.right = rootR
	root.left = rootLR

	return &[]*Tree{rootLLL, rootLL, rootLLR, rootL, rootLR, root, rootR}

}

/*
                root
             /        \
        rootL         rootR
       /     \       /     \
   rootLL   rootLR rootRL  rootRR
*/
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

func treeHeight2Balance0PreorderTraversal() *[]*Tree {
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

	return &[]*Tree{rootLeftLeft, rootLeft, rootLeftRight, root, rootRightLeft, rootRight, rootRightRight}

}
