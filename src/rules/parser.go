package rules

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/VKCOM/noverify/src/linter/lintapi"
	"github.com/VKCOM/noverify/src/php/parser/freefloating"
	"github.com/VKCOM/noverify/src/php/parser/node"
	"github.com/VKCOM/noverify/src/php/parser/node/stmt"
	"github.com/VKCOM/noverify/src/php/parser/php7"
	"github.com/VKCOM/noverify/src/phpdoc"
	"github.com/VKCOM/noverify/src/phpgrep"
)

var magicComment = regexp.MustCompile(`\* @(?:warning|error|info|maybe) `)

type parseError struct {
	filename string
	lineNum  int
	msg      string
}

func (e *parseError) Error() string {
	return fmt.Sprintf("%s:%d: %s", e.filename, e.lineNum, e.msg)
}

// parser parses rules file into a RuleSet.
type parser struct {
	filename   string
	sources    []byte
	res        *Set
	parentRule Rule // Used when phpdoc is shared between patterns
	compiler   phpgrep.Compiler
}

// Parse reads PHP code that represents a rule file from r and creates a RuleSet based on it.
func (p *parser) parse(filename string, r io.Reader) (*Set, error) {
	res := NewSet()

	// Parse PHP file.
	//
	// TODO: make phpgrep.compile accepting AST and stop
	// slurping sources here + don't parse it twice.
	sources, err := ioutil.ReadAll(r)
	if err != nil {
		return res, err
	}
	q := php7.NewParser(bytes.NewReader(sources), filename)
	q.WithFreeFloating()
	q.Parse()
	if errs := q.GetErrors(); len(errs) != 0 {
		return res, errors.New(errs[0].String())
	}
	root := q.GetRootNode()

	// Convert PHP file into the rule set.
	p.filename = filename
	p.sources = sources
	p.res = res
	for _, st := range root.Stmts {
		if err := p.parseRule(st); err != nil {
			return p.res, err
		}
	}

	return p.res, nil
}

func (p *parser) readRule(dst *Rule, st node.Node, source string) error {
	rule := dst

	var filterSet map[string]Filter

	for _, part := range phpdoc.Parse(source) {
		switch part.Name {
		case "name":
			if len(part.Params) != 1 {
				return p.errorf(st, "@name expects exactly 1 param, got %d", len(part.Params))
			}
			rule.Name = part.Params[0]

		case "location":
			if len(part.Params) != 1 {
				return p.errorf(st, "@type expects exactly 1 params, got %d", len(part.Params))
			}
			name := part.Params[0]
			if !strings.HasPrefix(name, "$") {
				return p.errorf(st, "@type 2nd param must be a phpgrep variable")
			}
			rule.Location = strings.TrimPrefix(name, "$")

		case "scope":
			if len(part.Params) != 1 {
				return p.errorf(st, "@scope expects exactly 1 params, got %d", len(part.Params))
			}
			switch part.Params[0] {
			case "any", "root", "local":
				rule.scope = part.Params[0]
			default:
				return p.errorf(st, "unknown @scope: %s", part.Params[0])
			}

		case "error":
			rule.Level = lintapi.LevelError
			rule.Message = part.ParamsText
		case "warning":
			rule.Level = lintapi.LevelWarning
			rule.Message = part.ParamsText
		case "info":
			rule.Level = lintapi.LevelInformation
			rule.Message = part.ParamsText
		case "maybe":
			rule.Level = lintapi.LevelMaybe
			rule.Message = part.ParamsText

		case "or":
			rule.Filters = append(rule.Filters, filterSet)
			filterSet = nil
		case "type":
			if len(part.Params) != 2 {
				return p.errorf(st, "@type expects exactly 2 params, got %d", len(part.Params))
			}
			typ := part.Params[0]
			name := part.Params[1]
			if !strings.HasPrefix(name, "$") {
				return p.errorf(st, "@type 2nd param must be a phpgrep variable")
			}
			name = strings.TrimPrefix(name, "$")
			if filterSet == nil {
				filterSet = map[string]Filter{}
			}
			filter := filterSet[name]
			if filter.Types != nil {
				return p.errorf(st, "$%s: duplicate type constraint", name)
			}
			filter.Types = strings.Split(typ, "|")
			filterSet[name] = filter

		default:
			return p.errorf(st, "unknown attribute @%s on line %d", part.Name, part.Line)
		}
	}

	if filterSet != nil {
		rule.Filters = append(rule.Filters, filterSet)
	}

	return nil
}

func (p *parser) parseRule(st node.Node) error {
	comment := ""
	for _, ff := range (*st.GetFreeFloating())[freefloating.Start] {
		if ff.StringType != freefloating.CommentType {
			continue
		}
		if strings.HasPrefix(ff.Value, "/**") && magicComment.MatchString(ff.Value) {
			comment = ff.Value
			break
		}
	}

	var rule Rule
	if comment == "" {
		// Inherit rule props from the parent rule.
		rule = p.parentRule
	} else {
		if err := p.readRule(&rule, st, comment); err != nil {
			return err
		}
		p.parentRule = rule
	}

	dst := p.res.Any // Use "any" set by default
	switch rule.scope {
	case "root":
		dst = p.res.Root
	case "local":
		dst = p.res.Local
	}

	if rule.Name == "" {
		rule.Name = fmt.Sprintf("%s:%d", filepath.Base(p.filename), st.GetPosition().StartLine)
	}
	p.res.ToAllow = append(p.res.ToAllow, rule.Name)

	if rule.Level == lintapi.LevelError || rule.Level == lintapi.LevelWarning {
		p.res.AlwaysCritical = append(p.res.AlwaysCritical, rule.Name)
	}

	pos := st.GetPosition()
	m, err := p.compiler.Compile(p.sources[pos.StartPos-1 : pos.EndPos])
	if err != nil {
		return p.errorf(st, "pattern compilation error: %v", err)
	}
	rule.Matcher = m

	if st2, ok := st.(*stmt.Expression); ok {
		st = st2.Expr
	}
	kind := CategorizeNode(st)
	if kind == KindNone {
		return p.errorf(st, "can't categorize pattern node: %T", st)
	}
	dst.RulesByKind[kind] = append(dst.RulesByKind[kind], rule)

	return nil
}

func (p *parser) errorf(n node.Node, format string, args ...interface{}) *parseError {
	pos := n.GetPosition()
	return &parseError{
		filename: p.filename,
		lineNum:  pos.StartLine,
		msg:      fmt.Sprintf(format, args...),
	}
}
