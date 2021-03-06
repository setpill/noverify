package expr_test

import (
	"bytes"
	"testing"

	"gotest.tools/assert"

	"github.com/setpill/noverify/src/php/parser/node"
	"github.com/setpill/noverify/src/php/parser/node/expr"
	"github.com/setpill/noverify/src/php/parser/node/stmt"
	"github.com/setpill/noverify/src/php/parser/php7"
	"github.com/setpill/noverify/src/php/parser/position"
)

func TestUnaryMinus(t *testing.T) {
	src := `<? -$a;`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  4,
			EndPos:    7,
		},
		Stmts: []node.Node{
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  4,
					EndPos:    7,
				},
				Expr: &expr.UnaryMinus{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  4,
						EndPos:    6,
					},
					Expr: &node.SimpleVar{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  5,
							EndPos:    6,
						},
						Name: "a",
					},
				},
			},
		},
	}

	php7parser := php7.NewParser(bytes.NewBufferString(src), "test.php")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestUnaryPlus(t *testing.T) {
	src := `<? +$a;`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  4,
			EndPos:    7,
		},
		Stmts: []node.Node{
			&stmt.Expression{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  4,
					EndPos:    7,
				},
				Expr: &expr.UnaryPlus{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  4,
						EndPos:    6,
					},
					Expr: &node.SimpleVar{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  5,
							EndPos:    6,
						},
						Name: "a",
					},
				},
			},
		},
	}

	php7parser := php7.NewParser(bytes.NewBufferString(src), "test.php")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}
