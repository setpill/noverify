package binary

import (
	"github.com/setpill/noverify/src/php/parser/freefloating"
	"github.com/setpill/noverify/src/php/parser/node"
	"github.com/setpill/noverify/src/php/parser/position"
	"github.com/setpill/noverify/src/php/parser/walker"
)

// Coalesce node
type Coalesce struct {
	FreeFloating freefloating.Collection
	Position     *position.Position
	Left         node.Node
	Right        node.Node
}

// NewCoalesce node constructor
func NewCoalesce(Variable node.Node, Expression node.Node) *Coalesce {
	return &Coalesce{
		FreeFloating: nil,
		Left:         Variable,
		Right:        Expression,
	}
}

// SetPosition sets node position
func (n *Coalesce) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Coalesce) GetPosition() *position.Position {
	return n.Position
}

func (n *Coalesce) GetFreeFloating() *freefloating.Collection {
	return &n.FreeFloating
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Coalesce) Walk(v walker.Visitor) {
	if !v.EnterNode(n) {
		return
	}

	if n.Left != nil {
		n.Left.Walk(v)
	}

	if n.Right != nil {
		n.Right.Walk(v)
	}

	v.LeaveNode(n)
}
