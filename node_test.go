package marchingsquares

import "testing"

func Test_NewNode(t *testing.T) {
	n := NewNode(Vector2{1, 2})
	t.Logf("n: %#+v\n", n)
}

func Test_NewControlNode(t *testing.T) {
	// topleft
	n := NewControlNode(Vector2{0, 0}, true, 10)
	t.Logf("topleft active: %v", n.active)
	t.Logf("above         : %v", n.above.position)
	t.Logf("node          : %v", n.node.position)
	t.Logf("right         : %v\n\n", n.right.position)

	// topright
	n = NewControlNode(Vector2{0, 10}, false, 10)
	t.Logf("topright active: %v", n.active)
	t.Logf("above          : %v", n.above.position)
	t.Logf("node           : %v", n.node.position)
	t.Logf("right          : %v\n\n", n.right.position)

	// bottomleft
	n = NewControlNode(Vector2{10, 0}, false, 10)
	t.Logf("bottomleft active: %v", n.active)
	t.Logf("above            : %v", n.above.position)
	t.Logf("node             : %v", n.node.position)
	t.Logf("right            : %v\n\n", n.right.position)

	// bottomright
	n = NewControlNode(Vector2{10, 10}, false, 10)
	t.Logf("bottomright active: %v", n.active)
	t.Logf("above             : %v", n.above.position)
	t.Logf("node              : %v", n.node.position)
	t.Logf("right             : %v\n\n", n.right.position)
}

func Test_Consumption(t *testing.T) {
	n := NewControlNode(Vector2{2, 1}, false, 10)
	t.Logf("IsConsumed: %#+v\n", n.above.IsConsumed())
	n.above.Consume()
	t.Logf("IsConsumed: %#+v\n\n", n.above.IsConsumed())
}