package parser

import (
	"fmt"
	"trivil/ast"
	"trivil/env"
	"trivil/lexer"
)

var _ = fmt.Printf

//=== statements

func (p *Parser) parseFn() *ast.Function {
	if p.trace {
		defer un(trace(p, "Функция"))
	}

	var n = &ast.Function{
		DeclBase: ast.DeclBase{Pos: p.pos},
	}

	p.expect(lexer.FN)

	if p.tok == lexer.LPAR { //receiver
		p.next()

		n.Recv = &ast.Param{
			DeclBase: ast.DeclBase{Pos: p.pos},
		}

		n.Recv.Name = p.parseIdent()
		p.expect(lexer.COLON)
		n.Recv.Typ = p.parseTypeRef()

		p.expect(lexer.RPAR)
	}

	n.Pos = p.pos // ident pos
	n.Name = p.parseIdent()

	if p.parseExportMark() {
		n.Exported = true
	}

	n.Typ = p.parseFuncType()

	if p.tok == lexer.MODIFIER {
		var mod = p.parseModifier()

		switch mod.Name {
		case "@внеш":
			n.External = true

		default:
			p.error(p.pos, "ПАР-ОШ-МОДИФИКАТОР", mod.Name)
		}
		/* перенесено в словарь
		for i, a := range mod.attrs {
			if a == "имя" {
				n.ExternalName = mod.values[i]
			} else {
				p.error(p.pos, "ПАР-ОШ-МОД-АТРИБУТ", a)
			}
		}
		*/
		n.Mod = mod
	} else {
		n.Seq = p.parseStatementSeq()
	}

	return n
}

func (p *Parser) parseFuncType() *ast.FuncType {
	if p.trace {
		defer un(trace(p, "Тип функции"))
	}

	var ft = &ast.FuncType{
		TypeBase: ast.TypeBase{Pos: p.pos},
	}

	p.expect(lexer.LPAR)

	p.parseParameters(ft)

	p.expect(lexer.RPAR)

	if p.tok == lexer.COLON {
		p.next()
		ft.ReturnTyp = p.parseTypeRef()
	}

	return ft
}

/* пока не используется
var skipToParam = tokens{
	lexer.EOF: true,

	lexer.RPAR:  true,
	lexer.COMMA: true,
}
*/

func (p *Parser) parseParameters(ft *ast.FuncType) {

	for p.tok != lexer.RPAR && p.tok != lexer.EOF {

		var param = &ast.Param{
			DeclBase: ast.DeclBase{Pos: p.pos},
		}

		param.Name = p.parseIdent()

		if p.tok == lexer.ASSIGN {
			param.Out = true
			p.next()
		} else {
			p.expect(lexer.COLON)
		}

		var variadic = p.tok == lexer.ELLIPSIS
		var variadic_pos = p.pos
		if variadic {
			p.next()
		}

		if p.tok == lexer.MUL {
			param.Typ = ast.TagPairType
			p.next()

		} else {
			param.Typ = p.parseTypeRef()
		}

		if variadic {
			param.Typ = &ast.VariadicType{
				TypeBase:   ast.TypeBase{Pos: variadic_pos},
				ElementTyp: param.Typ,
			}
		}

		ft.Params = append(ft.Params, param)

		if p.tok == lexer.RPAR {
			break
		}
		p.expect(lexer.COMMA)
	}

	for i, p := range ft.Params {
		if ast.IsVariadicType(p.Typ) && i != len(ft.Params)-1 {
			env.AddError(p.Pos, "ПАР-МЕСТО-ВАРИАДИК")
		}
	}

}
