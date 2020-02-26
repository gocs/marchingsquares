package marchingsquares

// Square ...
type Square struct {
	topLeft, topRight, bottomRight, bottomLeft *ControlNode
	topMid, midRight, bottomMid, midLeft       *Node
	config                                     int
}

// NewSquare ...
func NewSquare(topLeft, topRight, bottomRight, bottomLeft *ControlNode) *Square {
	config := 0
	if topLeft.active {
		config++
	}
	if topRight.active {
		config += 2
	}
	if bottomRight.active {
		config += 4
	}
	if bottomLeft.active {
		config += 8
	}

	return &Square{
		topLeft:     topLeft,
		topMid:      topLeft.right,
		topRight:    topRight,
		midLeft:     bottomLeft.above,
		midRight:    bottomRight.above,
		bottomLeft:  bottomLeft,
		bottomMid:   bottomLeft.right,
		bottomRight: bottomRight,
		config:      config,
	}
}

// SquareGrid ...
type SquareGrid struct {
	squares    [][]*Square
	squareSize float32
}

// NewSquareGrid ...
func NewSquareGrid(atlas [][]int, squareSize float32) *SquareGrid {
	nodeCountX := len(atlas)
	nodeCountY := len(atlas[0])
	mapWidth := float32(nodeCountX) * squareSize
	mapHeight := float32(nodeCountY) * squareSize

	controlNodes := make([][]*ControlNode, nodeCountX)
	for x := range controlNodes {
		controlNodes[x] = make([]*ControlNode, nodeCountY)
		for y := range controlNodes[x] {
			pos := Vector2{
				-mapWidth/2 + float32(x)*squareSize + squareSize/2,
				-mapHeight/2 + float32(y)*squareSize + squareSize/2,
			}
			controlNodes[x][y] = NewControlNode(pos, atlas[x][y] == 1, squareSize)
		}
	}
	squares := make([][]*Square, nodeCountX-1)
	for x := range squares {
		squares[x] = make([]*Square, nodeCountY-1)
		for y := range squares[x] {
			squares[x][y] = NewSquare(controlNodes[x][y], controlNodes[x+1][y], controlNodes[x+1][y+1], controlNodes[x][y+1])
		}
	}

	return &SquareGrid{squares: squares, squareSize: squareSize}
}
