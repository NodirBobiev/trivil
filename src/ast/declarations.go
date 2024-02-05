package ast

import (
	"fmt"
	//	"trivil/env"
)

var _ = fmt.Printf

//====

type DeclBase struct {
	Pos      int
	Name     string
	Typ      Type
	Host     *Module
	Exported bool
}

func (n *DeclBase) DeclNode() {}

func (n *DeclBase) GetPos() int {
	return n.Pos
}

func (n *DeclBase) GetName() string {
	return n.Name
}

func (n *DeclBase) GetType() Type {
	return n.Typ
}

func (n *DeclBase) GetHost() *Module {
	return n.Host
}

func (n *DeclBase) SetHost(host *Module) {
	n.Host = host
}

func (n *DeclBase) IsExported() bool {
	return n.Exported
}

//====

type InvalidDecl struct {
	DeclBase
}

//=== описания

type Function struct {
	DeclBase
	Recv     *Param
	Seq      *StatementSeq
	External bool
	Mod      *Modifier
}

type VarDecl struct {
	DeclBase
	Init       Expr
	Later      bool
	AssignOnce bool
	OutParam   bool // если это выходной параметр
}

type ConstDecl struct {
	DeclBase
	Value Expr
}

type TypeDecl struct {
	DeclBase
}

//====

type StdFunction struct {
	DeclBase
	Method bool
}
