package main

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"os"
	"regexp"
	"strings"

	"golang.org/x/tools/imports"
)

var fset = token.NewFileSet()

func main() {
	filename := os.Getenv("GOFILE")
	af, err := parser.ParseFile(fset, filename, nil, 0)

	if err != nil {
		log.Fatal(err)
	}

	af.Doc = nil
	af.Imports = nil
	af.Comments = nil

	var newDecl []ast.Decl

	for _, d := range af.Decls {
		fd, ok := d.(*ast.FuncDecl)

		if !ok {
			continue
		}

		if fd.Recv == nil || !fd.Name.IsExported() {
			continue
		}

		typ := fd.Type

		if len(typ.Params.List) < 1 {
			continue
		}

		arg0 := typ.Params.List[0]
		arg0Name := arg0.Names[0].Name
		arg0Type := arg0.Type.(*ast.SelectorExpr)

		if arg0Name != "ctx" || arg0Type.Sel.Name != "Context" {
			continue
		}

		recvType := fd.Recv.List[0].Type.(*ast.StarExpr).X.(*ast.Ident)
		recvType.Name = "ClientWithoutContext"

		origParams := []ast.Expr{
			&ast.CallExpr{
				Fun: &ast.SelectorExpr{
					X:   ast.NewIdent("context"),
					Sel: ast.NewIdent("Background"),
				},
			},
		}

		paramsList1 := typ.Params.List[1:len(typ.Params.List)]

		for _, p := range paramsList1 {
			origParams = append(origParams, &ast.BasicLit{
				Kind:  token.IDENT,
				Value: p.Names[0].Name,
			})
		}

		fd.Body = &ast.BlockStmt{
			List: []ast.Stmt{
				&ast.ReturnStmt{
					Results: []ast.Expr{
						&ast.CallExpr{
							Fun: &ast.SelectorExpr{
								X:   ast.NewIdent("client.withCtx"),
								Sel: ast.NewIdent(fd.Name.Name),
							},
							Args: origParams,
						},
					},
				},
			},
		}

		typ.Params.List = paramsList1
		newDecl = append(newDecl, fd)
	}

	af.Decls = newDecl
	var out bytes.Buffer

	if err := format.Node(&out, fset, af); err != nil {
		log.Fatal(err)
	}

	src, err := imports.Process(filename, out.Bytes(), nil)

	if err != nil {
		log.Fatal(err)
	}

	src = regexp.MustCompile(`(?m)^func `).ReplaceAll(src, []byte("\n// Auto-generated\nfunc "))
	src, err = format.Source(src)

	if err != nil {
		log.Fatal(err)
	}

	dst, err := os.Create(strings.ReplaceAll(filename, ".go", "_without_ctx.go"))

	if err != nil {
		log.Fatal(err)
	}

	defer dst.Close()
	fmt.Fprintf(dst, "// Code generated from %s using genzfunc.go; DO NOT EDIT.\n", filename)
	fmt.Fprintln(dst, string(bytes.TrimSpace(src)))
}
