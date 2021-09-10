// DO NOT EDIT: This file is autogenerated via the builtin command.

package debug

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
					Column: 13,
					Line:   14,
				},
				File:   "debug.flux",
				Source: "package debug\n\n\n// pass will pass any incoming tables directly next to the following transformation.\n// It is best used to interrupt any planner rules that rely on a specific ordering.\nbuiltin pass : (<-tables: [A]) => [A] where A: Record\n\n// slurp will read the incoming tables and concatenate buffers with the same group key\n// into a single table. This is useful for testing the performance impact of multiple\n// buffers versus a single buffer.\nbuiltin slurp : (<-tables: [A]) => [A] where A: Record\n\n// sink will discard all data that comes into it.\nbuiltin sink",
				Start: ast.Position{
					Column: 1,
					Line:   1,
				},
			},
		},
		Body: []ast.Statement{&ast.BuiltinStatement{
			BaseNode: ast.BaseNode{
				Comments: []ast.Comment{ast.Comment{Text: "// pass will pass any incoming tables directly next to the following transformation.\n"}, ast.Comment{Text: "// It is best used to interrupt any planner rules that rely on a specific ordering.\n"}},
				Errors:   nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 13,
						Line:   6,
					},
					File:   "debug.flux",
					Source: "builtin pass",
					Start: ast.Position{
						Column: 1,
						Line:   6,
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
							Column: 13,
							Line:   6,
						},
						File:   "debug.flux",
						Source: "pass",
						Start: ast.Position{
							Column: 9,
							Line:   6,
						},
					},
				},
				Name: "pass",
			},
			Ty: ast.TypeExpression{
				BaseNode: ast.BaseNode{
					Comments: nil,
					Errors:   nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 54,
							Line:   6,
						},
						File:   "debug.flux",
						Source: "(<-tables: [A]) => [A] where A: Record",
						Start: ast.Position{
							Column: 16,
							Line:   6,
						},
					},
				},
				Constraints: []*ast.TypeConstraint{&ast.TypeConstraint{
					BaseNode: ast.BaseNode{
						Comments: nil,
						Errors:   nil,
						Loc: &ast.SourceLocation{
							End: ast.Position{
								Column: 54,
								Line:   6,
							},
							File:   "debug.flux",
							Source: "A: Record",
							Start: ast.Position{
								Column: 45,
								Line:   6,
							},
						},
					},
					Kinds: []*ast.Identifier{&ast.Identifier{
						BaseNode: ast.BaseNode{
							Comments: nil,
							Errors:   nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 54,
									Line:   6,
								},
								File:   "debug.flux",
								Source: "Record",
								Start: ast.Position{
									Column: 48,
									Line:   6,
								},
							},
						},
						Name: "Record",
					}},
					Tvar: &ast.Identifier{
						BaseNode: ast.BaseNode{
							Comments: nil,
							Errors:   nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 46,
									Line:   6,
								},
								File:   "debug.flux",
								Source: "A",
								Start: ast.Position{
									Column: 45,
									Line:   6,
								},
							},
						},
						Name: "A",
					},
				}},
				Ty: &ast.FunctionType{
					BaseNode: ast.BaseNode{
						Comments: nil,
						Errors:   nil,
						Loc: &ast.SourceLocation{
							End: ast.Position{
								Column: 38,
								Line:   6,
							},
							File:   "debug.flux",
							Source: "(<-tables: [A]) => [A]",
							Start: ast.Position{
								Column: 16,
								Line:   6,
							},
						},
					},
					Parameters: []*ast.ParameterType{&ast.ParameterType{
						BaseNode: ast.BaseNode{
							Comments: nil,
							Errors:   nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 30,
									Line:   6,
								},
								File:   "debug.flux",
								Source: "<-tables: [A]",
								Start: ast.Position{
									Column: 17,
									Line:   6,
								},
							},
						},
						Kind: "Pipe",
						Name: &ast.Identifier{
							BaseNode: ast.BaseNode{
								Comments: nil,
								Errors:   nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 25,
										Line:   6,
									},
									File:   "debug.flux",
									Source: "tables",
									Start: ast.Position{
										Column: 19,
										Line:   6,
									},
								},
							},
							Name: "tables",
						},
						Ty: &ast.ArrayType{
							BaseNode: ast.BaseNode{
								Comments: nil,
								Errors:   nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 30,
										Line:   6,
									},
									File:   "debug.flux",
									Source: "[A]",
									Start: ast.Position{
										Column: 27,
										Line:   6,
									},
								},
							},
							ElementType: &ast.TvarType{
								BaseNode: ast.BaseNode{
									Comments: nil,
									Errors:   nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 29,
											Line:   6,
										},
										File:   "debug.flux",
										Source: "A",
										Start: ast.Position{
											Column: 28,
											Line:   6,
										},
									},
								},
								ID: &ast.Identifier{
									BaseNode: ast.BaseNode{
										Comments: nil,
										Errors:   nil,
										Loc: &ast.SourceLocation{
											End: ast.Position{
												Column: 29,
												Line:   6,
											},
											File:   "debug.flux",
											Source: "A",
											Start: ast.Position{
												Column: 28,
												Line:   6,
											},
										},
									},
									Name: "A",
								},
							},
						},
					}},
					Return: &ast.ArrayType{
						BaseNode: ast.BaseNode{
							Comments: nil,
							Errors:   nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 38,
									Line:   6,
								},
								File:   "debug.flux",
								Source: "[A]",
								Start: ast.Position{
									Column: 35,
									Line:   6,
								},
							},
						},
						ElementType: &ast.TvarType{
							BaseNode: ast.BaseNode{
								Comments: nil,
								Errors:   nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 37,
										Line:   6,
									},
									File:   "debug.flux",
									Source: "A",
									Start: ast.Position{
										Column: 36,
										Line:   6,
									},
								},
							},
							ID: &ast.Identifier{
								BaseNode: ast.BaseNode{
									Comments: nil,
									Errors:   nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 37,
											Line:   6,
										},
										File:   "debug.flux",
										Source: "A",
										Start: ast.Position{
											Column: 36,
											Line:   6,
										},
									},
								},
								Name: "A",
							},
						},
					},
				},
			},
		}, &ast.BuiltinStatement{
			BaseNode: ast.BaseNode{
				Comments: []ast.Comment{ast.Comment{Text: "// slurp will read the incoming tables and concatenate buffers with the same group key\n"}, ast.Comment{Text: "// into a single table. This is useful for testing the performance impact of multiple\n"}, ast.Comment{Text: "// buffers versus a single buffer.\n"}},
				Errors:   nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 14,
						Line:   11,
					},
					File:   "debug.flux",
					Source: "builtin slurp",
					Start: ast.Position{
						Column: 1,
						Line:   11,
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
							Column: 14,
							Line:   11,
						},
						File:   "debug.flux",
						Source: "slurp",
						Start: ast.Position{
							Column: 9,
							Line:   11,
						},
					},
				},
				Name: "slurp",
			},
			Ty: ast.TypeExpression{
				BaseNode: ast.BaseNode{
					Comments: nil,
					Errors:   nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 55,
							Line:   11,
						},
						File:   "debug.flux",
						Source: "(<-tables: [A]) => [A] where A: Record",
						Start: ast.Position{
							Column: 17,
							Line:   11,
						},
					},
				},
				Constraints: []*ast.TypeConstraint{&ast.TypeConstraint{
					BaseNode: ast.BaseNode{
						Comments: nil,
						Errors:   nil,
						Loc: &ast.SourceLocation{
							End: ast.Position{
								Column: 55,
								Line:   11,
							},
							File:   "debug.flux",
							Source: "A: Record",
							Start: ast.Position{
								Column: 46,
								Line:   11,
							},
						},
					},
					Kinds: []*ast.Identifier{&ast.Identifier{
						BaseNode: ast.BaseNode{
							Comments: nil,
							Errors:   nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 55,
									Line:   11,
								},
								File:   "debug.flux",
								Source: "Record",
								Start: ast.Position{
									Column: 49,
									Line:   11,
								},
							},
						},
						Name: "Record",
					}},
					Tvar: &ast.Identifier{
						BaseNode: ast.BaseNode{
							Comments: nil,
							Errors:   nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 47,
									Line:   11,
								},
								File:   "debug.flux",
								Source: "A",
								Start: ast.Position{
									Column: 46,
									Line:   11,
								},
							},
						},
						Name: "A",
					},
				}},
				Ty: &ast.FunctionType{
					BaseNode: ast.BaseNode{
						Comments: nil,
						Errors:   nil,
						Loc: &ast.SourceLocation{
							End: ast.Position{
								Column: 39,
								Line:   11,
							},
							File:   "debug.flux",
							Source: "(<-tables: [A]) => [A]",
							Start: ast.Position{
								Column: 17,
								Line:   11,
							},
						},
					},
					Parameters: []*ast.ParameterType{&ast.ParameterType{
						BaseNode: ast.BaseNode{
							Comments: nil,
							Errors:   nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 31,
									Line:   11,
								},
								File:   "debug.flux",
								Source: "<-tables: [A]",
								Start: ast.Position{
									Column: 18,
									Line:   11,
								},
							},
						},
						Kind: "Pipe",
						Name: &ast.Identifier{
							BaseNode: ast.BaseNode{
								Comments: nil,
								Errors:   nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 26,
										Line:   11,
									},
									File:   "debug.flux",
									Source: "tables",
									Start: ast.Position{
										Column: 20,
										Line:   11,
									},
								},
							},
							Name: "tables",
						},
						Ty: &ast.ArrayType{
							BaseNode: ast.BaseNode{
								Comments: nil,
								Errors:   nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 31,
										Line:   11,
									},
									File:   "debug.flux",
									Source: "[A]",
									Start: ast.Position{
										Column: 28,
										Line:   11,
									},
								},
							},
							ElementType: &ast.TvarType{
								BaseNode: ast.BaseNode{
									Comments: nil,
									Errors:   nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 30,
											Line:   11,
										},
										File:   "debug.flux",
										Source: "A",
										Start: ast.Position{
											Column: 29,
											Line:   11,
										},
									},
								},
								ID: &ast.Identifier{
									BaseNode: ast.BaseNode{
										Comments: nil,
										Errors:   nil,
										Loc: &ast.SourceLocation{
											End: ast.Position{
												Column: 30,
												Line:   11,
											},
											File:   "debug.flux",
											Source: "A",
											Start: ast.Position{
												Column: 29,
												Line:   11,
											},
										},
									},
									Name: "A",
								},
							},
						},
					}},
					Return: &ast.ArrayType{
						BaseNode: ast.BaseNode{
							Comments: nil,
							Errors:   nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 39,
									Line:   11,
								},
								File:   "debug.flux",
								Source: "[A]",
								Start: ast.Position{
									Column: 36,
									Line:   11,
								},
							},
						},
						ElementType: &ast.TvarType{
							BaseNode: ast.BaseNode{
								Comments: nil,
								Errors:   nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 38,
										Line:   11,
									},
									File:   "debug.flux",
									Source: "A",
									Start: ast.Position{
										Column: 37,
										Line:   11,
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
											Line:   11,
										},
										File:   "debug.flux",
										Source: "A",
										Start: ast.Position{
											Column: 37,
											Line:   11,
										},
									},
								},
								Name: "A",
							},
						},
					},
				},
			},
		}, &ast.BuiltinStatement{
			BaseNode: ast.BaseNode{
				Comments: []ast.Comment{ast.Comment{Text: "// sink will discard all data that comes into it.\n"}},
				Errors:   nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 13,
						Line:   14,
					},
					File:   "debug.flux",
					Source: "builtin sink",
					Start: ast.Position{
						Column: 1,
						Line:   14,
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
							Column: 13,
							Line:   14,
						},
						File:   "debug.flux",
						Source: "sink",
						Start: ast.Position{
							Column: 9,
							Line:   14,
						},
					},
				},
				Name: "sink",
			},
			Ty: ast.TypeExpression{
				BaseNode: ast.BaseNode{
					Comments: nil,
					Errors:   nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 54,
							Line:   14,
						},
						File:   "debug.flux",
						Source: "(<-tables: [A]) => [A] where A: Record",
						Start: ast.Position{
							Column: 16,
							Line:   14,
						},
					},
				},
				Constraints: []*ast.TypeConstraint{&ast.TypeConstraint{
					BaseNode: ast.BaseNode{
						Comments: nil,
						Errors:   nil,
						Loc: &ast.SourceLocation{
							End: ast.Position{
								Column: 54,
								Line:   14,
							},
							File:   "debug.flux",
							Source: "A: Record",
							Start: ast.Position{
								Column: 45,
								Line:   14,
							},
						},
					},
					Kinds: []*ast.Identifier{&ast.Identifier{
						BaseNode: ast.BaseNode{
							Comments: nil,
							Errors:   nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 54,
									Line:   14,
								},
								File:   "debug.flux",
								Source: "Record",
								Start: ast.Position{
									Column: 48,
									Line:   14,
								},
							},
						},
						Name: "Record",
					}},
					Tvar: &ast.Identifier{
						BaseNode: ast.BaseNode{
							Comments: nil,
							Errors:   nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 46,
									Line:   14,
								},
								File:   "debug.flux",
								Source: "A",
								Start: ast.Position{
									Column: 45,
									Line:   14,
								},
							},
						},
						Name: "A",
					},
				}},
				Ty: &ast.FunctionType{
					BaseNode: ast.BaseNode{
						Comments: nil,
						Errors:   nil,
						Loc: &ast.SourceLocation{
							End: ast.Position{
								Column: 38,
								Line:   14,
							},
							File:   "debug.flux",
							Source: "(<-tables: [A]) => [A]",
							Start: ast.Position{
								Column: 16,
								Line:   14,
							},
						},
					},
					Parameters: []*ast.ParameterType{&ast.ParameterType{
						BaseNode: ast.BaseNode{
							Comments: nil,
							Errors:   nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 30,
									Line:   14,
								},
								File:   "debug.flux",
								Source: "<-tables: [A]",
								Start: ast.Position{
									Column: 17,
									Line:   14,
								},
							},
						},
						Kind: "Pipe",
						Name: &ast.Identifier{
							BaseNode: ast.BaseNode{
								Comments: nil,
								Errors:   nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 25,
										Line:   14,
									},
									File:   "debug.flux",
									Source: "tables",
									Start: ast.Position{
										Column: 19,
										Line:   14,
									},
								},
							},
							Name: "tables",
						},
						Ty: &ast.ArrayType{
							BaseNode: ast.BaseNode{
								Comments: nil,
								Errors:   nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 30,
										Line:   14,
									},
									File:   "debug.flux",
									Source: "[A]",
									Start: ast.Position{
										Column: 27,
										Line:   14,
									},
								},
							},
							ElementType: &ast.TvarType{
								BaseNode: ast.BaseNode{
									Comments: nil,
									Errors:   nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 29,
											Line:   14,
										},
										File:   "debug.flux",
										Source: "A",
										Start: ast.Position{
											Column: 28,
											Line:   14,
										},
									},
								},
								ID: &ast.Identifier{
									BaseNode: ast.BaseNode{
										Comments: nil,
										Errors:   nil,
										Loc: &ast.SourceLocation{
											End: ast.Position{
												Column: 29,
												Line:   14,
											},
											File:   "debug.flux",
											Source: "A",
											Start: ast.Position{
												Column: 28,
												Line:   14,
											},
										},
									},
									Name: "A",
								},
							},
						},
					}},
					Return: &ast.ArrayType{
						BaseNode: ast.BaseNode{
							Comments: nil,
							Errors:   nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 38,
									Line:   14,
								},
								File:   "debug.flux",
								Source: "[A]",
								Start: ast.Position{
									Column: 35,
									Line:   14,
								},
							},
						},
						ElementType: &ast.TvarType{
							BaseNode: ast.BaseNode{
								Comments: nil,
								Errors:   nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 37,
										Line:   14,
									},
									File:   "debug.flux",
									Source: "A",
									Start: ast.Position{
										Column: 36,
										Line:   14,
									},
								},
							},
							ID: &ast.Identifier{
								BaseNode: ast.BaseNode{
									Comments: nil,
									Errors:   nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 37,
											Line:   14,
										},
										File:   "debug.flux",
										Source: "A",
										Start: ast.Position{
											Column: 36,
											Line:   14,
										},
									},
								},
								Name: "A",
							},
						},
					},
				},
			},
		}},
		Eof:      nil,
		Imports:  nil,
		Metadata: "parser-type=rust",
		Name:     "debug.flux",
		Package: &ast.PackageClause{
			BaseNode: ast.BaseNode{
				Comments: nil,
				Errors:   nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 14,
						Line:   1,
					},
					File:   "debug.flux",
					Source: "package debug",
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
							Column: 14,
							Line:   1,
						},
						File:   "debug.flux",
						Source: "debug",
						Start: ast.Position{
							Column: 9,
							Line:   1,
						},
					},
				},
				Name: "debug",
			},
		},
	}},
	Package: "debug",
	Path:    "internal/debug",
}