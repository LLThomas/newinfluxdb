// Generated by tmpl
// https://github.com/benbjohnson/tmpl
//
// DO NOT EDIT!
// Source: iterator.gen_test.go.tmpl

package arrowutil_test

import (
	"math/rand"
	"testing"

	"github.com/apache/arrow/go/arrow/array"
	"github.com/apache/arrow/go/arrow/memory"
	"github.com/google/go-cmp/cmp"
	"github.com/influxdata/flux/internal/arrowutil"
)

func TestIterateInt64s(t *testing.T) {
	arrs := make([]array.Interface, 0, 3)
	for i := 0; i < 3; i++ {
		b := arrowutil.NewInt64Builder(memory.DefaultAllocator)
		for j := 0; j < 100; j++ {
			if 0.05 > rand.Float64() {
				b.AppendNull()
				continue
			}
			v := generateInt64()
			b.Append(v)
		}
		arrs = append(arrs, b.NewArray())
	}

	itr := arrowutil.IterateInt64s(arrs)
	for i := 0; i < 300; i++ {
		if !itr.Next() {
			t.Fatalf("expected next value, but got false at index %d", i)
		}

		arr := arrs[i/100].(*array.Int64)
		if want, got := arr.IsValid(i%100), itr.IsValid(); !cmp.Equal(want, got) {
			t.Fatalf("unexpected valid value at index %d -want/+got:\n%s", i, cmp.Diff(want, got))
		} else if want && got {
			if want, got := arr.Value(i%100), itr.Value(); !cmp.Equal(want, got) {
				t.Fatalf("unexpected value at index %d -want/+got:\n%s", i, cmp.Diff(want, got))
			}
		}
		if want, got := arr.IsNull(i%100), itr.IsNull(); !cmp.Equal(want, got) {
			t.Fatalf("unexpected null value at index %d -want/+got:\n%s", i, cmp.Diff(want, got))
		}
	}
}

func TestIterateUint64s(t *testing.T) {
	arrs := make([]array.Interface, 0, 3)
	for i := 0; i < 3; i++ {
		b := arrowutil.NewUint64Builder(memory.DefaultAllocator)
		for j := 0; j < 100; j++ {
			if 0.05 > rand.Float64() {
				b.AppendNull()
				continue
			}
			v := generateUint64()
			b.Append(v)
		}
		arrs = append(arrs, b.NewArray())
	}

	itr := arrowutil.IterateUint64s(arrs)
	for i := 0; i < 300; i++ {
		if !itr.Next() {
			t.Fatalf("expected next value, but got false at index %d", i)
		}

		arr := arrs[i/100].(*array.Uint64)
		if want, got := arr.IsValid(i%100), itr.IsValid(); !cmp.Equal(want, got) {
			t.Fatalf("unexpected valid value at index %d -want/+got:\n%s", i, cmp.Diff(want, got))
		} else if want && got {
			if want, got := arr.Value(i%100), itr.Value(); !cmp.Equal(want, got) {
				t.Fatalf("unexpected value at index %d -want/+got:\n%s", i, cmp.Diff(want, got))
			}
		}
		if want, got := arr.IsNull(i%100), itr.IsNull(); !cmp.Equal(want, got) {
			t.Fatalf("unexpected null value at index %d -want/+got:\n%s", i, cmp.Diff(want, got))
		}
	}
}

func TestIterateFloat64s(t *testing.T) {
	arrs := make([]array.Interface, 0, 3)
	for i := 0; i < 3; i++ {
		b := arrowutil.NewFloat64Builder(memory.DefaultAllocator)
		for j := 0; j < 100; j++ {
			if 0.05 > rand.Float64() {
				b.AppendNull()
				continue
			}
			v := generateFloat64()
			b.Append(v)
		}
		arrs = append(arrs, b.NewArray())
	}

	itr := arrowutil.IterateFloat64s(arrs)
	for i := 0; i < 300; i++ {
		if !itr.Next() {
			t.Fatalf("expected next value, but got false at index %d", i)
		}

		arr := arrs[i/100].(*array.Float64)
		if want, got := arr.IsValid(i%100), itr.IsValid(); !cmp.Equal(want, got) {
			t.Fatalf("unexpected valid value at index %d -want/+got:\n%s", i, cmp.Diff(want, got))
		} else if want && got {
			if want, got := arr.Value(i%100), itr.Value(); !cmp.Equal(want, got) {
				t.Fatalf("unexpected value at index %d -want/+got:\n%s", i, cmp.Diff(want, got))
			}
		}
		if want, got := arr.IsNull(i%100), itr.IsNull(); !cmp.Equal(want, got) {
			t.Fatalf("unexpected null value at index %d -want/+got:\n%s", i, cmp.Diff(want, got))
		}
	}
}

func TestIterateBooleans(t *testing.T) {
	arrs := make([]array.Interface, 0, 3)
	for i := 0; i < 3; i++ {
		b := arrowutil.NewBooleanBuilder(memory.DefaultAllocator)
		for j := 0; j < 100; j++ {
			if 0.05 > rand.Float64() {
				b.AppendNull()
				continue
			}
			v := generateBoolean()
			b.Append(v)
		}
		arrs = append(arrs, b.NewArray())
	}

	itr := arrowutil.IterateBooleans(arrs)
	for i := 0; i < 300; i++ {
		if !itr.Next() {
			t.Fatalf("expected next value, but got false at index %d", i)
		}

		arr := arrs[i/100].(*array.Boolean)
		if want, got := arr.IsValid(i%100), itr.IsValid(); !cmp.Equal(want, got) {
			t.Fatalf("unexpected valid value at index %d -want/+got:\n%s", i, cmp.Diff(want, got))
		} else if want && got {
			if want, got := arr.Value(i%100), itr.Value(); !cmp.Equal(want, got) {
				t.Fatalf("unexpected value at index %d -want/+got:\n%s", i, cmp.Diff(want, got))
			}
		}
		if want, got := arr.IsNull(i%100), itr.IsNull(); !cmp.Equal(want, got) {
			t.Fatalf("unexpected null value at index %d -want/+got:\n%s", i, cmp.Diff(want, got))
		}
	}
}

func TestIterateStrings(t *testing.T) {
	arrs := make([]array.Interface, 0, 3)
	for i := 0; i < 3; i++ {
		b := arrowutil.NewStringBuilder(memory.DefaultAllocator)
		for j := 0; j < 100; j++ {
			if 0.05 > rand.Float64() {
				b.AppendNull()
				continue
			}
			v := generateString()
			b.AppendString(v)
		}
		arrs = append(arrs, b.NewArray())
	}

	itr := arrowutil.IterateStrings(arrs)
	for i := 0; i < 300; i++ {
		if !itr.Next() {
			t.Fatalf("expected next value, but got false at index %d", i)
		}

		arr := arrs[i/100].(*array.Binary)
		if want, got := arr.IsValid(i%100), itr.IsValid(); !cmp.Equal(want, got) {
			t.Fatalf("unexpected valid value at index %d -want/+got:\n%s", i, cmp.Diff(want, got))
		} else if want && got {
			if want, got := arr.ValueString(i%100), itr.ValueString(); !cmp.Equal(want, got) {
				t.Fatalf("unexpected value at index %d -want/+got:\n%s", i, cmp.Diff(want, got))
			}
		}
		if want, got := arr.IsNull(i%100), itr.IsNull(); !cmp.Equal(want, got) {
			t.Fatalf("unexpected null value at index %d -want/+got:\n%s", i, cmp.Diff(want, got))
		}
	}
}
