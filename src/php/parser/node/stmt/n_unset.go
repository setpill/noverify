package stmt

import (
	"github.com/setpill/noverify/src/php/parser/freefloating"
	"github.com/setpill/noverify/src/php/parser/node"
	"github.com/setpill/noverify/src/php/parser/position"
	"github.com/setpill/noverify/src/php/parser/walker"
)

// Unset node
type Unset struct {
	FreeFloating freefloating.Collection
	Position     *position.Position
	Vars         []node.Node
}

// NewUnset node constructor
func NewUnset(Vars []node.Node) *Unset {
	return &Unset{
		FreeFloating: nil,
		Vars:         Vars,
	}
}

// SetPosition sets node position
func (n *Unset) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *Unset) GetPosition() *position.Position {
	return n.Position
}

func (n *Unset) GetFreeFloating() *freefloating.Collection {
	return &n.FreeFloating
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *Unset) Walk(v walker.Visitor) {
	if !v.EnterNode(n) {
		return
	}

	if n.Vars != nil {
		for _, nn := range n.Vars {
			if nn != nil {
				nn.Walk(v)
			}
		}
	}

	v.LeaveNode(n)
}
