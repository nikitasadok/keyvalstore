package avltree

type Node struct {
	parent     *Node
	leftChild  *Node
	rightChild *Node
	key        string
	val        []byte
}
