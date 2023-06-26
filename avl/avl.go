package avl

type Node struct {
	Key    string
	Value  interface{}
	Height int
	Left   *Node
	Right  *Node
}

// A utility function to get the height of the node
func height(N *Node) int {
	if N == nil {
		return 0
	}
	return N.Height
}

// A utility function to get the maximum of two integers
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Function to create a new node
func newNode(key string, value interface{}) *Node {
	node := &Node{Key: key, Value: value, Height: 1}
	return node
}

// Function to right rotate subtree rooted with y
func rightRotate(y *Node) *Node {
	x := y.Left
	T2 := x.Right

	// Perform rotation
	x.Right = y
	y.Left = T2

	// Update heights
	y.Height = max(height(y.Left), height(y.Right)) + 1
	x.Height = max(height(x.Left), height(x.Right)) + 1

	// Return new root
	return x
}

// Function to left rotate subtree rooted with x
func leftRotate(x *Node) *Node {
	y := x.Right
	T2 := y.Left

	// Perform rotation
	y.Left = x
	x.Right = T2

	// Update heights
	x.Height = max(height(x.Left), height(x.Right)) + 1
	y.Height = max(height(y.Left), height(y.Right)) + 1

	// Return new root
	return y
}

// Get balance factor of node N
func getBalance(N *Node) int {
	if N == nil {
		return 0
	}
	return height(N.Left) - height(N.Right)
}

// Insert a node
func insert(node *Node, key string, value interface{}) *Node {
	// 1. Perform the normal BST rotation
	if node == nil {
		return newNode(key, value)
	}

	if key < node.Key {
		node.Left = insert(node.Left, key, value)
	} else if key > node.Key {
		node.Right = insert(node.Right, key, value)
	} else {
		// Equal keys are not allowed in BST, updating value instead
		node.Value = value
		return node
	}

	// 2. Update height of this ancestor node
	node.Height = 1 + max(height(node.Left), height(node.Right))

	// 3. Get the balance factor
	balance := getBalance(node)

	// If node is unbalanced, then there are 4 cases

	// Left Left Case
	if balance > 1 && key < node.Left.Key {
		return rightRotate(node)
	}

	// Right Right Case
	if balance < -1 && key > node.Right.Key {
		return leftRotate(node)
	}

	// Left Right Case
	if balance > 1 && key > node.Left.Key {
		node.Left = leftRotate(node.Left)
		return rightRotate(node)
	}

	// Right Left Case
	if balance < -1 && key < node.Right.Key {
		node.Right = rightRotate(node.Right)
		return leftRotate(node)
	}

	// return the (unchanged) node pointer
	return node
}

// Function to search a node with given key
func search(root *Node, key string) *Node {
	// Base Cases: root is null or key is present at root
	if root == nil || root.Key == key {
		return root
	}

	// Key is greater than root's key
	if root.Key < key {
		return search(root.Right, key)
	}

	// Key is smaller than root's key
	return search(root.Left, key)
}
