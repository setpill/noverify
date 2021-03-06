package stmt

import (
	"github.com/setpill/noverify/src/php/parser/freefloating"
	"github.com/setpill/noverify/src/php/parser/position"
	"github.com/setpill/noverify/src/php/parser/walker"
)

// InlineHtml node
type InlineHtml struct {
	FreeFloating freefloating.Collection
	Position     *position.Position
	Value        string
}

// NewInlineHtml node constructor
func NewInlineHtml(Value string) *InlineHtml {
	return &InlineHtml{
		FreeFloating: nil,
		Value:        Value,
	}
}

// SetPosition sets node position
func (n *InlineHtml) SetPosition(p *position.Position) {
	n.Position = p
}

// GetPosition returns node positions
func (n *InlineHtml) GetPosition() *position.Position {
	return n.Position
}

func (n *InlineHtml) GetFreeFloating() *freefloating.Collection {
	return &n.FreeFloating
}

// Walk traverses nodes
// Walk is invoked recursively until v.EnterNode returns true
func (n *InlineHtml) Walk(v walker.Visitor) {
	if !v.EnterNode(n) {
		return
	}

	v.LeaveNode(n)
}
