package stmt

import (
	"github.com/setpill/noverify/src/php/parser/freefloating"
	"github.com/setpill/noverify/src/php/parser/node"
	"github.com/setpill/noverify/src/php/parser/position"
	"github.com/setpill/noverify/src/php/parser/walker"
)

// UseList node
type UseList struct {
	FreeFloating freefloating.Collection
	Position     *position.Position
	UseType      node.Node
	Uses         []node.Node
}

// NewUseList node constructor
func NewUseList(UseType node.Node, Uses []node.Node) *UseList {
	return &UseList{
		FreeFloating: nil,
		UseType:      UseType,
		Uses:         Uses,
	}
}

// SetPosition sets node position
func (n *UseList) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *UseList) GetPosition() *position.Position {
	return n.Position
}

func (n *UseList) GetFreeFloating() *freefloating.Collection {
	return &n.FreeFloating
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *UseList) Walk(v walker.Visitor) {
	if !v.EnterNode(n) {
		return
	}

	if n.UseType != nil {
		n.UseType.Walk(v)
	}

	if n.Uses != nil {
		for _, nn := range n.Uses {
			if nn != nil {
				nn.Walk(v)
			}
		}
	}

	v.LeaveNode(n)
}
