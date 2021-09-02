// DO NOT EDIT: This file is autogenerated via the builtin command.

package runtime

import (
	ast "github.com/influxdata/flux/ast"
	runtime "github.com/influxdata/flux/runtime"
)

func init() {
	runtime.RegisterPackage(pkgAST)
}

var pkgAST = &ast.Package{
	BaseNode: ast.BaseNode{
		Comments: nil,
		Errors:   nil,
		Loc:      nil,
	},
	Files: []*ast.File{&ast.File{
		BaseNode: ast.BaseNode{
			Comments: nil,
			Errors:   nil,
			Loc: &ast.SourceLocation{
				End: ast.Position{
					Column: 16,
					Line:   4,
				},
				File:   "runtime.flux",
				Source: "package runtime\n\n\nbuiltin version",
				Start: ast.Position{
					Column: 1,
					Line:   1,
				},
			},
		},
		Body: []ast.Statement{&ast.BuiltinStatement{
			BaseNode: ast.BaseNode{
				Comments: nil,
				Errors:   nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 16,
						Line:   4,
					},
					File:   "runtime.flux",
					Source: "builtin version",
					Start: ast.Position{
						Column: 1,
						Line:   4,
					},
				},
			},
			Colon: nil,
			ID: &ast.Identifier{
				BaseNode: ast.BaseNode{
					Comments: nil,
					Errors:   nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 16,
							Line:   4,
						},
						File:   "runtime.flux",
						Source: "version",
						Start: ast.Position{
							Column: 9,
							Line:   4,
						},
					},
				},
				Name: "version",
			},
			Ty: ast.TypeExpression{
				BaseNode: ast.BaseNode{
					Comments: nil,
					Errors:   nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 31,
							Line:   4,
						},
						File:   "runtime.flux",
						Source: "() => string",
						Start: ast.Position{
							Column: 19,
							Line:   4,
						},
					},
				},
				Constraints: []*ast.TypeConstraint{},
				Ty: &ast.FunctionType{
					BaseNode: ast.BaseNode{
						Comments: nil,
						Errors:   nil,
						Loc: &ast.SourceLocation{
							End: ast.Position{
								Column: 31,
								Line:   4,
							},
							File:   "runtime.flux",
							Source: "() => string",
							Start: ast.Position{
								Column: 19,
								Line:   4,
							},
						},
					},
					Parameters: []*ast.ParameterType{},
					Return: &ast.NamedType{
						BaseNode: ast.BaseNode{
							Comments: nil,
							Errors:   nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 31,
									Line:   4,
								},
								File:   "runtime.flux",
								Source: "string",
								Start: ast.Position{
									Column: 25,
									Line:   4,
								},
							},
						},
						ID: &ast.Identifier{
							BaseNode: ast.BaseNode{
								Comments: nil,
								Errors:   nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 31,
										Line:   4,
									},
									File:   "runtime.flux",
									Source: "string",
									Start: ast.Position{
										Column: 25,
										Line:   4,
									},
								},
							},
							Name: "string",
						},
					},
				},
			},
		}},
		Eof:      nil,
		Imports:  nil,
		Metadata: "parser-type=rust",
		Name:     "runtime.flux",
		Package: &ast.PackageClause{
			BaseNode: ast.BaseNode{
				Comments: nil,
				Errors:   nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 16,
						Line:   1,
					},
					File:   "runtime.flux",
					Source: "package runtime",
					Start: ast.Position{
						Column: 1,
						Line:   1,
					},
				},
			},
			Name: &ast.Identifier{
				BaseNode: ast.BaseNode{
					Comments: nil,
					Errors:   nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 16,
							Line:   1,
						},
						File:   "runtime.flux",
						Source: "runtime",
						Start: ast.Position{
							Column: 9,
							Line:   1,
						},
					},
				},
				Name: "runtime",
			},
		},
	}},
	Package: "runtime",
	Path:    "runtime",
}
