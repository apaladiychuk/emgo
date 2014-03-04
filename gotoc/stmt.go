package gotoc

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/token"
	"strconv"

	"code.google.com/p/go.tools/go/types"
)

func (cdd *CDD) ReturnStmt(w *bytes.Buffer, s *ast.ReturnStmt, resultT string, tup *types.Tuple) (end bool) {
	switch len(s.Results) {
	case 0:
		if resultT == "void" {
			w.WriteString("return;\n")
		} else {
			w.WriteString("goto __end;\n")
			end = true
		}

	case 1:
		w.WriteString("return ")
		cdd.Expr(w, s.Results[0], tup.At(0).Type())
		w.WriteString(";\n")

	default:
		w.WriteString("return (" + resultT + ") {")
		for i, e := range s.Results {
			if i > 0 {
				w.WriteString(", ")
			}
			cdd.Expr(w, e, tup.At(i).Type())
		}
		w.WriteString("};\n")
	}
	return
}

func (cdd *CDD) label(w *bytes.Buffer, label, suffix string) {
	cdd.il--
	cdd.indent(w)
	w.WriteString(label)
	w.WriteString(suffix)
	w.WriteString(":;\n")
	cdd.il++
}

func (cdd *CDD) Stmt(w *bytes.Buffer, stmt ast.Stmt, label, resultT string, tup *types.Tuple) (end bool, acds []*CDD) {
	updateEA := func(e bool, a []*CDD) {
		if e {
			end = true
		}
		acds = append(acds, a...)
	}

	cdd.Complexity++

	switch s := stmt.(type) {
	case *ast.AssignStmt:
		rhs := make([]string, len(s.Lhs))
		typ := make([]types.Type, len(s.Lhs))

		if len(s.Lhs) > 1 && len(s.Rhs) == 1 {
			// Tuple on RHS
			tup := cdd.gtc.ti.Types[s.Rhs[0]].Type.(*types.Tuple)
			tupName, _ := cdd.tupleName(tup)
			w.WriteString(tupName)
			tupName = "__tmp" + cdd.gtc.uniqueId()
			w.WriteString(" " + tupName + " = ")
			cdd.Expr(w, s.Rhs[0], nil)
			w.WriteString(";\n")
			cdd.indent(w)
			for i, n := 0, tup.Len(); i < n; i++ {
				rhs[i] = tupName + "._" + strconv.Itoa(i)
				if s.Tok == token.DEFINE {
					typ[i] = tup.At(i).Type()
				}
			}
		} else {
			for i, e := range s.Rhs {
				var t types.Type
				if s.Tok == token.DEFINE {
					t = cdd.gtc.ti.Types[e].Type
				} else {
					t = cdd.gtc.ti.Types[s.Lhs[i]].Type
				}
				rhs[i] = cdd.ExprStr(e, t)
				typ[i] = t
			}
		}

		lhs := make([]string, len(s.Lhs))

		if s.Tok == token.DEFINE {
			for i, e := range s.Lhs {
				name := cdd.NameStr(cdd.object(e.(*ast.Ident)), true)
				t, dim, a := cdd.TypeStr(typ[i])
				acds = append(acds, a...)
				lhs[i] = t + " " + dimFuncPtr(name, dim)
			}
		} else {
			for i, e := range s.Lhs {
				lhs[i] = cdd.ExprStr(e, nil)
			}
		}

		if len(s.Rhs) == len(s.Lhs) && len(s.Lhs) > 1 && s.Tok != token.DEFINE {
			for i, t := range typ {
				if i > 0 {
					cdd.indent(w)
				}
				dim, a := cdd.Type(w, t)
				acds = append(acds, a...)
				tmp := "__tmp" + cdd.gtc.uniqueId()
				w.WriteString(" " + dimFuncPtr(tmp, dim))
				w.WriteString(" = " + rhs[i] + ";\n")
				rhs[i] = tmp
			}
			cdd.indent(w)
		}

		var atok string
		switch s.Tok {
		case token.DEFINE:
			atok = " = "

		case token.AND_NOT_ASSIGN:
			atok = " &= ~("

		default:
			atok = " " + s.Tok.String() + " "
		}

		for i := 0; i < len(lhs); i++ {
			if i > 0 {
				cdd.indent(w)
			}
			w.WriteString(lhs[i])
			w.WriteString(atok)
			w.WriteString(rhs[i])
			if s.Tok == token.AND_NOT_ASSIGN {
				w.WriteString(");\n")
			} else {
				w.WriteString(";\n")
			}
		}

	case *ast.ExprStmt:
		cdd.Expr(w, s.X, nil)
		w.WriteString(";\n")

	case *ast.IfStmt:
		if s.Init != nil {
			w.WriteString("{\n")
			cdd.il++
			cdd.indent(w)
			updateEA(cdd.Stmt(w, s.Init, "", resultT, tup))
			cdd.indent(w)
		}

		w.WriteString("if (")
		cdd.Expr(w, s.Cond, nil)
		w.WriteString(") ")
		updateEA(cdd.BlockStmt(w, s.Body, resultT, tup))
		if s.Else == nil {
			w.WriteByte('\n')
		} else {
			w.WriteString(" else ")
			updateEA(cdd.Stmt(w, s.Else, "", resultT, tup))
		}

		if s.Init != nil {
			cdd.il--
			cdd.indent(w)
			w.WriteString("}\n")
		}

	case *ast.IncDecStmt:
		w.WriteString(s.Tok.String())
		w.WriteByte('(')
		cdd.Expr(w, s.X, nil)
		w.WriteString(");\n")

	case *ast.BlockStmt:
		updateEA(cdd.BlockStmt(w, s, resultT, tup))
		w.WriteByte('\n')

	case *ast.ForStmt:
		if s.Init != nil {
			w.WriteString("{\n")
			cdd.il++
			cdd.indent(w)
			updateEA(cdd.Stmt(w, s.Init, "", resultT, tup))
		}
		if label != "" {
			cdd.label(w, label, "_continue")
		}
		if s.Init != nil {
			cdd.indent(w)
		}
		w.WriteString("while (")
		if s.Cond != nil {
			cdd.Expr(w, s.Cond, nil)
		} else {
			w.WriteString("true")
		}
		w.WriteString(") ")

		if s.Post != nil {
			w.WriteString("{\n")
			cdd.il++
			cdd.indent(w)
		}
		updateEA(cdd.BlockStmt(w, s.Body, resultT, tup))
		w.WriteByte('\n')

		if s.Post != nil {
			cdd.indent(w)
			updateEA(cdd.Stmt(w, s.Post, "", resultT, tup))
			cdd.il--
			cdd.indent(w)
			w.WriteString("}\n")
		}

		if s.Init != nil {
			cdd.il--
			cdd.indent(w)
			w.WriteString("}\n")
		}

		if label != "" {
			cdd.label(w, label, "_break")
		}

	case *ast.ReturnStmt:
		if cdd.ReturnStmt(w, s, resultT, tup) {
			end = true
		}

	case *ast.SwitchStmt:
		w.WriteString("switch(0) {\n")
		cdd.indent(w)
		w.WriteString("case 0:;\n")
		cdd.il++

		if s.Init != nil {
			cdd.indent(w)
			updateEA(cdd.Stmt(w, s.Init, "", resultT, tup))
		}

		cdd.indent(w)

		var typ types.Type
		if s.Tag != nil {
			typ = cdd.gtc.ti.Types[s.Tag].Type
			cdd.Type(w, typ)
			w.WriteString(" __tag = ")
			cdd.Expr(w, s.Tag, typ)
			w.WriteString(";\n")
		} else {
			typ = types.Typ[types.Bool]
			w.WriteString("bool __tag = true;\n")
		}

		for _, stmt := range s.Body.List {
			cdd.indent(w)
			
			cs := stmt.(*ast.CaseClause)
			if cs.List != nil {
				w.WriteString("if (")
				for i, e := range cs.List {
					if i != 0 {
						w.WriteString(" || ")
					}
					eq(w, "__tag", "==", cdd.ExprStr(e, typ), typ)
				}
				w.WriteString(") ")
			}
			w.WriteString("{\n")
			cdd.il++

			var ftLabel string
			for _, s := range cs.Body {
				cdd.indent(w)
				if bs, ok := s.(*ast.BranchStmt); ok && bs.Tok == token.FALLTHROUGH {
					if ftLabel == "" {
						ftLabel = "__fallthrough" + cdd.gtc.uniqueId()
					}
					w.WriteString("goto " + ftLabel + ";\n")
				} else {
					updateEA(cdd.Stmt(w, s, "", resultT, tup))
				}
			}

			cdd.indent(w)
			w.WriteString("break;\n")

			cdd.il--
			cdd.indent(w)
			w.WriteString("}\n")
			if ftLabel != "" {
				cdd.il--
				cdd.indent(w)
				w.WriteString(ftLabel + ":;\n")
				cdd.il++
			}
		}

		cdd.il--
		cdd.indent(w)
		w.WriteString("}\n")

		if label != "" {
			cdd.label(w, label, "_break")
		}

	case *ast.BranchStmt:
		if s.Label == nil {
			w.WriteString(s.Tok.String())
		} else {
			w.WriteString("goto " + s.Label.Name)
			switch s.Tok {
			case token.BREAK:
				w.WriteString("_break")
			case token.CONTINUE:
				w.WriteString("_continue")
			}
		}
		w.WriteString(";\n")

	default:
		fmt.Fprintf(w, "#<%T>#", stmt)
	}
	return
}

func (cdd *CDD) BlockStmt(w *bytes.Buffer, bs *ast.BlockStmt, resultT string, tup *types.Tuple) (end bool, acds []*CDD) {
	updateEA := func(e bool, a []*CDD) {
		if e {
			end = true
		}
		acds = append(acds, a...)
	}

	w.WriteString("{\n")
	cdd.il++
	for _, stmt := range bs.List {
		switch s := stmt.(type) {
		case *ast.DeclStmt:
			cdds := cdd.gtc.Decl(s.Decl, cdd.il)
			for _, c := range cdds {
				for u, typPtr := range c.BodyUses {
					cdd.BodyUses[u] = typPtr
				}
				w.Write(c.Decl)
			}
			for _, c := range cdds {
				w.Write(c.Def)
			}

		case *ast.LabeledStmt:
			cdd.label(w, s.Label.Name, "")
			cdd.indent(w)
			updateEA(cdd.Stmt(w, s.Stmt, s.Label.Name, resultT, tup))

		default:
			cdd.indent(w)
			updateEA(cdd.Stmt(w, s, "", resultT, tup))
		}
	}
	cdd.il--
	cdd.indent(w)
	w.WriteString("}")
	return
}
