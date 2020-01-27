package node

import (
	"github.com/setpill/noverify/src/php/parser/freefloating"
	"github.com/setpill/noverify/src/php/parser/position"
	"github.com/setpill/noverify/src/php/parser/walker"
)

// Node interface
type Node interface {
	walker.Walkable
	SetPosition(p *position.Position)
	GetPosition() *position.Position
	GetFreeFloating() *freefloating.Collection
}
