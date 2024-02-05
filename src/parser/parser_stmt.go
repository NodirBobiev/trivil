package parser

import (
	"fmt"
	"trivil/ast"
	"trivil/lexer"
)

var _ = fmt.Printf

var validSimpleStmToken = tokens{
	lexer.IDENT: true,
	lexer.LPAR:  true,

	// literals
	lexer.INT:    true,
	lexer.FLOAT:  true,
	lexer.STRING: true,
	lexer.SYMBOL: true,

	// unary ops
	lexer.ADD: true,
	lexer.SUB: true,
	lexer.NOT: true,
}

var skipToStatement = tokens{
	lexer.EOF: true,

	lexer.RBRACE: true,

	lexer.IF:     true,
	lexer.WHILE:  true,
	lexer.RETURN: true,
	//TODO
}

var endStatementSeq = tokens{
	lexer.EOF:    true,
	lexer.RBRACE: true,
}

var endWhenCase = tokens{
	lexer.EOF:    true,
	lexer.RBRACE: true,

	lexer.WHEN:  true,
	lexer.OTHER: true,
}

//=== statements

func (p *Parser) parseStatementSeq() *ast.StatementSeq {
	if p.trace {
		defer un(trace(p, "Список операторов"))
	}

	if p.tok != lexer.LBRACE {
		p.expect(lexer.LBRACE)
		return &ast.StatementSeq{
			StatementBase: ast.StatementBase{Pos: p.pos},
			Statements:    make([]ast.Statement, 0),
		}
	}
	p.next()

	var n = p.parseStatementList(endStatementSeq)

	p.expect(lexer.RBRACE)

	return n
}

func (p *Parser) parseStatementList(stop tokens) *ast.StatementSeq {

	var n = &ast.StatementSeq{
		StatementBase: ast.StatementBase{Pos: p.pos},
		Statements:    make([]ast.Statement, 0),
	}

	for !stop[p.tok] {
		var s = p.parseStatement()
		if s != nil {
			n.Statements = append(n.Statements, s)
		}

		if stop[p.tok] {
			break
		}
		p.sep()
	}

	return n
}

func (p *Parser) parseStatement() ast.Statement {
	if p.trace {
		defer un(trace(p, "Оператор"))
	}

	switch p.tok {
	case lexer.VAR:
		return &ast.DeclStatement{
			StatementBase: ast.StatementBase{Pos: p.pos},
			D:             p.parseVarDecl(),
		}
	case lexer.IF:
		return p.parseIf()
	case lexer.GUARD:
		return p.parseGuard()
	case lexer.SELECT:
		return p.parseSelect()
	case lexer.WHILE:
		return p.parseWhile()
	case lexer.CYCLE:
		return p.parseCycle()
	case lexer.RETURN:
		return p.parseReturn()
	case lexer.BREAK:
		return p.parseBreak()
	case lexer.CRASH:
		return p.parseCrash()

	default:
		if validSimpleStmToken[p.tok] {
			return p.parseSimpleStatement()
		}
		p.error(p.pos, "ПАР-ОШ-ОПЕРАТОР", p.tok.String())
		p.skipTo((skipToStatement))
	}

	return nil
}

func (p *Parser) parseSimpleStatement() ast.Statement {
	if p.trace {
		defer un(trace(p, "Простой оператор"))
	}

	var expr = p.parseExpression()

	switch p.tok {
	case lexer.ASSIGN:
		return p.parseAssign(expr)
	case lexer.INC:
		var n = &ast.IncStatement{
			StatementBase: ast.StatementBase{Pos: p.pos},
			L:             expr,
		}
		p.next()
		return n
	case lexer.DEC:
		var n = &ast.DecStatement{
			StatementBase: ast.StatementBase{Pos: p.pos},
			L:             expr,
		}
		p.next()
		return n

	default:
		var s = &ast.ExprStatement{
			StatementBase: ast.StatementBase{Pos: p.pos},

			X: expr,
		}

		return s
	}
}

func (p *Parser) parseAssign(l ast.Expr) ast.Statement {
	if p.trace {
		defer un(trace(p, "Оператор присваивания"))
	}

	var n = &ast.AssignStatement{
		StatementBase: ast.StatementBase{Pos: p.pos},
		L:             l,
	}

	p.next()
	n.R = p.parseExpression()

	return n
}

//====

func (p *Parser) parseIf() ast.Statement {
	if p.trace {
		defer un(trace(p, "Оператор если"))
	}

	var n = &ast.If{
		StatementBase: ast.StatementBase{Pos: p.pos},
	}

	p.next()
	n.Cond = p.parseExpression()
	n.Then = p.parseStatementSeq()

	if p.tok != lexer.ELSE {
		return n
	}

	p.next()

	if p.tok == lexer.IF {
		n.Else = p.parseIf()
	} else {
		n.Else = p.parseStatementSeq()
	}

	return n
}

func (p *Parser) parseWhile() ast.Statement {
	if p.trace {
		defer un(trace(p, "Оператор пока"))
	}

	var n = &ast.While{
		StatementBase: ast.StatementBase{Pos: p.pos},
	}

	p.next()
	n.Cond = p.parseExpression()
	n.Seq = p.parseStatementSeq()

	return n
}

// Не делаю пока возможности задавать тип переменных
func (p *Parser) parseCycle() ast.Statement {
	if p.trace {
		defer un(trace(p, "Оператор цикла"))
	}

	var n = &ast.Cycle{
		StatementBase: ast.StatementBase{Pos: p.pos},
	}

	p.next()

	if p.tok == lexer.LBRACK {
		p.next()
		n.IndexVar = &ast.VarDecl{
			DeclBase: ast.DeclBase{
				Pos:  p.pos,
				Name: p.parseIdent(),
			},
			AssignOnce: true,
		}
		p.expect(lexer.RBRACK)
	}

	if p.tok == lexer.IDENT {
		n.ElementVar = &ast.VarDecl{
			DeclBase: ast.DeclBase{
				Pos:  p.pos,
				Name: p.parseIdent(),
			},
			AssignOnce: true,
		}
	} else if n.IndexVar == nil {
		p.expect(lexer.IDENT)
	}

	p.expect(lexer.AMONG)

	n.Expr = p.parseExpression()

	n.Seq = p.parseStatementSeq()

	return n
}

func (p *Parser) parseReturn() ast.Statement {
	if p.trace {
		defer un(trace(p, "Оператор вернуть"))
	}

	var n = &ast.Return{
		StatementBase: ast.StatementBase{Pos: p.pos},
	}

	p.next()

	if p.afterNL || p.tok == lexer.SEMI || p.tok == lexer.RBRACE {
		return n
	}

	n.X = p.parseExpression()

	return n
}

func (p *Parser) parseBreak() ast.Statement {
	if p.trace {
		defer un(trace(p, "Оператор прервать"))
	}

	var n = &ast.Break{
		StatementBase: ast.StatementBase{Pos: p.pos},
	}

	p.next()

	return n
}

func (p *Parser) parseCrash() ast.Statement {
	if p.trace {
		defer un(trace(p, "Оператор авария"))
	}

	var n = &ast.Crash{
		StatementBase: ast.StatementBase{Pos: p.pos},
	}

	p.next()
	p.expect(lexer.LPAR)
	n.X = p.parseExpression()
	p.expect(lexer.RPAR)

	return n
}

func (p *Parser) parseGuard() ast.Statement {
	if p.trace {
		defer un(trace(p, "Оператор надо"))
	}

	var n = &ast.Guard{
		StatementBase: ast.StatementBase{Pos: p.pos},
	}

	p.next()

	n.Cond = p.parseExpression()

	p.expect(lexer.ELSE)

	switch p.tok {
	case lexer.RETURN:
		n.Else = p.parseReturn()
	case lexer.BREAK:
		n.Else = p.parseBreak()
	case lexer.CRASH:
		n.Else = p.parseCrash()
	default:
		n.Else = p.parseStatementSeq()
	}

	return n
}

//==== оператор выбора

func (p *Parser) parseSelect() ast.Statement {
	p.next()
	if p.tok == lexer.VAR || p.tok == lexer.TYPE {
		return p.parseSelectType()
	}
	return p.parseSelectExpr()
}

//==== оператор выбора по выражению

func (p *Parser) parseSelectExpr() ast.Statement {
	if p.trace {
		defer un(trace(p, "Оператор выбора по выражению"))
	}
	var n = &ast.Select{
		StatementBase: ast.StatementBase{Pos: p.pos},
	}

	if p.tok == lexer.LBRACE {
		p.next()
	} else {
		n.X = p.parseExpression()
		p.expect(lexer.LBRACE)
	}

	for p.tok == lexer.WHEN {
		var c = p.parseCaseExpr()
		n.Cases = append(n.Cases, c)
	}

	if p.tok == lexer.OTHER {
		p.next()
		n.Else = p.parseStatementList(endStatementSeq)
	}
	p.expect(lexer.RBRACE)

	return n
}

func (p *Parser) parseCaseExpr() *ast.Case {
	if p.trace {
		defer un(trace(p, "выбор когда"))
	}

	var c = &ast.Case{
		StatementBase: ast.StatementBase{Pos: p.pos},
		Exprs:         make([]ast.Expr, 0),
	}
	p.next()

	for {
		var x = p.parseExpression()
		c.Exprs = append(c.Exprs, x)
		if p.tok != lexer.COMMA {
			break
		}
		p.next()
	}
	p.expect(lexer.COLON)

	c.Seq = p.parseStatementList(endWhenCase)

	return c
}

//==== оператор выбора по типу

func (p *Parser) parseSelectType() ast.Statement {
	if p.trace {
		defer un(trace(p, "Оператор выбора по типу"))
	}

	var n = &ast.SelectType{
		StatementBase: ast.StatementBase{Pos: p.pos},
	}

	var varPos = 0

	if p.tok == lexer.VAR {
		p.next()
		varPos = p.pos
		n.VarIdent = p.parseIdent()
		p.expect(lexer.COLON)
	}
	p.expect(lexer.TYPE)

	n.X = p.parseExpression()
	p.expect(lexer.LBRACE)

	for p.tok == lexer.WHEN {
		var c = p.parseCaseType()
		if n.VarIdent != "" {
			c.Var = &ast.VarDecl{
				DeclBase: ast.DeclBase{
					Pos:  varPos,
					Name: n.VarIdent,
				},
				AssignOnce: true,
			}
		}

		n.Cases = append(n.Cases, c)
	}

	if p.tok == lexer.OTHER {
		p.next()
		n.Else = p.parseStatementList(endStatementSeq)
	}
	p.expect(lexer.RBRACE)

	return n
}

func (p *Parser) parseCaseType() *ast.CaseType {
	if p.trace {
		defer un(trace(p, "выбор когда по типу"))
	}

	var c = &ast.CaseType{
		StatementBase: ast.StatementBase{Pos: p.pos},
		Types:         make([]ast.Type, 0),
	}
	p.next()

	for {
		var x = p.parseTypeRef()
		c.Types = append(c.Types, x)
		if p.tok != lexer.COMMA {
			break
		}
		p.next()
	}
	p.expect(lexer.COLON)

	c.Seq = p.parseStatementList(endWhenCase)

	return c
}
