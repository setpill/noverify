package stmt_test

import (
	"bytes"
	"testing"

	"gotest.tools/assert"

	"github.com/setpill/noverify/src/php/parser/node/scalar"
	"github.com/setpill/noverify/src/php/parser/position"

	"github.com/setpill/noverify/src/php/parser/node"
	"github.com/setpill/noverify/src/php/parser/node/stmt"
	"github.com/setpill/noverify/src/php/parser/php7"
)

func TestDo(t *testing.T) {
	src := `<? do {} while(1);`

	expected := &node.Root{
		Position: &position.Position{
			StartLine: 1,
			EndLine:   1,
			StartPos:  4,
			EndPos:    18,
		},
		Stmts: []node.Node{
			&stmt.Do{
				Position: &position.Position{
					StartLine: 1,
					EndLine:   1,
					StartPos:  4,
					EndPos:    18,
				},
				Stmt: &stmt.StmtList{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  7,
						EndPos:    8,
					},
					Stmts: []node.Node{},
				},
				Cond: &scalar.Lnumber{
					Position: &position.Position{
						StartLine: 1,
						EndLine:   1,
						StartPos:  16,
						EndPos:    16,
					},
					Value: "1",
				},
			},
		},
	}

	php7parser := php7.NewParser(bytes.NewBufferString(src), "test.php")
	php7parser.Parse()
	actual := php7parser.GetRootNode()
	assert.DeepEqual(t, expected, actual)
}
