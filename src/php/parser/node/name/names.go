package name

import (
	"github.com/setpill/noverify/src/php/parser/node"
)

// Names is generalizing the Name types
type Names interface {
	node.Node
	GetParts() []node.Node
}
