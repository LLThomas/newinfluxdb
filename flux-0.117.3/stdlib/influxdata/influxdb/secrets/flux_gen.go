// DO NOT EDIT: This file is autogenerated via the builtin command.

package secrets

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
					Column: 12,
					Line:   4,
				},
				File:   "secrets.flux",
				Source: "package secrets\n\n\nbuiltin get",
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
						Column: 12,
						Line:   4,
					},
					File:   "secrets.flux",
					Source: "builtin get",
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
							Column: 12,
							Line:   4,
						},
						File:   "secrets.flux",
						Source: "get",
						Start: ast.Position{
							Column: 9,
							Line:   4,
						},
					},
				},
				Name: "get",
			},
			Ty: ast.TypeExpression{
				BaseNode: ast.BaseNode{
					Comments: nil,
					Errors:   nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 38,
							Line:   4,
						},
						File:   "secrets.flux",
						Source: "(key: string) => string",
						Start: ast.Position{
							Column: 15,
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
								Column: 38,
								Line:   4,
							},
							File:   "secrets.flux",
							Source: "(key: string) => string",
							Start: ast.Position{
								Column: 15,
								Line:   4,
							},
						},
					},
					Parameters: []*ast.ParameterType{&ast.ParameterType{
						BaseNode: ast.BaseNode{
							Comments: nil,
							Errors:   nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 27,
									Line:   4,
								},
								File:   "secrets.flux",
								Source: "key: string",
								Start: ast.Position{
									Column: 16,
									Line:   4,
								},
							},
						},
						Kind: "Required",
						Name: &ast.Identifier{
							BaseNode: ast.BaseNode{
								Comments: nil,
								Errors:   nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 19,
										Line:   4,
									},
									File:   "secrets.flux",
									Source: "key",
									Start: ast.Position{
										Column: 16,
										Line:   4,
									},
								},
							},
							Name: "key",
						},
						Ty: &ast.NamedType{
							BaseNode: ast.BaseNode{
								Comments: nil,
								Errors:   nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 27,
										Line:   4,
									},
									File:   "secrets.flux",
									Source: "string",
									Start: ast.Position{
										Column: 21,
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
											Column: 27,
											Line:   4,
										},
										File:   "secrets.flux",
										Source: "string",
										Start: ast.Position{
											Column: 21,
											Line:   4,
										},
									},
								},
								Name: "string",
							},
						},
					}},
					Return: &ast.NamedType{
						BaseNode: ast.BaseNode{
							Comments: nil,
							Errors:   nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 38,
									Line:   4,
								},
								File:   "secrets.flux",
								Source: "string",
								Start: ast.Position{
									Column: 32,
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
										Column: 38,
										Line:   4,
									},
									File:   "secrets.flux",
									Source: "string",
									Start: ast.Position{
										Column: 32,
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
		Name:     "secrets.flux",
		Package: &ast.PackageClause{
			BaseNode: ast.BaseNode{
				Comments: nil,
				Errors:   nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 16,
						Line:   1,
					},
					File:   "secrets.flux",
					Source: "package secrets",
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
						File:   "secrets.flux",
						Source: "secrets",
						Start: ast.Position{
							Column: 9,
							Line:   1,
						},
					},
				},
				Name: "secrets",
			},
		},
	}},
	Package: "secrets",
	Path:    "influxdata/influxdb/secrets",
}