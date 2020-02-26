package marchingsquares

// Node gives position based on index
type Node struct {
	position    Vector2
	vertexIndex uint16
	consumed    bool // because uint don't have -1
}

// NewNode initializes node with default index, given position
func NewNode(pos Vector2) *Node {
	return &Node{
		position:    pos,
		vertexIndex: 0,
	}
}

// Consume will be used if vertexindex is set
func (n *Node) Consume() {
	n.consumed = true
}

// IsConsumed gives true if vertexIndex is already consumed or set
func (n *Node) IsConsumed() bool {
	return n.consumed
}

// ControlNode contains current, above, and right nodes, and active
type ControlNode struct {
	node, above, right *Node
	active             bool
}

// NewControlNode set an above node and right node depending on the given size when active
func NewControlNode(pos Vector2, active bool, squareSize float32) *ControlNode {
	ss := squareSize / 2
	return &ControlNode{
		node:   NewNode(pos),
		active: active,
		// pos + Vector2.forward * ss
		above:  NewNode(Vector2{x: pos.x, y: pos.y - ss}),
		// pos + Vector2.right * ss
		right:  NewNode(Vector2{x: pos.x + ss, y: pos.y}),
	}
}
