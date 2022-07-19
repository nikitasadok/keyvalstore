package memtable

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

type treeEntry struct {
	key string
	val []byte
}

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
			args: args{t: constructTreeHeight2Balance1(), key: "t_rootRightLeft_key"},
			want: &Tree{
				left:  nil,
				right: nil,
				key:   "t_rootRightLeft_key",
				val:   []byte("rootRightLeft_val"),
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

func Test_leftRightRotate(t *testing.T) {
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
			args: args{t: constructTreeHeight3BalanceNeg2LR()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			upd := leftRightRotate(tt.args.t)
			pre := &[]*Tree{}
			preorder(upd, pre)
			assert.Equal(t, treeHeight3BalanceNeg2LRPreorderTraversalAfterLRRotate(), pre)
		})
	}
}

func Test_rightLeftRotate(t *testing.T) {
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
			args: args{t: constructTreeHeight3Balance2RL()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			upd := rightLeftRotate(tt.args.t)
			pre := &[]*Tree{}
			preorder(upd, pre)
			assert.Equal(t, treeHeight3Balance2LRPreorderTraversalAfterRLRotate(), pre)
		})
	}
}


func Test_insertRec(t *testing.T) {
	type args struct {
		t   *Tree
		key string
		val []byte
	}
	/*
	{
				key:   "c_rootLRL_key",
				val:   []byte("c_rootLRL_val"),
			},
	 */
	tests := []struct {
		name string
		args args
		want *Tree
	}{
		{
			name: "right-heavy after insert to right subtree",
			args: args{
				t:   constructTreeHeight2Balance1(),
				key: "z_rootRRR_key",
				val: []byte("z_rootRRR_val"),
			},
			want: treeHeight2Balance1AfterRRRInsert(),
		},
		{
			name: "right-heavy after insert to left subtree",
			args: args{
				t:   constructTreeHeight2Balance1(),
				key: "s_rootRLL_key",
				val: []byte("s_rootRLL_val"),
			},
			want: treeHeight2Balance1AfterRLLInsert(),
		},
		{
			name: "left-heavy after insert to left subtree",
			args: args{
				t:   constructTreeHeight2BalanceNeg1(),
				key: "a_rootLLL_key",
				val: []byte("rootLLL_val"),
			},
			want: treeHeight2BalanceNeg1AfterInsertLLL(),
		},
		{
			name: "left-heavy after insert to right subtree",
			args: args{
				t:   constructTreeHeight2BalanceNeg1(),
				key: "m_rootLRL_key",
				val: []byte("m_rootLRL_val"),
			},
			want: treeHeight2BalanceNeg1AfterInsertLRL(),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insertRec(tt.args.t, tt.args.key, tt.args.val); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insertRec() = %v, want %v", got, tt.want)
			}
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
	var root *Tree
	root = fillTree(root, []treeEntry{
		{
			key: "root_key",
			val: []byte("root_val"),
		},
		{
			key: "k_rootLeft_key",
			val: []byte("rootLeft_val"),
		},
		{
			key:   "v_rootRight_key",
			val:   []byte("rootRight_val"),
		},
		{
			key:   "t_rootRightLeft_key",
			val:   []byte("rootRightLeft_val"),
		},
		{
			key:   "y_rootRightRight_key",
			val:   []byte("rootRightRight_val"),
		},
	})

	return root
}

/*
                rootR
             /         \
        root           rootRR
       /    \                \
   rootL   rootRL          rootRRR
 */
func treeHeight2Balance1AfterRRRInsert() *Tree {
	var root *Tree
	root = fillTree(root, []treeEntry{
		{
			key:   "v_rootRight_key",
			val:   []byte("rootRight_val"),
		},
		{
			key: "root_key",
			val: []byte("root_val"),
		},
		{
			key:   "y_rootRightRight_key",
			val:   []byte("rootRightRight_val"),
		},
		{
			key: "k_rootLeft_key",
			val: []byte("rootLeft_val"),
		},
		{
			key:   "t_rootRightLeft_key",
			val:   []byte("rootRightLeft_val"),
		},
		{
			key:   "z_rootRRR_key",
			val:   []byte("z_rootRRR_val"),
		},
	})

	return root
}

/*
                 rootRL
             /        \
        root         rootR
       /     \            \
   rootL   rootRLL      rootRR
*/
func treeHeight2Balance1AfterRLLInsert() *Tree {
	var root *Tree
	root = fillTree(root, []treeEntry{
		{
			key:   "t_rootRightLeft_key",
			val:   []byte("rootRightLeft_val"),
		},
		{
			key: "root_key",
			val: []byte("root_val"),
		},
		{
			key:   "v_rootRight_key",
			val:   []byte("rootRight_val"),
		},
		{
			key: "k_rootLeft_key",
			val: []byte("rootLeft_val"),
		},
		{
			key:   "y_rootRightRight_key",
			val:   []byte("rootRightRight_val"),
		},
		{
			key:   "s_rootRLL_key",
			val:   []byte("s_rootRLL_val"),
		},
	})

	return root
}

func treeHeight2Balance1PreorderTraversal() *[]*Tree {
	root := &Tree{
		key:   "root_key",
		val:   []byte("root_val"),
	}

	rootLeft := &Tree{
		key:   "k_rootLeft_key",
		val:   []byte("rootLeft_val"),
	}

	rootRight := &Tree{
		key:   "v_rootRight_key",
		val:   []byte("rootRight_val"),
	}
	root.left = rootLeft
	root.right = rootRight

	rootRightLeft := &Tree{
		key:   "t_rootRightLeft_key",
		val:   []byte("rootRightLeft_val"),
	}

	rootRightRight := &Tree{
		key:   "y_rootRightRight_key",
		val:   []byte("rootRightRight_val"),
	}

	rootRight.left = rootRightLeft
	rootRight.right = rootRightRight

	return &[]*Tree{rootLeft, root, rootRightLeft, rootRight, rootRightRight}
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
	var root *Tree
	root = fillTree(root, []treeEntry{
		{key: "root_key",
			val: []byte("root_val"),
		},
		{
			key: "k_rootLeft_key",
			val: []byte("k_rootLeft_val"),
		},
		{
			key:   "v_rootRight_key",
			val:   []byte("rootRight_val"),
		},
		{
			key:   "u_rootRightLeft_key",
			val:   []byte("rootRightLeft_val"),
		},
		{
			key:   "w_rootRightRight_key",
			val:   []byte("rootRightRight_val"),
		},
		{
			key:   "v_rootRRL_key",
			val:   []byte("v_rootRRL_val"),
		},
		{
			key:   "z_rootRRR_key",
			val:   []byte("z_rootRRR_val"),
		},
	})

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
	var (
		root *Tree
		path = &[]*Tree{}
	)
	root = fillTree(root, []treeEntry{
		{
			key:   "v_rootRight_key",
			val:   []byte("rootRight_val"),
		},
		{
			key: "root_key",
			val: []byte("root_val"),
		},
		{
			key:   "w_rootRightRight_key",
			val:   []byte("rootRightRight_val"),
		},
		{
			key: "k_rootLeft_key",
			val: []byte("k_rootLeft_val"),
		},
		{
			key:   "u_rootRightLeft_key",
			val:   []byte("rootRightLeft_val"),
		},
		{
			key:   "v_rootRRL_key",
			val:   []byte("v_rootRRL_val"),
		},
		{
			key:   "z_rootRRR_key",
			val:   []byte("z_rootRRR_val"),
		},
	})

	preorder(root, path)
	return path
}


/*
                 root
               /      \
         rootLeft    rootRight
        /        \
rootLeftLeft    rootLeftRight
*/
func constructTreeHeight2BalanceNeg1() *Tree {
	var root *Tree
	root = fillTree(root, []treeEntry{
		{
			key: "root_key",
			val: []byte("root_val"),
		},
		{
			key: "k_rootLeft_key",
			val: []byte("rootLeft_val"),
		},
		{
			key:   "v_rootRight_key",
			val:   []byte("rootRight_val"),
		},
		{
			key:   "b_rootLeftLeft_key",
			val:   []byte("rootLeftLeft_val"),
		},
		{
			key:   "l_rootLeftRight_key",
			val:   []byte("rootLeftRight_val"),
		},
	})

	return root
}

/*
                rootL
             /         \
        rootLL           root
       /                /     \
   rootLLL           rootLR  rootR
 */
func treeHeight2BalanceNeg1AfterInsertLLL() *Tree {
	var root *Tree
	root = fillTree(root, []treeEntry{
		{
			key:   "k_rootLeft_key",
			val:   []byte("rootLeft_val"),
		},
		{
			key:   "b_rootLeftLeft_key",
			val:   []byte("rootLeftLeft_val"),
		},
		{
			key:   "root_key",
			val:   []byte("root_val"),
		},
		{
			key:   "a_rootLLL_key",
			val:   []byte("rootLLL_val"),
		},
		{
			key:   "l_rootLeftRight_key",
			val:   []byte("rootLeftRight_val"),
		},
		{
			key:   "v_rootRight_key",
			val:   []byte("rootRight_val"),
		},
	})

	return root
}
/*
               rootLR
            /         \
       rootL           root
      /     \               \
  rootLL   rootLRL          rootR
 */
func treeHeight2BalanceNeg1AfterInsertLRL() *Tree {
	var root *Tree

	root = fillTree(root, []treeEntry{
		{
			key:   "l_rootLeftRight_key",
			val:   []byte("rootLeftRight_val"),
		},
		{
			key: "root_key",
			val: []byte("root_val"),
		},
		{
			key: "k_rootLeft_key",
			val: []byte("rootLeft_val"),
		},
		{
			key:   "v_rootRight_key",
			val:   []byte("rootRight_val"),
		},
		{
			key:   "b_rootLeftLeft_key",
			val:   []byte("rootLeftLeft_val"),
		},
		{
			key:   "m_rootLRL_key",
			val:   []byte("m_rootLRL_val"),
		},
	})

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
	var root *Tree
	path := &[]*Tree{}

	root = fillTree(root, []treeEntry{
		{
			key:   "k_rootLeft_key",
			val:   []byte("rootLeft_val"),
		},
		{
			key:   "b_rootLeftLeft_key",
			val:   []byte("rootLeftLeft_val"),
		},
		{
			key:   "root_key",
			val:   []byte("root_val"),
		},
		{
			key:   "a_rootLLL_key",
			val:   []byte("rootLLL_val"),
		},
		{
			key:   "c_rootLLR_key",
			val:   []byte("rootLLR_key"),
		},
		{
			key:   "l_rootLeftRight_key",
			val:   []byte("rootLeftRight_val"),
		},
		{
			key:   "s_rootRight_key",
			val:   []byte("rootRight_val"),
		},
	})

	preorder(root, path)
	return path
}

/*
                root(-2)
             /        \
        rootL(1)    rootR(0)
       /     \
   rootLL(0) rootLR(0)
            /        \
       rootLRL(0)  rootLRR(0)
*/
func constructTreeHeight3BalanceNeg2LR() *Tree {
	var root *Tree
	root = fillTree(root, []treeEntry{
		{
			key: "root_key",
			val: []byte("root_val"),
		},
		{
			key: "k_rootLeft_key",
			val: []byte("k_rootLeft_val"),
		},
		{
			key:   "v_rootRight_key",
			val:   []byte("rootRight_val"),
		},
		{
			key:   "a_rootLeftLeft_key",
			val:   []byte("rootLeftLeft_val"),
		},
		{
			key:   "l_rootLeftRight_key",
			val:   []byte("rootLeftRight_val"),
		},
		{
			key:   "c_rootLRL_key",
			val:   []byte("c_rootLRL_val"),
		},
		{
			key:   "m_rootLRR_key",
			val:   []byte("m_rootLRR_val"),
		},
	})


	return root
}

/*
               rootLR
            /         \
       rootL           root
      /     \          /     \
  rootLL   rootLRL   rootLRR  rootR
see constructTreeHeight3BalanceNeg2 for previous tree configuration and node names
*/
func treeHeight3BalanceNeg2LRPreorderTraversalAfterLRRotate() *[]*Tree {
	var (
		root *Tree
		path *[]*Tree
	)

	path = &[]*Tree{}
	root = fillTree(root, []treeEntry{
		{
			key:   "l_rootLeftRight_key",
			val:   []byte("rootLeftRight_val"),
		},
		{
			key: "root_key",
			val: []byte("root_val"),
		},
		{
			key: "k_rootLeft_key",
			val: []byte("k_rootLeft_val"),
		},
		{
			key:   "v_rootRight_key",
			val:   []byte("rootRight_val"),
		},
		{
			key:   "a_rootLeftLeft_key",
			val:   []byte("rootLeftLeft_val"),
		},
		{
			key:   "c_rootLRL_key",
			val:   []byte("c_rootLRL_val"),
		},
		{
			key:   "m_rootLRR_key",
			val:   []byte("m_rootLRR_val"),
		},
	})

	preorder(root, path)
	return path
}

/*
                root(2)
               /      \
        rootL(0)    rootR(-1)
                    /     \
              rootRL(0) rootRR(0)
            /        \
       rootRLL(0)  rootRLR(0)
*/
func constructTreeHeight3Balance2RL() *Tree {
	var root *Tree
	root = fillTree(root, []treeEntry{
		{key: "root_key",
			val: []byte("root_val"),
		},
		{
			key: "k_rootLeft_key",
			val: []byte("k_rootLeft_val"),
		},
		{
			key:   "v_rootRight_key",
			val:   []byte("rootRight_val"),
		},
		{
			key:   "t_rootRightLeft_key",
			val:   []byte("rootRightLeft_val"),
		},
		{
			key:   "z_rootRightRight_key",
			val:   []byte("rootLeftRight_val"),
		},
		{
			key:   "s_rootRLL_key",
			val:   []byte("s_rootRLL_val"),
		},
		{
			key:   "u_rootRLR_key",
			val:   []byte("u_rootRLR_val"),
		},
	})


	return root
}

/*
                rootRL
             /        \
        root         rootR
       /     \       /     \
   rootL   rootRLL rootRLR  rootRR
*/
func treeHeight3Balance2LRPreorderTraversalAfterRLRotate() *[]*Tree {
	var (
		root *Tree
		path *[]*Tree
	)
	path = &[]*Tree{}
	root = fillTree(root, []treeEntry{
		{
			key:   "t_rootRightLeft_key",
			val:   []byte("rootRightLeft_val"),
		},
		{key: "root_key",
			val: []byte("root_val"),
		},
		{
			key:   "v_rootRight_key",
			val:   []byte("rootRight_val"),
		},
		{
			key: "k_rootLeft_key",
			val: []byte("k_rootLeft_val"),
		},
		{
			key:   "z_rootRightRight_key",
			val:   []byte("rootLeftRight_val"),
		},
		{
			key:   "s_rootRLL_key",
			val:   []byte("s_rootRLL_val"),
		},
		{
			key:   "u_rootRLR_key",
			val:   []byte("u_rootRLR_val"),
		},
	})

	preorder(root,path)
	return path
}

/*
                root
             /        \
        rootL         rootR
       /     \       /     \
   rootLL   rootLR rootRL  rootRR
*/
func constructTreeHeight2Balance0() *Tree {
	var root *Tree
	root = fillTree(root, []treeEntry{
		{key: "root_key",
			val: []byte("root_val"),
		},
		{
			key: "k_rootLeft_key",
			val: []byte("rootLeft_val"),
		},
		{
			key:   "v_rootRight_key",
			val:   []byte("rootRight_val"),
		},
		{
			key:   "a_rootLeftLeft_key",
			val:   []byte("rootLeftLeft_val"),
		},
		{
			key:   "l_rootLeftRight_key",
			val:   []byte("rootLeftRight_val"),
		},
		{
			key:   "u_rootRightLeft_key",
			val:   []byte("rootRightLeft_val"),
		},
		{
			key:   "w_rootRightRight_key",
			val:   []byte("rootRightRight_val"),
		},
	})

	return root
}

func treeHeight2Balance0PreorderTraversal() *[]*Tree {
	root := &Tree{
		key:   "root_key",
		val:   []byte("root_val"),
	}

	rootLeft := &Tree{
		key:   "k_rootLeft_key",
		val:   []byte("rootLeft_val"),
	}

	rootRight := &Tree{
		key:   "v_rootRight_key",
		val:   []byte("rootRight_val"),
	}

	root.left = rootLeft
	root.right = rootRight

	rootLeftLeft := &Tree{
		key:   "a_rootLeftLeft_key",
		val:   []byte("rootLeftLeft_val"),
	}

	rootLeftRight := &Tree{
		key:   "l_rootLeftRight_key",
		val:   []byte("rootLeftRight_val"),
	}

	rootLeft.left = rootLeftLeft
	rootLeft.right = rootLeftRight

	rootRightLeft := &Tree{
		key:   "u_rootRightLeft_key",
		val:   []byte("rootRightLeft_val"),
	}

	rootRightRight := &Tree{
		key:   "w_rootRightRight_key",
		val:   []byte("rootRightRight_val"),
	}

	rootRight.left = rootRightLeft
	rootRight.right = rootRightRight

	return &[]*Tree{rootLeftLeft, rootLeft, rootLeftRight, root, rootRightLeft, rootRight, rootRightRight}

}

func fillTree(t *Tree, entries []treeEntry) *Tree {
	for _, e := range entries {
		t = insertWithoutRebuild(t, e.key, e.val)
	}

	return t
}
