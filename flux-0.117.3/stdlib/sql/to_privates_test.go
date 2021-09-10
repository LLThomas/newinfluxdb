package sql

import (
	"strconv"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	flux "github.com/influxdata/flux"
	"github.com/influxdata/flux/codes"
	"github.com/influxdata/flux/internal/errors"
)

// represents unsupported type
var unsupportedType flux.ColType = 666

// additional and seperate tests that can be run without needing functions to be Exported in sql, just to be testable
func TestCorrectBatchSize(t *testing.T) {
	// given the combination of row width and supplied batchSize argument from user, verify that it is modified as required
	userBatchSize := 1000
	rowWidth := 10
	correctedSize := correctBatchSize(userBatchSize, rowWidth)
	if !cmp.Equal(99, correctedSize) {
		t.Log(cmp.Diff(90, correctedSize))
		t.Fail()
	}

	// verify that the batchSoze is not lower than the width of a single row - if it ever is, we have a big problem
	userBatchSize = 1
	correctedSize = correctBatchSize(userBatchSize, rowWidth)
	if !cmp.Equal(10, correctedSize) {
		t.Log(cmp.Diff(10, correctedSize))
		t.Fail()
	}

	userBatchSize = -1
	correctedSize = correctBatchSize(userBatchSize, rowWidth)
	if !cmp.Equal(10, correctedSize) {
		t.Log(cmp.Diff(10, correctedSize))
		t.Fail()
	}
}

func TestTranslationDriverReturn(t *testing.T) {

	// verify invalid return error
	_, err := getTranslationFunc("bananas")
	if !cmp.Equal(errors.New(codes.Internal, "invalid driverName: bananas").Error(), err.Error()) {
		t.Log(cmp.Diff(errors.New(codes.Internal, "invalid driverName: bananas").Error(), err.Error()))
		t.Fail()
	}

	// verify that valid returns expected happiness for SQLITE
	_, err = getTranslationFunc("sqlite3")
	if !cmp.Equal(nil, err) {
		t.Log(cmp.Diff(nil, err))
		t.Fail()
	}

	// verify that valid returns expected happiness for Postgres
	_, err = getTranslationFunc("postgres")
	if !cmp.Equal(nil, err) {
		t.Log(cmp.Diff(nil, err))
		t.Fail()
	}

	// verify that valid returns expected happiness for MySQL
	_, err = getTranslationFunc("mysql")
	if !cmp.Equal(nil, err) {
		t.Log(cmp.Diff(nil, err))
		t.Fail()
	}

	// verify that valid returns expected happiness for Snowflake
	_, err = getTranslationFunc("snowflake")
	if !cmp.Equal(nil, err) {
		t.Log(cmp.Diff(nil, err))
		t.Fail()
	}

	// verify that valid returns expected happiness for Mssql
	_, err = getTranslationFunc("sqlserver")
	if !cmp.Equal(nil, err) {
		t.Log(cmp.Diff(nil, err))
		t.Fail()
	}

	// verify that valid returns expected happiness for BigQuery
	_, err = getTranslationFunc("bigquery")
	if !cmp.Equal(nil, err) {
		t.Log(cmp.Diff(nil, err))
		t.Fail()
	}

	// verify that valid returns expected error for AWS Athena (yes, error)
	expectedErr := errors.Newf(codes.Invalid, "writing is not supported for awsathena")
	_, err = getTranslationFunc("awsathena")
	if !cmp.Equal(expectedErr, err) {
		t.Log(cmp.Diff(expectedErr, err))
		t.Fail()
	}

	// verify that valid returns expected happiness for SAP HANA
	_, err = getTranslationFunc("hdb")
	if !cmp.Equal(nil, err) {
		t.Log(cmp.Diff(nil, err))
		t.Fail()
	}

}

func TestSqliteTranslation(t *testing.T) {
	sqliteTypeTranslations := map[string]flux.ColType{
		"FLOAT":    flux.TFloat,
		"INT":      flux.TInt,
		"TEXT":     flux.TString,
		"DATETIME": flux.TTime,
	}
	columnLabel := "apples"
	sqlT, err := getTranslationFunc("sqlite3")
	if !cmp.Equal(nil, err) {
		t.Log(cmp.Diff(nil, err))
		t.Fail()
	}

	for dbTypeString, fluxType := range sqliteTypeTranslations {
		v, err := sqlT()(fluxType, columnLabel)
		if !cmp.Equal(nil, err) {
			t.Log(cmp.Diff(nil, err))
			t.Fail()
		}
		if !cmp.Equal(columnLabel+" "+dbTypeString, v) {
			t.Log(cmp.Diff(columnLabel+" "+dbTypeString, v))
			t.Fail()
		}
	}

	// as SQLITE has NO BOOLEAN column type, we need to return an error rather than doing implicit conversions
	_, err = sqlT()(flux.TBool, columnLabel)
	if cmp.Equal(nil, err) {
		t.Log(cmp.Diff(nil, err))
		t.Fail()
	}
	if !cmp.Equal("SQLite does not support column type bool", err.Error()) {
		t.Log(cmp.Diff("SQLite does not support column type bool", err.Error()))
		t.Fail()
	}

}

func TestPostgresTranslation(t *testing.T) {
	postgresTypeTranslations := map[string]flux.ColType{
		"FLOAT":     flux.TFloat,
		"TEXT":      flux.TString,
		"BIGINT":    flux.TInt,
		"TIMESTAMP": flux.TTime,
		"BOOL":      flux.TBool,
	}

	columnLabel := "apples"
	// verify that valid returns expected happiness for postgres
	sqlT, err := getTranslationFunc("postgres")
	if !cmp.Equal(nil, err) {
		t.Log(cmp.Diff(nil, err))
		t.Fail()
	}

	for dbTypeString, fluxType := range postgresTypeTranslations {
		v, err := sqlT()(fluxType, columnLabel)
		if !cmp.Equal(nil, err) {
			t.Log(cmp.Diff(nil, err))
			t.Fail()
		}
		if !cmp.Equal(columnLabel+" "+dbTypeString, v) {
			t.Log(cmp.Diff(columnLabel+" "+dbTypeString, v))
			t.Fail()
		}
	}

	// test no match
	_, err = sqlT()(unsupportedType, columnLabel)
	if cmp.Equal(nil, err) {
		t.Log(cmp.Diff(nil, err))
		t.Fail()
	}
	if !cmp.Equal("PostgreSQL does not support column type unknown", err.Error()) {
		t.Log(cmp.Diff("PostgreSQL does not support column type unknown", err.Error()))
		t.Fail()
	}
}

func TestMysqlTranslation(t *testing.T) {
	mysqlTypeTranslations := map[string]flux.ColType{
		"FLOAT":       flux.TFloat,
		"BIGINT":      flux.TInt,
		"TEXT(16383)": flux.TString,
		"DATETIME":    flux.TTime,
		"BOOL":        flux.TBool,
	}

	columnLabel := "apples"
	// verify that valid returns expected happiness for mysql
	sqlT, err := getTranslationFunc("mysql")
	if !cmp.Equal(nil, err) {
		t.Log(cmp.Diff(nil, err))
		t.Fail()
	}

	for dbTypeString, fluxType := range mysqlTypeTranslations {
		v, err := sqlT()(fluxType, columnLabel)
		if !cmp.Equal(nil, err) {
			t.Log(cmp.Diff(nil, err))
			t.Fail()
		}
		if !cmp.Equal(columnLabel+" "+dbTypeString, v) {
			t.Log(cmp.Diff(columnLabel+" "+dbTypeString, v))
			t.Fail()
		}
	}

	// test no match
	_, err = sqlT()(unsupportedType, columnLabel)
	if cmp.Equal(nil, err) {
		t.Log(cmp.Diff(nil, err))
		t.Fail()
	}
	if !cmp.Equal("MySQL does not support column type unknown", err.Error()) {
		t.Log(cmp.Diff("MySQL does not support column type unknown", err.Error()))
		t.Fail()
	}
}

func TestSnowflakeTranslation(t *testing.T) {
	snowflakeTypeTranslations := map[string]flux.ColType{
		"FLOAT":         flux.TFloat,
		"NUMBER":        flux.TInt,
		"TEXT":          flux.TString,
		"TIMESTAMP_LTZ": flux.TTime,
		"BOOLEAN":       flux.TBool,
	}

	columnLabel := "apples"
	// verify that valid returns expected happiness for snowflake
	sqlT, err := getTranslationFunc("snowflake")
	if !cmp.Equal(nil, err) {
		t.Log(cmp.Diff(nil, err))
		t.Fail()
	}

	for dbTypeString, fluxType := range snowflakeTypeTranslations {
		v, err := sqlT()(fluxType, columnLabel)
		if !cmp.Equal(nil, err) {
			t.Log(cmp.Diff(nil, err))
			t.Fail()
		}
		if !cmp.Equal(columnLabel+" "+dbTypeString, v) {
			t.Log(cmp.Diff(columnLabel+" "+dbTypeString, v))
			t.Fail()
		}
	}

	// test no match
	_, err = sqlT()(unsupportedType, columnLabel)
	if cmp.Equal(nil, err) {
		t.Log(cmp.Diff(nil, err))
		t.Fail()
	}
	if !cmp.Equal("Snowflake does not support column type unknown", err.Error()) {
		t.Log(cmp.Diff("Snowflake does not support column type unknown", err.Error()))
		t.Fail()
	}
}

func TestMssqlTranslation(t *testing.T) {
	mssqlTypeTranslations := map[string]flux.ColType{
		"FLOAT":          flux.TFloat,
		"BIGINT":         flux.TInt,
		"VARCHAR(MAX)":   flux.TString,
		"DATETIMEOFFSET": flux.TTime,
		"BIT":            flux.TBool,
	}

	columnLabel := "apples"
	// verify that valid returns expected happiness for mssql
	sqlT, err := getTranslationFunc("sqlserver")
	if !cmp.Equal(nil, err) {
		t.Log(cmp.Diff(nil, err))
		t.Fail()
	}

	for dbTypeString, fluxType := range mssqlTypeTranslations {
		v, err := sqlT()(fluxType, columnLabel)
		if !cmp.Equal(nil, err) {
			t.Log(cmp.Diff(nil, err))
			t.Fail()
		}
		if !cmp.Equal(columnLabel+" "+dbTypeString, v) {
			t.Log(cmp.Diff(columnLabel+" "+dbTypeString, v))
			t.Fail()
		}
	}

	// test no match
	_, err = sqlT()(unsupportedType, columnLabel)
	if cmp.Equal(nil, err) {
		t.Log(cmp.Diff(nil, err))
		t.Fail()
	}
	if !cmp.Equal("SQLServer does not support column type unknown", err.Error()) {
		t.Log(cmp.Diff("SQLServer does not support column type unknown", err.Error()))
		t.Fail()
	}
}

func TestBigQueryTranslation(t *testing.T) {
	bigqueryTypeTranslations := map[string]flux.ColType{
		"FLOAT64":   flux.TFloat,
		"INT64":     flux.TInt,
		"STRING":    flux.TString,
		"TIMESTAMP": flux.TTime,
		"BOOL":      flux.TBool,
	}

	columnLabel := "apples"
	// verify that valid returns expected happiness for bigquery
	sqlT, err := getTranslationFunc("bigquery")
	if !cmp.Equal(nil, err) {
		t.Log(cmp.Diff(nil, err))
		t.Fail()
	}

	for dbTypeString, fluxType := range bigqueryTypeTranslations {
		v, err := sqlT()(fluxType, columnLabel)
		if !cmp.Equal(nil, err) {
			t.Log(cmp.Diff(nil, err))
			t.Fail()
		}
		if !cmp.Equal(columnLabel+" "+dbTypeString, v) {
			t.Log(cmp.Diff(columnLabel+" "+dbTypeString, v))
			t.Fail()
		}
	}

	// test no match
	_, err = sqlT()(unsupportedType, columnLabel)
	if cmp.Equal(nil, err) {
		t.Log(cmp.Diff(nil, err))
		t.Fail()
	}
	if !cmp.Equal("BigQuery does not support column type unknown", err.Error()) {
		t.Log(cmp.Diff("BigQuery does not support column type unknown", err.Error()))
		t.Fail()
	}
}

func TestHdbTranslation(t *testing.T) {
	hdbTypeTranslations := map[string]flux.ColType{
		"DOUBLE":         flux.TFloat,
		"BIGINT":         flux.TInt,
		"NVARCHAR(5000)": flux.TString,
		"TIMESTAMP":      flux.TTime,
		"BOOLEAN":        flux.TBool,
	}

	columnLabel := "apples"
	// verify that valid returns expected happiness for hdb
	sqlT, err := getTranslationFunc("hdb")
	if !cmp.Equal(nil, err) {
		t.Log(cmp.Diff(nil, err))
		t.Fail()
	}

	for dbTypeString, fluxType := range hdbTypeTranslations {
		v, err := sqlT()(fluxType, columnLabel)
		if !cmp.Equal(nil, err) {
			t.Log(cmp.Diff(nil, err))
			t.Fail()
		}
		if !cmp.Equal(strconv.Quote(strings.ToUpper(columnLabel))+" "+dbTypeString, v) {
			t.Log(cmp.Diff(strconv.Quote(strings.ToUpper(columnLabel))+" "+dbTypeString, v))
			t.Fail()
		}
	}

	// test no match
	var _unsupportedType flux.ColType = 666
	_, err = sqlT()(_unsupportedType, columnLabel)
	if cmp.Equal(nil, err) {
		t.Log(cmp.Diff(nil, err))
		t.Fail()
	}
	if !cmp.Equal("SAP HANA does not support column type unknown", err.Error()) {
		t.Log(cmp.Diff("SAP HANA does not support column type unknown", err.Error()))
		t.Fail()
	}
}