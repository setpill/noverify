package stmt_test

import (
	"bytes"
	"testing"

	"gotest.tools/assert"

	"github.com/setpill/noverify/src/php/parser/node/expr/binary"
	"github.com/setpill/noverify/src/php/parser/position"

	"github.com/setpill/noverify/src/php/parser/node/expr"
	"github.com/setpill/noverify/src/php/parser/node/expr/assign"

	"github.com/setpill/noverify/src/php/parser/node/scalar"

	"github.com/setpill/noverify/src/php/parser/node"
	"github.com/setpill/noverify/src/php/parser/node/stmt"
	"github.com/setpill/noverify/src/php/parser/php7"
)

func TestFor(t *testing.T) {
	src := `<? for($i = 0; $i < 10; $i++, $i++) {}`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  4,
			EndPos:    38,
		},
		Stmts: []node.Node{
			&stmt.For{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  4,
					EndPos:    38,
				},
				Init: []node.Node{
					&assign.Assign{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  8,
							EndPos:    13,
						},
						Variable: &node.SimpleVar{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  8,
								EndPos:    9,
							},
							Name: "i",
						},
						Expression: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  13,
								EndPos:    13,
							},
							Value: "0",
						},
					},
				},
				Cond: []node.Node{
					&binary.Smaller{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  16,
							EndPos:    22,
						},
						Left: &node.SimpleVar{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  16,
								EndPos:    17,
							},
							Name: "i",
						},
						Right: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  21,
								EndPos:    22,
							},
							Value: "10",
						},
					},
				},
				Loop: []node.Node{
					&expr.PostInc{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  25,
							EndPos:    28,
						},
						Variable: &node.SimpleVar{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  25,
								EndPos:    26,
							},
							Name: "i",
						},
					},
					&expr.PostInc{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  31,
							EndPos:    34,
						},
						Variable: &node.SimpleVar{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  31,
								EndPos:    32,
							},
							Name: "i",
						},
					},
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  37,
						EndPos:    38,
					},
					Stmts: []node.Node{},
				},
			},
		},
	}

	php7parser := php7.NewParser(bytes.NewBufferString(src), "test.php")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}

func TestAltFor(t *testing.T) {
	src := `<? for(; $i < 10; $i++) : endfor;`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  4,
			EndPos:    33,
		},
		Stmts: []node.Node{
			&stmt.For{
				AltSyntax: true,
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  4,
					EndPos:    33,
				},
				Cond: []node.Node{
					&binary.Smaller{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  10,
							EndPos:    16,
						},
						Left: &node.SimpleVar{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  10,
								EndPos:    11,
							},
							Name: "i",
						},
						Right: &scalar.Lnumber{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  15,
								EndPos:    16,
							},
							Value: "10",
						},
					},
				},
				Loop: []node.Node{
					&expr.PostInc{
						Position: &position.Position{
							StartLine: 1,
							EndLine:   1,
							StartPos:  19,
							EndPos:    22,
						},
						Variable: &node.SimpleVar{
							Position: &position.Position{
								StartLine: 1,
								EndLine:   1,
								StartPos:  19,
								EndPos:    20,
							},
							Name: "i",
						},
					},
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: -1,
						EndLine:   -1,
						StartPos:  -1,
						EndPos:    -1,
					},
					Stmts: []node.Node{},
				},
			},
		},
	}

	php7parser := php7.NewParser(bytes.NewBufferString(src), "test.php")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}
