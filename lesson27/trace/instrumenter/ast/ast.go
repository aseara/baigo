package ast

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"

	"golang.org/x/tools/go/ast/astutil"

	inst "github.com/aseara/baigo/lesson27/trace/instrumenter"
)

type instrumenter struct {
	trceImport string
	tracePkg   string
	traceFunc  string
}

// New factory method for inner instrumenter
func New(trceImport string, tracePkg string, traceFunc string) inst.Instrumenter {
	return instrumenter{
		trceImport: trceImport,
		tracePkg:   tracePkg,
		traceFunc:  traceFunc,
	}
}

func (a instrumenter) Instrument(filename string) ([]byte, error) {
	fset := token.NewFileSet()
	curAST, err := parser.ParseFile(fset, filename, nil, parser.ParseComments)
	if err != nil {
		return nil, fmt.Errorf("error parsing %s: %w", filename, err)
	}

	if !hasFuncDecl(curAST) {
		return nil, nil
	}

	astutil.AddImport(fset, curAST, a.trceImport)

	a.addDeferTraceIntoFuncDecls(curAST)

	buf := &bytes.Buffer{}
	err = format.Node(buf, fset, curAST)
	if err != nil {
		return nil, fmt.Errorf("error formatting new code: %w", err)
	}
	return buf.Bytes(), nil
}

func hasFuncDecl(f *ast.File) bool {
	for _, decl := range f.Decls {
		_, ok := decl.(*ast.FuncDecl)
		if ok {
			return true
		}
	}
	return false
}

func (a instrumenter) addDeferTraceIntoFuncDecls(f *ast.File) {
	for _, decl := range f.Decls {
		fd, ok := decl.(*ast.FuncDecl)
		if ok {
			_ = a.addDeferStmt(fd)
		}
	}
}

func (a instrumenter) addDeferStmt(fd *ast.FuncDecl) (added bool) {
	stmts := fd.Body.List

	for _, stmt := range stmts {
		ds, ok := stmt.(*ast.DeferStmt)
		if !ok {
			continue
		}
		ce, ok := ds.Call.Fun.(*ast.CallExpr)
		if !ok {
			continue
		}

		se, ok := ce.Fun.(*ast.SelectorExpr)
		if !ok {
			continue
		}

		x, ok := se.X.(*ast.Ident)
		if !ok {
			continue
		}
		if (x.Name == a.tracePkg) && (se.Sel.Name == a.traceFunc) {
			return false
		}
	}

	ds := &ast.DeferStmt{
		Call: &ast.CallExpr{
			Fun: &ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X: &ast.Ident{
						Name: a.tracePkg,
					},
					Sel: &ast.Ident{
						Name: a.traceFunc,
					},
				},
			},
		},
	}

	newList := make([]ast.Stmt, len(stmts)+1)
	copy(newList[1:], stmts)
	newList[0] = ds
	fd.Body.List = newList
	return true
}
