package parser

import (
	"fmt"
	"strconv"
	"strings"
	"unicode/utf8"

	"trivil/ast"
	"trivil/lexer"
)

var _ = fmt.Printf

//=== приоритеты

const lowestPrecedence = 0

func precedence(tok lexer.Token) int {
	switch tok {
	case lexer.OR:
		return 1
	case lexer.AND:
		return 2
	case lexer.EQ, lexer.NEQ, lexer.LSS, lexer.LEQ, lexer.GTR, lexer.GEQ,
		lexer.OFTYPE:
		return 3
	case lexer.ADD, lexer.SUB, lexer.BITOR, lexer.BITXOR:
		return 4
	case lexer.MUL, lexer.QUO, lexer.REM, lexer.BITAND, lexer.SHL, lexer.SHR:
		return 5
	default:
		return lowestPrecedence
	}
}

//=== выражения

func (p *Parser) parseExpression() ast.Expr {
	if p.trace {
		defer un(trace(p, "Выражение"))
	}

	return p.parseBinaryExpression(lowestPrecedence + 1)
}

func (p *Parser) parseBinaryExpression(prec int) ast.Expr {
	if p.trace {
		defer un(trace(p, "Выражение бинарное"))
	}

	var x = p.parseUnaryExpression()
	for {
		op := p.tok
		var pos = p.pos
		opPrec := precedence(op)
		if opPrec < prec {
			return x
		}

		if op == lexer.OFTYPE {
			x = p.parseOfTypeExpression(x)
		} else {
			p.next()
			var y = p.parseBinaryExpression(opPrec + 1)
			x = &ast.BinaryExpr{
				ExprBase: ast.ExprBase{Pos: pos},
				X:        x,
				Op:       op,
				Y:        y,
			}
		}
	}
}

func (p *Parser) parseUnaryExpression() ast.Expr {
	if p.trace {
		defer un(trace(p, "Выражение унарное"))
	}

	switch p.tok {
	case lexer.SUB, lexer.NOT, lexer.BITNOT:
		var pos = p.pos
		var op = p.tok
		p.next()
		var x = p.parseUnaryExpression()
		return &ast.UnaryExpr{
			ExprBase: ast.ExprBase{Pos: pos},
			Op:       op,
			X:        x,
		}
	case lexer.ADD:
		p.next()
		return p.parseUnaryExpression()
	}

	var x = p.parsePrimaryExpression()

	// check ?
	return x
}

func (p *Parser) parseOfTypeExpression(x ast.Expr) ast.Expr {

	var pos = p.pos
	p.next()

	var n = &ast.OfTypeExpr{
		ExprBase:  ast.ExprBase{Pos: pos},
		X:         x,
		TargetTyp: p.parseTypeRef(),
	}

	return n
}

func (p *Parser) parsePrimaryExpression() ast.Expr {
	if p.trace {
		defer un(trace(p, "Выражение первичное"))
	}

	var x ast.Expr

	switch p.tok {
	case lexer.INT:
		var l = &ast.LiteralExpr{
			ExprBase: ast.ExprBase{Pos: p.pos},
			Kind:     ast.Lit_Int,
		}
		if strings.HasPrefix(p.lit, "0x") {
			u, err := strconv.ParseUint(p.lit, 0, 64)
			if err != nil {
				p.error(p.pos, "ПАР-ОШ-ЛИТЕРАЛ", err.Error())
			} else {
				l.Kind = ast.Lit_Word
				l.WordVal = u
			}
		} else {
			i, err := strconv.ParseInt(p.lit, 0, 64)
			if err != nil {
				p.error(p.pos, "ПАР-ОШ-ЛИТЕРАЛ", err.Error())
			} else {
				l.IntVal = i
			}
		}
		p.next()
		x = l
	case lexer.SYMBOL:
		var l = &ast.LiteralExpr{
			ExprBase: ast.ExprBase{Pos: p.pos},
			Kind:     ast.Lit_Symbol,
		}
		if !utf8.ValidString(p.lit) {
			p.error(p.pos, "ПАР-ОШ-ЛИТЕРАЛ", "неверная кодировка символа")
		} else {
			s, err := strconv.Unquote("'" + p.lit + "'")
			if err != nil {
				p.error(p.pos, "ПАР-ОШ-ЛИТЕРАЛ", "unqoute "+err.Error())
				//panic("assert - unquote: " + err.Error())
			}
			r, _ := utf8.DecodeRuneInString(s)
			l.WordVal = uint64(r)
		}
		p.next()
		x = l
	case lexer.FLOAT:
		var l = &ast.LiteralExpr{
			ExprBase: ast.ExprBase{Pos: p.pos},
			Kind:     ast.Lit_Float,
		}
		_, err := strconv.ParseFloat(p.lit, 64)
		if err != nil {
			p.error(p.pos, "ПАР-ОШ-ЛИТЕРАЛ", err.Error())
			l.FloatStr = "1.0"
		} else {
			l.FloatStr = p.lit
		}
		p.next()
		x = l
	case lexer.STRING:
		var l = &ast.LiteralExpr{
			ExprBase: ast.ExprBase{Pos: p.pos},
			Kind:     ast.Lit_String,
		}
		if !utf8.ValidString(p.lit) {
			p.error(p.pos, "ПАР-ОШ-ЛИТЕРАЛ", "неверная кодировка строки")
		} else {
			s, err := strconv.Unquote("\"" + p.lit + "\"")
			if err != nil {
				p.error(p.pos, "ПАР-ОШ-ЛИТЕРАЛ", "unqoute "+err.Error())
				//panic("assert - unquote: " + err.Error())
			}
			l.StrVal = make([]rune, utf8.RuneCountInString(s))

			var b = []byte(s)
			var i = 0
			for len(b) > 0 {
				r, size := utf8.DecodeRune(b)
				l.StrVal[i] = r
				i++
				b = b[size:]
			}
			//fmt.Printf("!! %v\n", l.StrVal)
		}
		p.next()
		x = l
	case lexer.IDENT:
		x = &ast.IdentExpr{
			ExprBase: ast.ExprBase{Pos: p.pos},
			Name:     p.lit,
		}
		p.next()
	case lexer.LPAR:
		p.next()
		x = p.parseExpression()
		p.expect(lexer.RPAR)
	default:
		p.error(p.pos, "ПАР-ОШ-ОПЕРАНД", p.tok.String())
		return &ast.InvalidExpr{}
	}

	for {
		switch p.tok {
		case lexer.DOT:
			x = p.parseSelector(x)
		case lexer.LPAR:
			x = p.parseArguments(x)
		case lexer.LCONV:
			x = p.parseConversion(x)
		case lexer.LBRACK:
			x = p.parseIndex(x)
		case lexer.LBRACE:
			if p.lex.WhitespaceBefore('{') {
				return x
			}
			x = p.parseClassComposite(x)
		case lexer.NOTNIL:
			x = p.parseNotNil(x)
		default:
			return x
		}
	}
}

func (p *Parser) parseSelector(x ast.Expr) ast.Expr {
	if p.trace {
		defer un(trace(p, "Селектор"))
	}
	p.next()

	var n = &ast.SelectorExpr{
		ExprBase: ast.ExprBase{Pos: p.pos},
		X:        x,
	}

	n.Name = p.parseIdent()

	return n
}

func (p *Parser) parseArguments(x ast.Expr) ast.Expr {
	if p.trace {
		defer un(trace(p, "Аргументы"))
	}

	var n = &ast.CallExpr{
		ExprBase: ast.ExprBase{Pos: p.pos},
		X:        x,
		Args:     make([]ast.Expr, 0),
	}

	p.expect(lexer.LPAR)

	for p.tok != lexer.RPAR && p.tok != lexer.EOF {

		var expr = p.parseExpression()

		if p.tok == lexer.ELLIPSIS {
			var u = &ast.UnfoldExpr{
				ExprBase: ast.ExprBase{Pos: p.pos},
				X:        expr,
			}
			expr = u
			p.next()
		}

		n.Args = append(n.Args, expr)

		if p.tok == lexer.RPAR {
			break
		}
		p.expect(lexer.COMMA)
	}

	p.expect(lexer.RPAR)

	return n
}

func (p *Parser) parseConversion(x ast.Expr) ast.Expr {
	if p.trace {
		defer un(trace(p, "Конверсия"))
	}

	var n = &ast.ConversionExpr{
		ExprBase: ast.ExprBase{Pos: p.pos},
		X:        x,
	}

	p.next()

	if p.tok == lexer.CAUTION {
		p.next()
		n.Caution = true
		p.cautionUse++
		if !p.caution {
			p.error(p.pos, "ПАР-ОШ-ИСП-ОСТОРОЖНО")
		}
	}

	n.TargetTyp = p.parseTypeRef()

	p.expect(lexer.RPAR)

	return n
}

func (p *Parser) parseIndex(x ast.Expr) ast.Expr {
	if p.trace {
		defer un(trace(p, "Индексация"))
	}

	var n = &ast.GeneralBracketExpr{
		ExprBase: ast.ExprBase{Pos: p.pos},
		X:        x,
		Composite: &ast.ArrayCompositeExpr{
			ExprBase: ast.ExprBase{Pos: p.pos},
			Indexes:  make([]ast.Expr, 0),
			Values:   make([]ast.Expr, 0),
			Length:   -1,
			MaxIndex: -1,
		},
	}

	p.expect(lexer.LBRACK)

	var inx ast.Expr
	var val ast.Expr

	for p.tok != lexer.RBRACK && p.tok != lexer.EOF {

		if p.tok == lexer.MUL {
			p.next()
			p.expect(lexer.COLON)
			n.Composite.Default = p.parseExpression()
		} else {
			inx = p.parseExpression()

			if p.tok == lexer.COLON {
				p.next()
				val = p.parseExpression()

				if isVectorProperty(inx, ast.StdLen) {
					n.Composite.LenExpr = val
					val = nil
				} else if isVectorProperty(inx, ast.VectorAllocate) {
					n.Composite.CapExpr = val
					val = nil
				} else {
					n.Composite.Indexes = append(n.Composite.Indexes, inx)
					n.Composite.Values = append(n.Composite.Values, val)
				}
			} else {
				n.Composite.Values = append(n.Composite.Values, inx)
			}
		}

		if p.tok == lexer.RBRACK {
			break
		}
		p.expect(lexer.COMMA)
	}

	p.expect(lexer.RBRACK)

	p.checkElements(n.Composite)

	return n
}

func isVectorProperty(p ast.Expr, name string) bool {
	ident, ok := p.(*ast.IdentExpr)
	if !ok {
		return false
	}

	return ident.Name == name
}

func (p *Parser) checkElements(n *ast.ArrayCompositeExpr) {

	// если все элементы - это пары
	if len(n.Indexes) == len(n.Values) {
		return
	}

	// если ни одной пары
	if len(n.Indexes) == 0 && n.LenExpr == nil && n.CapExpr == nil && n.Default == nil {
		return
	}

	p.error(n.Pos, "ПАР-СМЕСЬ-МАССИВ")
}

//=== class composite

func (p *Parser) parseClassComposite(x ast.Expr) ast.Expr {
	if p.trace {
		defer un(trace(p, "Композит класса"))
	}

	var n = &ast.ClassCompositeExpr{
		ExprBase: ast.ExprBase{Pos: p.pos},
		X:        x,
		Values:   make([]ast.ValuePair, 0),
	}

	p.expect(lexer.LBRACE)

	for p.tok != lexer.RBRACE && p.tok != lexer.EOF {

		var vp = ast.ValuePair{Pos: p.pos}

		vp.Name = p.parseIdent()
		p.expect(lexer.COLON)
		vp.Value = p.parseExpression()

		n.Values = append(n.Values, vp)

		if p.tok == lexer.RBRACE {
			break
		}
		p.expect(lexer.COMMA)
	}

	p.expect(lexer.RBRACE)

	return n
}

func (p *Parser) parseNotNil(x ast.Expr) ast.Expr {

	var n = &ast.NotNilExpr{
		ExprBase: ast.ExprBase{Pos: p.pos},
		X:        x,
	}

	p.next()

	return n
}
