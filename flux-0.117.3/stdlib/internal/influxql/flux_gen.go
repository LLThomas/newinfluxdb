// DO NOT EDIT: This file is autogenerated via the builtin command.

package influxql

import (
	ast "github.com/influxdata/flux/ast"
	parser "github.com/influxdata/flux/internal/parser"
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
					Column: 41,
					Line:   6,
				},
				File:   "influxql.flux",
				Source: "package influxql\n\n\nepoch = 1970-01-01T00:00:00Z\nminTime = 1677-09-21T00:12:43.145224194Z\nmaxTime = 2262-04-11T23:47:16.854775806Z",
				Start: ast.Position{
					Column: 1,
					Line:   1,
				},
			},
		},
		Body: []ast.Statement{&ast.VariableAssignment{
			BaseNode: ast.BaseNode{
				Comments: nil,
				Errors:   nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 29,
						Line:   4,
					},
					File:   "influxql.flux",
					Source: "epoch = 1970-01-01T00:00:00Z",
					Start: ast.Position{
						Column: 1,
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
							Column: 6,
							Line:   4,
						},
						File:   "influxql.flux",
						Source: "epoch",
						Start: ast.Position{
							Column: 1,
							Line:   4,
						},
					},
				},
				Name: "epoch",
			},
			Init: &ast.DateTimeLiteral{
				BaseNode: ast.BaseNode{
					Comments: nil,
					Errors:   nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 29,
							Line:   4,
						},
						File:   "influxql.flux",
						Source: "1970-01-01T00:00:00Z",
						Start: ast.Position{
							Column: 9,
							Line:   4,
						},
					},
				},
				Value: parser.MustParseTime("1970-01-01T00:00:00Z"),
			},
		}, &ast.VariableAssignment{
			BaseNode: ast.BaseNode{
				Comments: nil,
				Errors:   nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 41,
						Line:   5,
					},
					File:   "influxql.flux",
					Source: "minTime = 1677-09-21T00:12:43.145224194Z",
					Start: ast.Position{
						Column: 1,
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
							Column: 8,
							Line:   5,
						},
						File:   "influxql.flux",
						Source: "minTime",
						Start: ast.Position{
							Column: 1,
							Line:   5,
						},
					},
				},
				Name: "minTime",
			},
			Init: &ast.DateTimeLiteral{
				BaseNode: ast.BaseNode{
					Comments: nil,
					Errors:   nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 41,
							Line:   5,
						},
						File:   "influxql.flux",
						Source: "1677-09-21T00:12:43.145224194Z",
						Start: ast.Position{
							Column: 11,
							Line:   5,
						},
					},
				},
				Value: parser.MustParseTime("1677-09-21T00:12:43.145224194Z"),
			},
		}, &ast.VariableAssignment{
			BaseNode: ast.BaseNode{
				Comments: nil,
				Errors:   nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 41,
						Line:   6,
					},
					File:   "influxql.flux",
					Source: "maxTime = 2262-04-11T23:47:16.854775806Z",
					Start: ast.Position{
						Column: 1,
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
							Column: 8,
							Line:   6,
						},
						File:   "influxql.flux",
						Source: "maxTime",
						Start: ast.Position{
							Column: 1,
							Line:   6,
						},
					},
				},
				Name: "maxTime",
			},
			Init: &ast.DateTimeLiteral{
				BaseNode: ast.BaseNode{
					Comments: nil,
					Errors:   nil,
					Loc: &ast.SourceLocation{
						End: ast.Position{
							Column: 41,
							Line:   6,
						},
						File:   "influxql.flux",
						Source: "2262-04-11T23:47:16.854775806Z",
						Start: ast.Position{
							Column: 11,
							Line:   6,
						},
					},
				},
				Value: parser.MustParseTime("2262-04-11T23:47:16.854775806Z"),
			},
		}},
		Eof:      nil,
		Imports:  nil,
		Metadata: "parser-type=rust",
		Name:     "influxql.flux",
		Package: &ast.PackageClause{
			BaseNode: ast.BaseNode{
				Comments: nil,
				Errors:   nil,
				Loc: &ast.SourceLocation{
					End: ast.Position{
						Column: 17,
						Line:   1,
					},
					File:   "influxql.flux",
					Source: "package influxql",
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
							Column: 17,
							Line:   1,
						},
						File:   "influxql.flux",
						Source: "influxql",
						Start: ast.Position{
							Column: 9,
							Line:   1,
						},
					},
				},
				Name: "influxql",
			},
		},
	}},
	Package: "influxql",
	Path:    "internal/influxql",
}