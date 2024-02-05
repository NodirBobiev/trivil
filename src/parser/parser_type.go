package parser

import (
	"fmt"
	"trivil/ast"
	"trivil/env"
	"trivil/lexer"
)

var _ = fmt.Printf

//=== types

func (p *Parser) parseTypeRef() ast.Type {
	if p.trace {
		defer un(trace(p, "Cсылка на тип"))
	}

	var maybe *ast.MayBeType = nil

	if p.tok == lexer.MAYBE {
		maybe = &ast.MayBeType{
			TypeBase: ast.TypeBase{Pos: p.pos},
		}
		p.next()
	}

	var t = &ast.TypeRef{
		TypeBase: ast.TypeBase{Pos: p.pos},
	}

	var s = p.parseIdent()

	if p.tok == lexer.DOT {
		p.next()
		t.ModuleName = s
		t.TypeName = p.parseIdent()
	} else {
		t.TypeName = s
	}

	if maybe != nil {
		maybe.Typ = t
		return maybe
	}

	return t
}

func (p *Parser) parseTypeDecl() *ast.TypeDecl {
	if p.trace {
		defer un(trace(p, "Описание типа"))
	}

	p.next()

	var n = &ast.TypeDecl{
		DeclBase: ast.DeclBase{Pos: p.pos},
	}

	n.Name = p.parseIdent()
	if p.parseExportMark() {
		n.Exported = true
	}

	p.expect(lexer.EQ)
	n.Typ = p.parseTypeDef()

	return n
}

func (p *Parser) parseTypeDef() ast.Type {

	switch p.tok {
	case lexer.LBRACK:
		return p.parseVectorType()
	case lexer.CLASS:
		return p.parseClassType()
	default:
		return p.parseTypeRef()
		/*
			p.error(p.pos, "ПАР-ОШ-ОП-ТИПА", p.tok.String())
			return &ast.InvalidType{
				TypeBase: ast.TypeBase{Pos: p.pos},
			}
		*/
	}
}

func (p *Parser) parseVectorType() *ast.VectorType {

	var t = &ast.VectorType{
		TypeBase: ast.TypeBase{Pos: p.pos},
	}

	p.next()
	p.expect(lexer.RBRACK)

	t.ElementTyp = p.parseTypeRef()

	return t
}

//==== класс

func (p *Parser) parseClassType() *ast.ClassType {

	var t = &ast.ClassType{
		TypeBase: ast.TypeBase{Pos: p.pos},
		Members:  make(map[string]ast.Decl),
	}

	p.next()

	if p.tok == lexer.LPAR {
		p.next()
		t.BaseTyp = p.parseTypeRef()
		p.expect(lexer.RPAR)
	}

	p.expect(lexer.LBRACE)

	for p.tok != lexer.RBRACE && p.tok != lexer.EOF {

		var f = p.parseField()
		t.Fields = append(t.Fields, f)

		if p.tok == lexer.RBRACE {
			break
		}
		p.sep()
	}

	p.expect(lexer.RBRACE)

	return t
}

func (p *Parser) parseField() *ast.Field {

	var n = &ast.Field{
		DeclBase: ast.DeclBase{
			Pos: p.pos,
			//Host: p.module,
		},
	}

	n.Name = p.parseIdent()
	if p.parseExportMark() {
		n.Exported = true
	}

	if p.tok == lexer.COLON {
		p.next()
		n.Typ = p.parseTypeRef()
	}

	if p.tok == lexer.EQ {
		n.AssignOnce = true
		p.next()
	} else if p.tok == lexer.ASSIGN {
		p.next()
	} else {
		env.AddError(p.pos, "ПАР-ПЕРЕМ-ИНИТ")
		return n
	}

	if p.tok == lexer.LATER {
		n.Later = true
		p.next()
	} else {
		n.Init = p.parseExpression()
	}

	return n
}
