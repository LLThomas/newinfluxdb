// DO NOT EDIT: This file is autogenerated via the builtin command.

package mqtt

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
					Column: 11,
					Line:   4,
				},
				File:   "mqtt.flux",
				Source: "package mqtt\n\n\nbuiltin to",
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
						Column: 11,
						Line:   4,
					},
					File:   "mqtt.flux",
					Source: "builtin to",
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
							Column: 11,
							Line:   4,
						},
						File:   "mqtt.flux",
						Source: "to",
						Start: ast.Position{
							Column: 9,
							Line:   4,
						},
					},
				},
				Name: "to",
			},
			Ty: ast.TypeExpression{
				BaseNode: ast.BaseNode{
					Comments: nil,
					Errors:   nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 14,
							Line:   20,
						},
						File:   "mqtt.flux",
						Source: "(\n    <-tables: [A],\n    broker: string,\n    ?topic: string,\n    ?message: string,\n    ?qos: int,\n    ?clientid: string,\n    ?username: string,\n    ?password: string,\n    ?name: string,\n    ?timeout: duration,\n    ?timeColumn: string,\n    ?tagColumns: [string],\n    ?valueColumns: [string],\n) => [B] where\n    A: Record,\n    B: Record",
						Start: ast.Position{
							Column: 14,
							Line:   4,
						},
					},
				},
				Constraints: []*ast.TypeConstraint{&ast.TypeConstraint{
					BaseNode: ast.BaseNode{
						Comments: nil,
						Errors:   nil,
						Loc: &ast.SourceLocation{
							End: ast.Position{
								Column: 14,
								Line:   19,
							},
							File:   "mqtt.flux",
							Source: "A: Record",
							Start: ast.Position{
								Column: 5,
								Line:   19,
							},
						},
					},
					Kinds: []*ast.Identifier{&ast.Identifier{
						BaseNode: ast.BaseNode{
							Comments: nil,
							Errors:   nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 14,
									Line:   19,
								},
								File:   "mqtt.flux",
								Source: "Record",
								Start: ast.Position{
									Column: 8,
									Line:   19,
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
									Column: 6,
									Line:   19,
								},
								File:   "mqtt.flux",
								Source: "A",
								Start: ast.Position{
									Column: 5,
									Line:   19,
								},
							},
						},
						Name: "A",
					},
				}, &ast.TypeConstraint{
					BaseNode: ast.BaseNode{
						Comments: nil,
						Errors:   nil,
						Loc: &ast.SourceLocation{
							End: ast.Position{
								Column: 14,
								Line:   20,
							},
							File:   "mqtt.flux",
							Source: "B: Record",
							Start: ast.Position{
								Column: 5,
								Line:   20,
							},
						},
					},
					Kinds: []*ast.Identifier{&ast.Identifier{
						BaseNode: ast.BaseNode{
							Comments: nil,
							Errors:   nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 14,
									Line:   20,
								},
								File:   "mqtt.flux",
								Source: "Record",
								Start: ast.Position{
									Column: 8,
									Line:   20,
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
									Column: 6,
									Line:   20,
								},
								File:   "mqtt.flux",
								Source: "B",
								Start: ast.Position{
									Column: 5,
									Line:   20,
								},
							},
						},
						Name: "B",
					},
				}},
				Ty: &ast.FunctionType{
					BaseNode: ast.BaseNode{
						Comments: nil,
						Errors:   nil,
						Loc: &ast.SourceLocation{
							End: ast.Position{
								Column: 9,
								Line:   18,
							},
							File:   "mqtt.flux",
							Source: "(\n    <-tables: [A],\n    broker: string,\n    ?topic: string,\n    ?message: string,\n    ?qos: int,\n    ?clientid: string,\n    ?username: string,\n    ?password: string,\n    ?name: string,\n    ?timeout: duration,\n    ?timeColumn: string,\n    ?tagColumns: [string],\n    ?valueColumns: [string],\n) => [B]",
							Start: ast.Position{
								Column: 14,
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
									Column: 18,
									Line:   5,
								},
								File:   "mqtt.flux",
								Source: "<-tables: [A]",
								Start: ast.Position{
									Column: 5,
									Line:   5,
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
										Column: 13,
										Line:   5,
									},
									File:   "mqtt.flux",
									Source: "tables",
									Start: ast.Position{
										Column: 7,
										Line:   5,
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
										Column: 18,
										Line:   5,
									},
									File:   "mqtt.flux",
									Source: "[A]",
									Start: ast.Position{
										Column: 15,
										Line:   5,
									},
								},
							},
							ElementType: &ast.TvarType{
								BaseNode: ast.BaseNode{
									Comments: nil,
									Errors:   nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 17,
											Line:   5,
										},
										File:   "mqtt.flux",
										Source: "A",
										Start: ast.Position{
											Column: 16,
											Line:   5,
										},
									},
								},
								ID: &ast.Identifier{
									BaseNode: ast.BaseNode{
										Comments: nil,
										Errors:   nil,
										Loc: &ast.SourceLocation{
											End: ast.Position{
												Column: 17,
												Line:   5,
											},
											File:   "mqtt.flux",
											Source: "A",
											Start: ast.Position{
												Column: 16,
												Line:   5,
											},
										},
									},
									Name: "A",
								},
							},
						},
					}, &ast.ParameterType{
						BaseNode: ast.BaseNode{
							Comments: nil,
							Errors:   nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 19,
									Line:   6,
								},
								File:   "mqtt.flux",
								Source: "broker: string",
								Start: ast.Position{
									Column: 5,
									Line:   6,
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
										Column: 11,
										Line:   6,
									},
									File:   "mqtt.flux",
									Source: "broker",
									Start: ast.Position{
										Column: 5,
										Line:   6,
									},
								},
							},
							Name: "broker",
						},
						Ty: &ast.NamedType{
							BaseNode: ast.BaseNode{
								Comments: nil,
								Errors:   nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 19,
										Line:   6,
									},
									File:   "mqtt.flux",
									Source: "string",
									Start: ast.Position{
										Column: 13,
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
											Column: 19,
											Line:   6,
										},
										File:   "mqtt.flux",
										Source: "string",
										Start: ast.Position{
											Column: 13,
											Line:   6,
										},
									},
								},
								Name: "string",
							},
						},
					}, &ast.ParameterType{
						BaseNode: ast.BaseNode{
							Comments: nil,
							Errors:   nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 19,
									Line:   7,
								},
								File:   "mqtt.flux",
								Source: "?topic: string",
								Start: ast.Position{
									Column: 5,
									Line:   7,
								},
							},
						},
						Kind: "Optional",
						Name: &ast.Identifier{
							BaseNode: ast.BaseNode{
								Comments: nil,
								Errors:   nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 11,
										Line:   7,
									},
									File:   "mqtt.flux",
									Source: "topic",
									Start: ast.Position{
										Column: 6,
										Line:   7,
									},
								},
							},
							Name: "topic",
						},
						Ty: &ast.NamedType{
							BaseNode: ast.BaseNode{
								Comments: nil,
								Errors:   nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 19,
										Line:   7,
									},
									File:   "mqtt.flux",
									Source: "string",
									Start: ast.Position{
										Column: 13,
										Line:   7,
									},
								},
							},
							ID: &ast.Identifier{
								BaseNode: ast.BaseNode{
									Comments: nil,
									Errors:   nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 19,
											Line:   7,
										},
										File:   "mqtt.flux",
										Source: "string",
										Start: ast.Position{
											Column: 13,
											Line:   7,
										},
									},
								},
								Name: "string",
							},
						},
					}, &ast.ParameterType{
						BaseNode: ast.BaseNode{
							Comments: nil,
							Errors:   nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 21,
									Line:   8,
								},
								File:   "mqtt.flux",
								Source: "?message: string",
								Start: ast.Position{
									Column: 5,
									Line:   8,
								},
							},
						},
						Kind: "Optional",
						Name: &ast.Identifier{
							BaseNode: ast.BaseNode{
								Comments: nil,
								Errors:   nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 13,
										Line:   8,
									},
									File:   "mqtt.flux",
									Source: "message",
									Start: ast.Position{
										Column: 6,
										Line:   8,
									},
								},
							},
							Name: "message",
						},
						Ty: &ast.NamedType{
							BaseNode: ast.BaseNode{
								Comments: nil,
								Errors:   nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 21,
										Line:   8,
									},
									File:   "mqtt.flux",
									Source: "string",
									Start: ast.Position{
										Column: 15,
										Line:   8,
									},
								},
							},
							ID: &ast.Identifier{
								BaseNode: ast.BaseNode{
									Comments: nil,
									Errors:   nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 21,
											Line:   8,
										},
										File:   "mqtt.flux",
										Source: "string",
										Start: ast.Position{
											Column: 15,
											Line:   8,
										},
									},
								},
								Name: "string",
							},
						},
					}, &ast.ParameterType{
						BaseNode: ast.BaseNode{
							Comments: nil,
							Errors:   nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 14,
									Line:   9,
								},
								File:   "mqtt.flux",
								Source: "?qos: int",
								Start: ast.Position{
									Column: 5,
									Line:   9,
								},
							},
						},
						Kind: "Optional",
						Name: &ast.Identifier{
							BaseNode: ast.BaseNode{
								Comments: nil,
								Errors:   nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 9,
										Line:   9,
									},
									File:   "mqtt.flux",
									Source: "qos",
									Start: ast.Position{
										Column: 6,
										Line:   9,
									},
								},
							},
							Name: "qos",
						},
						Ty: &ast.NamedType{
							BaseNode: ast.BaseNode{
								Comments: nil,
								Errors:   nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 14,
										Line:   9,
									},
									File:   "mqtt.flux",
									Source: "int",
									Start: ast.Position{
										Column: 11,
										Line:   9,
									},
								},
							},
							ID: &ast.Identifier{
								BaseNode: ast.BaseNode{
									Comments: nil,
									Errors:   nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 14,
											Line:   9,
										},
										File:   "mqtt.flux",
										Source: "int",
										Start: ast.Position{
											Column: 11,
											Line:   9,
										},
									},
								},
								Name: "int",
							},
						},
					}, &ast.ParameterType{
						BaseNode: ast.BaseNode{
							Comments: nil,
							Errors:   nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 22,
									Line:   10,
								},
								File:   "mqtt.flux",
								Source: "?clientid: string",
								Start: ast.Position{
									Column: 5,
									Line:   10,
								},
							},
						},
						Kind: "Optional",
						Name: &ast.Identifier{
							BaseNode: ast.BaseNode{
								Comments: nil,
								Errors:   nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 14,
										Line:   10,
									},
									File:   "mqtt.flux",
									Source: "clientid",
									Start: ast.Position{
										Column: 6,
										Line:   10,
									},
								},
							},
							Name: "clientid",
						},
						Ty: &ast.NamedType{
							BaseNode: ast.BaseNode{
								Comments: nil,
								Errors:   nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 22,
										Line:   10,
									},
									File:   "mqtt.flux",
									Source: "string",
									Start: ast.Position{
										Column: 16,
										Line:   10,
									},
								},
							},
							ID: &ast.Identifier{
								BaseNode: ast.BaseNode{
									Comments: nil,
									Errors:   nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 22,
											Line:   10,
										},
										File:   "mqtt.flux",
										Source: "string",
										Start: ast.Position{
											Column: 16,
											Line:   10,
										},
									},
								},
								Name: "string",
							},
						},
					}, &ast.ParameterType{
						BaseNode: ast.BaseNode{
							Comments: nil,
							Errors:   nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 22,
									Line:   11,
								},
								File:   "mqtt.flux",
								Source: "?username: string",
								Start: ast.Position{
									Column: 5,
									Line:   11,
								},
							},
						},
						Kind: "Optional",
						Name: &ast.Identifier{
							BaseNode: ast.BaseNode{
								Comments: nil,
								Errors:   nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 14,
										Line:   11,
									},
									File:   "mqtt.flux",
									Source: "username",
									Start: ast.Position{
										Column: 6,
										Line:   11,
									},
								},
							},
							Name: "username",
						},
						Ty: &ast.NamedType{
							BaseNode: ast.BaseNode{
								Comments: nil,
								Errors:   nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 22,
										Line:   11,
									},
									File:   "mqtt.flux",
									Source: "string",
									Start: ast.Position{
										Column: 16,
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
											Column: 22,
											Line:   11,
										},
										File:   "mqtt.flux",
										Source: "string",
										Start: ast.Position{
											Column: 16,
											Line:   11,
										},
									},
								},
								Name: "string",
							},
						},
					}, &ast.ParameterType{
						BaseNode: ast.BaseNode{
							Comments: nil,
							Errors:   nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 22,
									Line:   12,
								},
								File:   "mqtt.flux",
								Source: "?password: string",
								Start: ast.Position{
									Column: 5,
									Line:   12,
								},
							},
						},
						Kind: "Optional",
						Name: &ast.Identifier{
							BaseNode: ast.BaseNode{
								Comments: nil,
								Errors:   nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 14,
										Line:   12,
									},
									File:   "mqtt.flux",
									Source: "password",
									Start: ast.Position{
										Column: 6,
										Line:   12,
									},
								},
							},
							Name: "password",
						},
						Ty: &ast.NamedType{
							BaseNode: ast.BaseNode{
								Comments: nil,
								Errors:   nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 22,
										Line:   12,
									},
									File:   "mqtt.flux",
									Source: "string",
									Start: ast.Position{
										Column: 16,
										Line:   12,
									},
								},
							},
							ID: &ast.Identifier{
								BaseNode: ast.BaseNode{
									Comments: nil,
									Errors:   nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 22,
											Line:   12,
										},
										File:   "mqtt.flux",
										Source: "string",
										Start: ast.Position{
											Column: 16,
											Line:   12,
										},
									},
								},
								Name: "string",
							},
						},
					}, &ast.ParameterType{
						BaseNode: ast.BaseNode{
							Comments: nil,
							Errors:   nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 18,
									Line:   13,
								},
								File:   "mqtt.flux",
								Source: "?name: string",
								Start: ast.Position{
									Column: 5,
									Line:   13,
								},
							},
						},
						Kind: "Optional",
						Name: &ast.Identifier{
							BaseNode: ast.BaseNode{
								Comments: nil,
								Errors:   nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 10,
										Line:   13,
									},
									File:   "mqtt.flux",
									Source: "name",
									Start: ast.Position{
										Column: 6,
										Line:   13,
									},
								},
							},
							Name: "name",
						},
						Ty: &ast.NamedType{
							BaseNode: ast.BaseNode{
								Comments: nil,
								Errors:   nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 18,
										Line:   13,
									},
									File:   "mqtt.flux",
									Source: "string",
									Start: ast.Position{
										Column: 12,
										Line:   13,
									},
								},
							},
							ID: &ast.Identifier{
								BaseNode: ast.BaseNode{
									Comments: nil,
									Errors:   nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 18,
											Line:   13,
										},
										File:   "mqtt.flux",
										Source: "string",
										Start: ast.Position{
											Column: 12,
											Line:   13,
										},
									},
								},
								Name: "string",
							},
						},
					}, &ast.ParameterType{
						BaseNode: ast.BaseNode{
							Comments: nil,
							Errors:   nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 23,
									Line:   14,
								},
								File:   "mqtt.flux",
								Source: "?timeout: duration",
								Start: ast.Position{
									Column: 5,
									Line:   14,
								},
							},
						},
						Kind: "Optional",
						Name: &ast.Identifier{
							BaseNode: ast.BaseNode{
								Comments: nil,
								Errors:   nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 13,
										Line:   14,
									},
									File:   "mqtt.flux",
									Source: "timeout",
									Start: ast.Position{
										Column: 6,
										Line:   14,
									},
								},
							},
							Name: "timeout",
						},
						Ty: &ast.NamedType{
							BaseNode: ast.BaseNode{
								Comments: nil,
								Errors:   nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 23,
										Line:   14,
									},
									File:   "mqtt.flux",
									Source: "duration",
									Start: ast.Position{
										Column: 15,
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
											Column: 23,
											Line:   14,
										},
										File:   "mqtt.flux",
										Source: "duration",
										Start: ast.Position{
											Column: 15,
											Line:   14,
										},
									},
								},
								Name: "duration",
							},
						},
					}, &ast.ParameterType{
						BaseNode: ast.BaseNode{
							Comments: nil,
							Errors:   nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 24,
									Line:   15,
								},
								File:   "mqtt.flux",
								Source: "?timeColumn: string",
								Start: ast.Position{
									Column: 5,
									Line:   15,
								},
							},
						},
						Kind: "Optional",
						Name: &ast.Identifier{
							BaseNode: ast.BaseNode{
								Comments: nil,
								Errors:   nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 16,
										Line:   15,
									},
									File:   "mqtt.flux",
									Source: "timeColumn",
									Start: ast.Position{
										Column: 6,
										Line:   15,
									},
								},
							},
							Name: "timeColumn",
						},
						Ty: &ast.NamedType{
							BaseNode: ast.BaseNode{
								Comments: nil,
								Errors:   nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 24,
										Line:   15,
									},
									File:   "mqtt.flux",
									Source: "string",
									Start: ast.Position{
										Column: 18,
										Line:   15,
									},
								},
							},
							ID: &ast.Identifier{
								BaseNode: ast.BaseNode{
									Comments: nil,
									Errors:   nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 24,
											Line:   15,
										},
										File:   "mqtt.flux",
										Source: "string",
										Start: ast.Position{
											Column: 18,
											Line:   15,
										},
									},
								},
								Name: "string",
							},
						},
					}, &ast.ParameterType{
						BaseNode: ast.BaseNode{
							Comments: nil,
							Errors:   nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 26,
									Line:   16,
								},
								File:   "mqtt.flux",
								Source: "?tagColumns: [string]",
								Start: ast.Position{
									Column: 5,
									Line:   16,
								},
							},
						},
						Kind: "Optional",
						Name: &ast.Identifier{
							BaseNode: ast.BaseNode{
								Comments: nil,
								Errors:   nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 16,
										Line:   16,
									},
									File:   "mqtt.flux",
									Source: "tagColumns",
									Start: ast.Position{
										Column: 6,
										Line:   16,
									},
								},
							},
							Name: "tagColumns",
						},
						Ty: &ast.ArrayType{
							BaseNode: ast.BaseNode{
								Comments: nil,
								Errors:   nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 26,
										Line:   16,
									},
									File:   "mqtt.flux",
									Source: "[string]",
									Start: ast.Position{
										Column: 18,
										Line:   16,
									},
								},
							},
							ElementType: &ast.NamedType{
								BaseNode: ast.BaseNode{
									Comments: nil,
									Errors:   nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 25,
											Line:   16,
										},
										File:   "mqtt.flux",
										Source: "string",
										Start: ast.Position{
											Column: 19,
											Line:   16,
										},
									},
								},
								ID: &ast.Identifier{
									BaseNode: ast.BaseNode{
										Comments: nil,
										Errors:   nil,
										Loc: &ast.SourceLocation{
											End: ast.Position{
												Column: 25,
												Line:   16,
											},
											File:   "mqtt.flux",
											Source: "string",
											Start: ast.Position{
												Column: 19,
												Line:   16,
											},
										},
									},
									Name: "string",
								},
							},
						},
					}, &ast.ParameterType{
						BaseNode: ast.BaseNode{
							Comments: nil,
							Errors:   nil,
							Loc: &ast.SourceLocation{
								End: ast.Position{
									Column: 28,
									Line:   17,
								},
								File:   "mqtt.flux",
								Source: "?valueColumns: [string]",
								Start: ast.Position{
									Column: 5,
									Line:   17,
								},
							},
						},
						Kind: "Optional",
						Name: &ast.Identifier{
							BaseNode: ast.BaseNode{
								Comments: nil,
								Errors:   nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 18,
										Line:   17,
									},
									File:   "mqtt.flux",
									Source: "valueColumns",
									Start: ast.Position{
										Column: 6,
										Line:   17,
									},
								},
							},
							Name: "valueColumns",
						},
						Ty: &ast.ArrayType{
							BaseNode: ast.BaseNode{
								Comments: nil,
								Errors:   nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 28,
										Line:   17,
									},
									File:   "mqtt.flux",
									Source: "[string]",
									Start: ast.Position{
										Column: 20,
										Line:   17,
									},
								},
							},
							ElementType: &ast.NamedType{
								BaseNode: ast.BaseNode{
									Comments: nil,
									Errors:   nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 27,
											Line:   17,
										},
										File:   "mqtt.flux",
										Source: "string",
										Start: ast.Position{
											Column: 21,
											Line:   17,
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
												Line:   17,
											},
											File:   "mqtt.flux",
											Source: "string",
											Start: ast.Position{
												Column: 21,
												Line:   17,
											},
										},
									},
									Name: "string",
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
									Column: 9,
									Line:   18,
								},
								File:   "mqtt.flux",
								Source: "[B]",
								Start: ast.Position{
									Column: 6,
									Line:   18,
								},
							},
						},
						ElementType: &ast.TvarType{
							BaseNode: ast.BaseNode{
								Comments: nil,
								Errors:   nil,
								Loc: &ast.SourceLocation{
									End: ast.Position{
										Column: 8,
										Line:   18,
									},
									File:   "mqtt.flux",
									Source: "B",
									Start: ast.Position{
										Column: 7,
										Line:   18,
									},
								},
							},
							ID: &ast.Identifier{
								BaseNode: ast.BaseNode{
									Comments: nil,
									Errors:   nil,
									Loc: &ast.SourceLocation{
										End: ast.Position{
											Column: 8,
											Line:   18,
										},
										File:   "mqtt.flux",
										Source: "B",
										Start: ast.Position{
											Column: 7,
											Line:   18,
										},
									},
								},
								Name: "B",
							},
						},
					},
				},
			},
		}},
		Eof:      nil,
		Imports:  nil,
		Metadata: "parser-type=rust",
		Name:     "mqtt.flux",
		Package: &ast.PackageClause{
			BaseNode: ast.BaseNode{
				Comments: nil,
				Errors:   nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 13,
						Line:   1,
					},
					File:   "mqtt.flux",
					Source: "package mqtt",
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
							Column: 13,
							Line:   1,
						},
						File:   "mqtt.flux",
						Source: "mqtt",
						Start: ast.Position{
							Column: 9,
							Line:   1,
						},
					},
				},
				Name: "mqtt",
			},
		},
	}},
	Package: "mqtt",
	Path:    "experimental/mqtt",
}