// Generated by tmpl
// https://github.com/benbjohnson/tmpl
//
// DO NOT EDIT!
// Source: array_cursor_iterator.gen.go.tmpl

package tsm1

import (
	"context"
	"github.com/influxdata/influxdb/v2/influxql/query"
	"github.com/influxdata/influxdb/v2/models"
	"github.com/influxdata/influxdb/v2/tsdb"
)

// buildFloatArrayCursor creates an array cursor for a float field.
func (q *arrayCursorIterator) buildFloatArrayCursor(ctx context.Context, name []byte, tags models.Tags, field string, opt query.IteratorOptions) tsdb.FloatArrayCursor {
	key := q.seriesFieldKeyBytes(name, tags, field)
	cacheValues := q.e.Cache.Values(key)
	keyCursor := q.e.KeyCursor(ctx, key, opt.SeekTime(), opt.Ascending)
	if opt.Ascending {
		q.asc.Float = newFloatArrayAscendingCursor()
		//if q.asc.Float == nil {
		//	q.asc.Float = newFloatArrayAscendingCursor()
		//}
		q.asc.Float.reset(opt.SeekTime(), opt.StopTime(), cacheValues, keyCursor)
		return q.asc.Float
	} else {
		if q.desc.Float == nil {
			q.desc.Float = newFloatArrayDescendingCursor()
		}
		q.desc.Float.reset(opt.SeekTime(), opt.StopTime(), cacheValues, keyCursor)
		return q.desc.Float
	}
}

// buildIntegerArrayCursor creates an array cursor for a integer field.
func (q *arrayCursorIterator) buildIntegerArrayCursor(ctx context.Context, name []byte, tags models.Tags, field string, opt query.IteratorOptions) tsdb.IntegerArrayCursor {
	key := q.seriesFieldKeyBytes(name, tags, field)
	cacheValues := q.e.Cache.Values(key)
	keyCursor := q.e.KeyCursor(ctx, key, opt.SeekTime(), opt.Ascending)
	if opt.Ascending {
		if q.asc.Integer == nil {
			q.asc.Integer = newIntegerArrayAscendingCursor()
		}
		q.asc.Integer.reset(opt.SeekTime(), opt.StopTime(), cacheValues, keyCursor)
		return q.asc.Integer
	} else {
		if q.desc.Integer == nil {
			q.desc.Integer = newIntegerArrayDescendingCursor()
		}
		q.desc.Integer.reset(opt.SeekTime(), opt.StopTime(), cacheValues, keyCursor)
		return q.desc.Integer
	}
}

// buildUnsignedArrayCursor creates an array cursor for a unsigned field.
func (q *arrayCursorIterator) buildUnsignedArrayCursor(ctx context.Context, name []byte, tags models.Tags, field string, opt query.IteratorOptions) tsdb.UnsignedArrayCursor {
	key := q.seriesFieldKeyBytes(name, tags, field)
	cacheValues := q.e.Cache.Values(key)
	keyCursor := q.e.KeyCursor(ctx, key, opt.SeekTime(), opt.Ascending)
	if opt.Ascending {
		if q.asc.Unsigned == nil {
			q.asc.Unsigned = newUnsignedArrayAscendingCursor()
		}
		q.asc.Unsigned.reset(opt.SeekTime(), opt.StopTime(), cacheValues, keyCursor)
		return q.asc.Unsigned
	} else {
		if q.desc.Unsigned == nil {
			q.desc.Unsigned = newUnsignedArrayDescendingCursor()
		}
		q.desc.Unsigned.reset(opt.SeekTime(), opt.StopTime(), cacheValues, keyCursor)
		return q.desc.Unsigned
	}
}

// buildStringArrayCursor creates an array cursor for a string field.
func (q *arrayCursorIterator) buildStringArrayCursor(ctx context.Context, name []byte, tags models.Tags, field string, opt query.IteratorOptions) tsdb.StringArrayCursor {
	key := q.seriesFieldKeyBytes(name, tags, field)
	cacheValues := q.e.Cache.Values(key)
	keyCursor := q.e.KeyCursor(ctx, key, opt.SeekTime(), opt.Ascending)
	if opt.Ascending {
		if q.asc.String == nil {
			q.asc.String = newStringArrayAscendingCursor()
		}
		q.asc.String.reset(opt.SeekTime(), opt.StopTime(), cacheValues, keyCursor)
		return q.asc.String
	} else {
		if q.desc.String == nil {
			q.desc.String = newStringArrayDescendingCursor()
		}
		q.desc.String.reset(opt.SeekTime(), opt.StopTime(), cacheValues, keyCursor)
		return q.desc.String
	}
}

// buildBooleanArrayCursor creates an array cursor for a boolean field.
func (q *arrayCursorIterator) buildBooleanArrayCursor(ctx context.Context, name []byte, tags models.Tags, field string, opt query.IteratorOptions) tsdb.BooleanArrayCursor {
	key := q.seriesFieldKeyBytes(name, tags, field)
	cacheValues := q.e.Cache.Values(key)
	keyCursor := q.e.KeyCursor(ctx, key, opt.SeekTime(), opt.Ascending)
	if opt.Ascending {
		if q.asc.Boolean == nil {
			q.asc.Boolean = newBooleanArrayAscendingCursor()
		}
		q.asc.Boolean.reset(opt.SeekTime(), opt.StopTime(), cacheValues, keyCursor)
		return q.asc.Boolean
	} else {
		if q.desc.Boolean == nil {
			q.desc.Boolean = newBooleanArrayDescendingCursor()
		}
		q.desc.Boolean.reset(opt.SeekTime(), opt.StopTime(), cacheValues, keyCursor)
		return q.desc.Boolean
	}
}
