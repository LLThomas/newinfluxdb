// Code generated by file_store.gen.go.tmpl. DO NOT EDIT.

package tsm1

import (
	"github.com/influxdata/influxdb/v2/tsdb"
)

// ReadFloatArrayBlock reads the next block as a set of float values.
func (c *KeyCursor) ReadFloatArrayBlock(values *tsdb.FloatArray) (*tsdb.FloatArray, error) {
LOOP:
	// No matching blocks to decode
	if len(c.current) == 0 {
		values.Timestamps = values.Timestamps[:0]
		values.Values = values.Values[:0]
		return values, nil
	}

	// First block is the oldest block containing the points we're searching for.
	first := c.current[0]
	err := first.r.ReadFloatArrayBlockAt(&first.entry, values)
	if err != nil {
		return nil, err
	}
	if c.col != nil {
		c.col.GetCounter(floatBlocksDecodedCounter).Add(1)
		c.col.GetCounter(floatBlocksSizeCounter).Add(int64(first.entry.Size))
	}

	// Remove values we already read
	values.Exclude(first.readMin, first.readMax)

	// Remove any tombstones
	tombstones := first.r.TombstoneRange(c.key)
	excludeTombstonesFloatArray(tombstones, values)
	// If there are no values in this first block (all tombstoned or previously read) and
	// we have more potential blocks too search.  Try again.
	if values.Len() == 0 && len(c.current) > 0 {
		c.current = c.current[1:]
		goto LOOP
	}

	// Only one block with this key and time range so return it
	if len(c.current) == 1 {
		if values.Len() > 0 {
			first.markRead(values.MinTime(), values.MaxTime())
		}
		return values, nil
	}

	// Use the current block time range as our overlapping window
	minT, maxT := first.readMin, first.readMax
	if values.Len() > 0 {
		minT, maxT = values.MinTime(), values.MaxTime()
	}
	if c.ascending {
		// Blocks are ordered by generation, we may have values in the past in later blocks, if so,
		// expand the window to include the min time range to ensure values are returned in ascending
		// order
		for i := 1; i < len(c.current); i++ {
			cur := c.current[i]
			if cur.entry.MinTime < minT && !cur.read() {
				minT = cur.entry.MinTime
			}
		}

		// Find first block that overlaps our window
		for i := 1; i < len(c.current); i++ {
			cur := c.current[i]
			if cur.entry.OverlapsTimeRange(minT, maxT) && !cur.read() {
				// Shrink our window so it's the intersection of the first overlapping block and the
				// first block.  We do this to minimize the region that overlaps and needs to
				// be merged.
				if cur.entry.MaxTime > maxT {
					maxT = cur.entry.MaxTime
				}
				values.Include(minT, maxT)
				break
			}
		}

		// Search the remaining blocks that overlap our window and append their values so we can
		// merge them.
		for i := 1; i < len(c.current); i++ {
			cur := c.current[i]
			// Skip this block if it doesn't contain points we looking for or they have already been read
			if !cur.entry.OverlapsTimeRange(minT, maxT) || cur.read() {
				cur.markRead(minT, maxT)
				continue
			}

			v := &tsdb.FloatArray{}
			err := cur.r.ReadFloatArrayBlockAt(&cur.entry, v)
			if err != nil {
				return nil, err
			}
			if c.col != nil {
				c.col.GetCounter(floatBlocksDecodedCounter).Add(1)
				c.col.GetCounter(floatBlocksSizeCounter).Add(int64(cur.entry.Size))
			}

			tombstones := cur.r.TombstoneRange(c.key)
			// Remove any tombstoned values
			excludeTombstonesFloatArray(tombstones, v)

			// Remove values we already read
			v.Exclude(cur.readMin, cur.readMax)

			if v.Len() > 0 {
				// Only use values in the overlapping window
				v.Include(minT, maxT)
				// Merge the remaining values with the existing
				values.Merge(v)
			}
			cur.markRead(minT, maxT)
		}

	} else {
		// Blocks are ordered by generation, we may have values in the past in later blocks, if so,
		// expand the window to include the max time range to ensure values are returned in descending
		// order
		for i := 1; i < len(c.current); i++ {
			cur := c.current[i]
			if cur.entry.MaxTime > maxT && !cur.read() {
				maxT = cur.entry.MaxTime
			}
		}

		// Find first block that overlaps our window
		for i := 1; i < len(c.current); i++ {
			cur := c.current[i]
			if cur.entry.OverlapsTimeRange(minT, maxT) && !cur.read() {
				// Shrink our window so it's the intersection of the first overlapping block and the
				// first block.  We do this to minimize the region that overlaps and needs to
				// be merged.
				if cur.entry.MinTime < minT {
					minT = cur.entry.MinTime
				}
				values.Include(minT, maxT)
				break
			}
		}

		// Search the remaining blocks that overlap our window and append their values so we can
		// merge them.
		for i := 1; i < len(c.current); i++ {
			cur := c.current[i]
			// Skip this block if it doesn't contain points we looking for or they have already been read
			if !cur.entry.OverlapsTimeRange(minT, maxT) || cur.read() {
				cur.markRead(minT, maxT)
				continue
			}

			v := &tsdb.FloatArray{}
			err := cur.r.ReadFloatArrayBlockAt(&cur.entry, v)
			if err != nil {
				return nil, err
			}
			if c.col != nil {
				c.col.GetCounter(floatBlocksDecodedCounter).Add(1)
				c.col.GetCounter(floatBlocksSizeCounter).Add(int64(cur.entry.Size))
			}
			tombstones := cur.r.TombstoneRange(c.key)
			// Remove any tombstoned values
			excludeTombstonesFloatArray(tombstones, v)

			// Remove values we already read
			v.Exclude(cur.readMin, cur.readMax)

			// If the block we decoded should have all of it's values included, mark it as read so we
			// don't use it again.
			if v.Len() > 0 {
				v.Include(minT, maxT)
				// Merge the remaining values with the existing
				v.Merge(values)
				*values = *v
			}
			cur.markRead(minT, maxT)
		}
	}

	first.markRead(minT, maxT)

	return values, err
}

func excludeTombstonesFloatArray(t []TimeRange, values *tsdb.FloatArray) {
	for i := range t {
		values.Exclude(t[i].Min, t[i].Max)
	}
}

// ReadIntegerArrayBlock reads the next block as a set of integer values.
func (c *KeyCursor) ReadIntegerArrayBlock(values *tsdb.IntegerArray) (*tsdb.IntegerArray, error) {
LOOP:
	// No matching blocks to decode
	if len(c.current) == 0 {
		values.Timestamps = values.Timestamps[:0]
		values.Values = values.Values[:0]
		return values, nil
	}

	// First block is the oldest block containing the points we're searching for.
	first := c.current[0]
	err := first.r.ReadIntegerArrayBlockAt(&first.entry, values)
	if err != nil {
		return nil, err
	}
	if c.col != nil {
		c.col.GetCounter(integerBlocksDecodedCounter).Add(1)
		c.col.GetCounter(integerBlocksSizeCounter).Add(int64(first.entry.Size))
	}

	// Remove values we already read
	values.Exclude(first.readMin, first.readMax)

	// Remove any tombstones
	tombstones := first.r.TombstoneRange(c.key)
	excludeTombstonesIntegerArray(tombstones, values)
	// If there are no values in this first block (all tombstoned or previously read) and
	// we have more potential blocks too search.  Try again.
	if values.Len() == 0 && len(c.current) > 0 {
		c.current = c.current[1:]
		goto LOOP
	}

	// Only one block with this key and time range so return it
	if len(c.current) == 1 {
		if values.Len() > 0 {
			first.markRead(values.MinTime(), values.MaxTime())
		}
		return values, nil
	}

	// Use the current block time range as our overlapping window
	minT, maxT := first.readMin, first.readMax
	if values.Len() > 0 {
		minT, maxT = values.MinTime(), values.MaxTime()
	}
	if c.ascending {
		// Blocks are ordered by generation, we may have values in the past in later blocks, if so,
		// expand the window to include the min time range to ensure values are returned in ascending
		// order
		for i := 1; i < len(c.current); i++ {
			cur := c.current[i]
			if cur.entry.MinTime < minT && !cur.read() {
				minT = cur.entry.MinTime
			}
		}

		// Find first block that overlaps our window
		for i := 1; i < len(c.current); i++ {
			cur := c.current[i]
			if cur.entry.OverlapsTimeRange(minT, maxT) && !cur.read() {
				// Shrink our window so it's the intersection of the first overlapping block and the
				// first block.  We do this to minimize the region that overlaps and needs to
				// be merged.
				if cur.entry.MaxTime > maxT {
					maxT = cur.entry.MaxTime
				}
				values.Include(minT, maxT)
				break
			}
		}

		// Search the remaining blocks that overlap our window and append their values so we can
		// merge them.
		for i := 1; i < len(c.current); i++ {
			cur := c.current[i]
			// Skip this block if it doesn't contain points we looking for or they have already been read
			if !cur.entry.OverlapsTimeRange(minT, maxT) || cur.read() {
				cur.markRead(minT, maxT)
				continue
			}

			v := &tsdb.IntegerArray{}
			err := cur.r.ReadIntegerArrayBlockAt(&cur.entry, v)
			if err != nil {
				return nil, err
			}
			if c.col != nil {
				c.col.GetCounter(integerBlocksDecodedCounter).Add(1)
				c.col.GetCounter(integerBlocksSizeCounter).Add(int64(cur.entry.Size))
			}

			tombstones := cur.r.TombstoneRange(c.key)
			// Remove any tombstoned values
			excludeTombstonesIntegerArray(tombstones, v)

			// Remove values we already read
			v.Exclude(cur.readMin, cur.readMax)

			if v.Len() > 0 {
				// Only use values in the overlapping window
				v.Include(minT, maxT)
				// Merge the remaining values with the existing
				values.Merge(v)
			}
			cur.markRead(minT, maxT)
		}

	} else {
		// Blocks are ordered by generation, we may have values in the past in later blocks, if so,
		// expand the window to include the max time range to ensure values are returned in descending
		// order
		for i := 1; i < len(c.current); i++ {
			cur := c.current[i]
			if cur.entry.MaxTime > maxT && !cur.read() {
				maxT = cur.entry.MaxTime
			}
		}

		// Find first block that overlaps our window
		for i := 1; i < len(c.current); i++ {
			cur := c.current[i]
			if cur.entry.OverlapsTimeRange(minT, maxT) && !cur.read() {
				// Shrink our window so it's the intersection of the first overlapping block and the
				// first block.  We do this to minimize the region that overlaps and needs to
				// be merged.
				if cur.entry.MinTime < minT {
					minT = cur.entry.MinTime
				}
				values.Include(minT, maxT)
				break
			}
		}

		// Search the remaining blocks that overlap our window and append their values so we can
		// merge them.
		for i := 1; i < len(c.current); i++ {
			cur := c.current[i]
			// Skip this block if it doesn't contain points we looking for or they have already been read
			if !cur.entry.OverlapsTimeRange(minT, maxT) || cur.read() {
				cur.markRead(minT, maxT)
				continue
			}

			v := &tsdb.IntegerArray{}
			err := cur.r.ReadIntegerArrayBlockAt(&cur.entry, v)
			if err != nil {
				return nil, err
			}
			if c.col != nil {
				c.col.GetCounter(integerBlocksDecodedCounter).Add(1)
				c.col.GetCounter(integerBlocksSizeCounter).Add(int64(cur.entry.Size))
			}
			tombstones := cur.r.TombstoneRange(c.key)
			// Remove any tombstoned values
			excludeTombstonesIntegerArray(tombstones, v)

			// Remove values we already read
			v.Exclude(cur.readMin, cur.readMax)

			// If the block we decoded should have all of it's values included, mark it as read so we
			// don't use it again.
			if v.Len() > 0 {
				v.Include(minT, maxT)
				// Merge the remaining values with the existing
				v.Merge(values)
				*values = *v
			}
			cur.markRead(minT, maxT)
		}
	}

	first.markRead(minT, maxT)

	return values, err
}

func excludeTombstonesIntegerArray(t []TimeRange, values *tsdb.IntegerArray) {
	for i := range t {
		values.Exclude(t[i].Min, t[i].Max)
	}
}

// ReadUnsignedArrayBlock reads the next block as a set of unsigned values.
func (c *KeyCursor) ReadUnsignedArrayBlock(values *tsdb.UnsignedArray) (*tsdb.UnsignedArray, error) {
LOOP:
	// No matching blocks to decode
	if len(c.current) == 0 {
		values.Timestamps = values.Timestamps[:0]
		values.Values = values.Values[:0]
		return values, nil
	}

	// First block is the oldest block containing the points we're searching for.
	first := c.current[0]
	err := first.r.ReadUnsignedArrayBlockAt(&first.entry, values)
	if err != nil {
		return nil, err
	}
	if c.col != nil {
		c.col.GetCounter(unsignedBlocksDecodedCounter).Add(1)
		c.col.GetCounter(unsignedBlocksSizeCounter).Add(int64(first.entry.Size))
	}

	// Remove values we already read
	values.Exclude(first.readMin, first.readMax)

	// Remove any tombstones
	tombstones := first.r.TombstoneRange(c.key)
	excludeTombstonesUnsignedArray(tombstones, values)
	// If there are no values in this first block (all tombstoned or previously read) and
	// we have more potential blocks too search.  Try again.
	if values.Len() == 0 && len(c.current) > 0 {
		c.current = c.current[1:]
		goto LOOP
	}

	// Only one block with this key and time range so return it
	if len(c.current) == 1 {
		if values.Len() > 0 {
			first.markRead(values.MinTime(), values.MaxTime())
		}
		return values, nil
	}

	// Use the current block time range as our overlapping window
	minT, maxT := first.readMin, first.readMax
	if values.Len() > 0 {
		minT, maxT = values.MinTime(), values.MaxTime()
	}
	if c.ascending {
		// Blocks are ordered by generation, we may have values in the past in later blocks, if so,
		// expand the window to include the min time range to ensure values are returned in ascending
		// order
		for i := 1; i < len(c.current); i++ {
			cur := c.current[i]
			if cur.entry.MinTime < minT && !cur.read() {
				minT = cur.entry.MinTime
			}
		}

		// Find first block that overlaps our window
		for i := 1; i < len(c.current); i++ {
			cur := c.current[i]
			if cur.entry.OverlapsTimeRange(minT, maxT) && !cur.read() {
				// Shrink our window so it's the intersection of the first overlapping block and the
				// first block.  We do this to minimize the region that overlaps and needs to
				// be merged.
				if cur.entry.MaxTime > maxT {
					maxT = cur.entry.MaxTime
				}
				values.Include(minT, maxT)
				break
			}
		}

		// Search the remaining blocks that overlap our window and append their values so we can
		// merge them.
		for i := 1; i < len(c.current); i++ {
			cur := c.current[i]
			// Skip this block if it doesn't contain points we looking for or they have already been read
			if !cur.entry.OverlapsTimeRange(minT, maxT) || cur.read() {
				cur.markRead(minT, maxT)
				continue
			}

			v := &tsdb.UnsignedArray{}
			err := cur.r.ReadUnsignedArrayBlockAt(&cur.entry, v)
			if err != nil {
				return nil, err
			}
			if c.col != nil {
				c.col.GetCounter(unsignedBlocksDecodedCounter).Add(1)
				c.col.GetCounter(unsignedBlocksSizeCounter).Add(int64(cur.entry.Size))
			}

			tombstones := cur.r.TombstoneRange(c.key)
			// Remove any tombstoned values
			excludeTombstonesUnsignedArray(tombstones, v)

			// Remove values we already read
			v.Exclude(cur.readMin, cur.readMax)

			if v.Len() > 0 {
				// Only use values in the overlapping window
				v.Include(minT, maxT)
				// Merge the remaining values with the existing
				values.Merge(v)
			}
			cur.markRead(minT, maxT)
		}

	} else {
		// Blocks are ordered by generation, we may have values in the past in later blocks, if so,
		// expand the window to include the max time range to ensure values are returned in descending
		// order
		for i := 1; i < len(c.current); i++ {
			cur := c.current[i]
			if cur.entry.MaxTime > maxT && !cur.read() {
				maxT = cur.entry.MaxTime
			}
		}

		// Find first block that overlaps our window
		for i := 1; i < len(c.current); i++ {
			cur := c.current[i]
			if cur.entry.OverlapsTimeRange(minT, maxT) && !cur.read() {
				// Shrink our window so it's the intersection of the first overlapping block and the
				// first block.  We do this to minimize the region that overlaps and needs to
				// be merged.
				if cur.entry.MinTime < minT {
					minT = cur.entry.MinTime
				}
				values.Include(minT, maxT)
				break
			}
		}

		// Search the remaining blocks that overlap our window and append their values so we can
		// merge them.
		for i := 1; i < len(c.current); i++ {
			cur := c.current[i]
			// Skip this block if it doesn't contain points we looking for or they have already been read
			if !cur.entry.OverlapsTimeRange(minT, maxT) || cur.read() {
				cur.markRead(minT, maxT)
				continue
			}

			v := &tsdb.UnsignedArray{}
			err := cur.r.ReadUnsignedArrayBlockAt(&cur.entry, v)
			if err != nil {
				return nil, err
			}
			if c.col != nil {
				c.col.GetCounter(unsignedBlocksDecodedCounter).Add(1)
				c.col.GetCounter(unsignedBlocksSizeCounter).Add(int64(cur.entry.Size))
			}
			tombstones := cur.r.TombstoneRange(c.key)
			// Remove any tombstoned values
			excludeTombstonesUnsignedArray(tombstones, v)

			// Remove values we already read
			v.Exclude(cur.readMin, cur.readMax)

			// If the block we decoded should have all of it's values included, mark it as read so we
			// don't use it again.
			if v.Len() > 0 {
				v.Include(minT, maxT)
				// Merge the remaining values with the existing
				v.Merge(values)
				*values = *v
			}
			cur.markRead(minT, maxT)
		}
	}

	first.markRead(minT, maxT)

	return values, err
}

func excludeTombstonesUnsignedArray(t []TimeRange, values *tsdb.UnsignedArray) {
	for i := range t {
		values.Exclude(t[i].Min, t[i].Max)
	}
}

// ReadStringArrayBlock reads the next block as a set of string values.
func (c *KeyCursor) ReadStringArrayBlock(values *tsdb.StringArray) (*tsdb.StringArray, error) {
LOOP:
	// No matching blocks to decode
	if len(c.current) == 0 {
		values.Timestamps = values.Timestamps[:0]
		values.Values = values.Values[:0]
		return values, nil
	}

	// First block is the oldest block containing the points we're searching for.
	first := c.current[0]
	err := first.r.ReadStringArrayBlockAt(&first.entry, values)
	if err != nil {
		return nil, err
	}
	if c.col != nil {
		c.col.GetCounter(stringBlocksDecodedCounter).Add(1)
		c.col.GetCounter(stringBlocksSizeCounter).Add(int64(first.entry.Size))
	}

	// Remove values we already read
	values.Exclude(first.readMin, first.readMax)

	// Remove any tombstones
	tombstones := first.r.TombstoneRange(c.key)
	excludeTombstonesStringArray(tombstones, values)
	// If there are no values in this first block (all tombstoned or previously read) and
	// we have more potential blocks too search.  Try again.
	if values.Len() == 0 && len(c.current) > 0 {
		c.current = c.current[1:]
		goto LOOP
	}

	// Only one block with this key and time range so return it
	if len(c.current) == 1 {
		if values.Len() > 0 {
			first.markRead(values.MinTime(), values.MaxTime())
		}
		return values, nil
	}

	// Use the current block time range as our overlapping window
	minT, maxT := first.readMin, first.readMax
	if values.Len() > 0 {
		minT, maxT = values.MinTime(), values.MaxTime()
	}
	if c.ascending {
		// Blocks are ordered by generation, we may have values in the past in later blocks, if so,
		// expand the window to include the min time range to ensure values are returned in ascending
		// order
		for i := 1; i < len(c.current); i++ {
			cur := c.current[i]
			if cur.entry.MinTime < minT && !cur.read() {
				minT = cur.entry.MinTime
			}
		}

		// Find first block that overlaps our window
		for i := 1; i < len(c.current); i++ {
			cur := c.current[i]
			if cur.entry.OverlapsTimeRange(minT, maxT) && !cur.read() {
				// Shrink our window so it's the intersection of the first overlapping block and the
				// first block.  We do this to minimize the region that overlaps and needs to
				// be merged.
				if cur.entry.MaxTime > maxT {
					maxT = cur.entry.MaxTime
				}
				values.Include(minT, maxT)
				break
			}
		}

		// Search the remaining blocks that overlap our window and append their values so we can
		// merge them.
		for i := 1; i < len(c.current); i++ {
			cur := c.current[i]
			// Skip this block if it doesn't contain points we looking for or they have already been read
			if !cur.entry.OverlapsTimeRange(minT, maxT) || cur.read() {
				cur.markRead(minT, maxT)
				continue
			}

			v := &tsdb.StringArray{}
			err := cur.r.ReadStringArrayBlockAt(&cur.entry, v)
			if err != nil {
				return nil, err
			}
			if c.col != nil {
				c.col.GetCounter(stringBlocksDecodedCounter).Add(1)
				c.col.GetCounter(stringBlocksSizeCounter).Add(int64(cur.entry.Size))
			}

			tombstones := cur.r.TombstoneRange(c.key)
			// Remove any tombstoned values
			excludeTombstonesStringArray(tombstones, v)

			// Remove values we already read
			v.Exclude(cur.readMin, cur.readMax)

			if v.Len() > 0 {
				// Only use values in the overlapping window
				v.Include(minT, maxT)
				// Merge the remaining values with the existing
				values.Merge(v)
			}
			cur.markRead(minT, maxT)
		}

	} else {
		// Blocks are ordered by generation, we may have values in the past in later blocks, if so,
		// expand the window to include the max time range to ensure values are returned in descending
		// order
		for i := 1; i < len(c.current); i++ {
			cur := c.current[i]
			if cur.entry.MaxTime > maxT && !cur.read() {
				maxT = cur.entry.MaxTime
			}
		}

		// Find first block that overlaps our window
		for i := 1; i < len(c.current); i++ {
			cur := c.current[i]
			if cur.entry.OverlapsTimeRange(minT, maxT) && !cur.read() {
				// Shrink our window so it's the intersection of the first overlapping block and the
				// first block.  We do this to minimize the region that overlaps and needs to
				// be merged.
				if cur.entry.MinTime < minT {
					minT = cur.entry.MinTime
				}
				values.Include(minT, maxT)
				break
			}
		}

		// Search the remaining blocks that overlap our window and append their values so we can
		// merge them.
		for i := 1; i < len(c.current); i++ {
			cur := c.current[i]
			// Skip this block if it doesn't contain points we looking for or they have already been read
			if !cur.entry.OverlapsTimeRange(minT, maxT) || cur.read() {
				cur.markRead(minT, maxT)
				continue
			}

			v := &tsdb.StringArray{}
			err := cur.r.ReadStringArrayBlockAt(&cur.entry, v)
			if err != nil {
				return nil, err
			}
			if c.col != nil {
				c.col.GetCounter(stringBlocksDecodedCounter).Add(1)
				c.col.GetCounter(stringBlocksSizeCounter).Add(int64(cur.entry.Size))
			}
			tombstones := cur.r.TombstoneRange(c.key)
			// Remove any tombstoned values
			excludeTombstonesStringArray(tombstones, v)

			// Remove values we already read
			v.Exclude(cur.readMin, cur.readMax)

			// If the block we decoded should have all of it's values included, mark it as read so we
			// don't use it again.
			if v.Len() > 0 {
				v.Include(minT, maxT)
				// Merge the remaining values with the existing
				v.Merge(values)
				*values = *v
			}
			cur.markRead(minT, maxT)
		}
	}

	first.markRead(minT, maxT)

	return values, err
}

func excludeTombstonesStringArray(t []TimeRange, values *tsdb.StringArray) {
	for i := range t {
		values.Exclude(t[i].Min, t[i].Max)
	}
}

// ReadBooleanArrayBlock reads the next block as a set of boolean values.
func (c *KeyCursor) ReadBooleanArrayBlock(values *tsdb.BooleanArray) (*tsdb.BooleanArray, error) {
LOOP:
	// No matching blocks to decode
	if len(c.current) == 0 {
		values.Timestamps = values.Timestamps[:0]
		values.Values = values.Values[:0]
		return values, nil
	}

	// First block is the oldest block containing the points we're searching for.
	first := c.current[0]
	err := first.r.ReadBooleanArrayBlockAt(&first.entry, values)
	if err != nil {
		return nil, err
	}
	if c.col != nil {
		c.col.GetCounter(booleanBlocksDecodedCounter).Add(1)
		c.col.GetCounter(booleanBlocksSizeCounter).Add(int64(first.entry.Size))
	}

	// Remove values we already read
	values.Exclude(first.readMin, first.readMax)

	// Remove any tombstones
	tombstones := first.r.TombstoneRange(c.key)
	excludeTombstonesBooleanArray(tombstones, values)
	// If there are no values in this first block (all tombstoned or previously read) and
	// we have more potential blocks too search.  Try again.
	if values.Len() == 0 && len(c.current) > 0 {
		c.current = c.current[1:]
		goto LOOP
	}

	// Only one block with this key and time range so return it
	if len(c.current) == 1 {
		if values.Len() > 0 {
			first.markRead(values.MinTime(), values.MaxTime())
		}
		return values, nil
	}

	// Use the current block time range as our overlapping window
	minT, maxT := first.readMin, first.readMax
	if values.Len() > 0 {
		minT, maxT = values.MinTime(), values.MaxTime()
	}
	if c.ascending {
		// Blocks are ordered by generation, we may have values in the past in later blocks, if so,
		// expand the window to include the min time range to ensure values are returned in ascending
		// order
		for i := 1; i < len(c.current); i++ {
			cur := c.current[i]
			if cur.entry.MinTime < minT && !cur.read() {
				minT = cur.entry.MinTime
			}
		}

		// Find first block that overlaps our window
		for i := 1; i < len(c.current); i++ {
			cur := c.current[i]
			if cur.entry.OverlapsTimeRange(minT, maxT) && !cur.read() {
				// Shrink our window so it's the intersection of the first overlapping block and the
				// first block.  We do this to minimize the region that overlaps and needs to
				// be merged.
				if cur.entry.MaxTime > maxT {
					maxT = cur.entry.MaxTime
				}
				values.Include(minT, maxT)
				break
			}
		}

		// Search the remaining blocks that overlap our window and append their values so we can
		// merge them.
		for i := 1; i < len(c.current); i++ {
			cur := c.current[i]
			// Skip this block if it doesn't contain points we looking for or they have already been read
			if !cur.entry.OverlapsTimeRange(minT, maxT) || cur.read() {
				cur.markRead(minT, maxT)
				continue
			}

			v := &tsdb.BooleanArray{}
			err := cur.r.ReadBooleanArrayBlockAt(&cur.entry, v)
			if err != nil {
				return nil, err
			}
			if c.col != nil {
				c.col.GetCounter(booleanBlocksDecodedCounter).Add(1)
				c.col.GetCounter(booleanBlocksSizeCounter).Add(int64(cur.entry.Size))
			}

			tombstones := cur.r.TombstoneRange(c.key)
			// Remove any tombstoned values
			excludeTombstonesBooleanArray(tombstones, v)

			// Remove values we already read
			v.Exclude(cur.readMin, cur.readMax)

			if v.Len() > 0 {
				// Only use values in the overlapping window
				v.Include(minT, maxT)
				// Merge the remaining values with the existing
				values.Merge(v)
			}
			cur.markRead(minT, maxT)
		}

	} else {
		// Blocks are ordered by generation, we may have values in the past in later blocks, if so,
		// expand the window to include the max time range to ensure values are returned in descending
		// order
		for i := 1; i < len(c.current); i++ {
			cur := c.current[i]
			if cur.entry.MaxTime > maxT && !cur.read() {
				maxT = cur.entry.MaxTime
			}
		}

		// Find first block that overlaps our window
		for i := 1; i < len(c.current); i++ {
			cur := c.current[i]
			if cur.entry.OverlapsTimeRange(minT, maxT) && !cur.read() {
				// Shrink our window so it's the intersection of the first overlapping block and the
				// first block.  We do this to minimize the region that overlaps and needs to
				// be merged.
				if cur.entry.MinTime < minT {
					minT = cur.entry.MinTime
				}
				values.Include(minT, maxT)
				break
			}
		}

		// Search the remaining blocks that overlap our window and append their values so we can
		// merge them.
		for i := 1; i < len(c.current); i++ {
			cur := c.current[i]
			// Skip this block if it doesn't contain points we looking for or they have already been read
			if !cur.entry.OverlapsTimeRange(minT, maxT) || cur.read() {
				cur.markRead(minT, maxT)
				continue
			}

			v := &tsdb.BooleanArray{}
			err := cur.r.ReadBooleanArrayBlockAt(&cur.entry, v)
			if err != nil {
				return nil, err
			}
			if c.col != nil {
				c.col.GetCounter(booleanBlocksDecodedCounter).Add(1)
				c.col.GetCounter(booleanBlocksSizeCounter).Add(int64(cur.entry.Size))
			}
			tombstones := cur.r.TombstoneRange(c.key)
			// Remove any tombstoned values
			excludeTombstonesBooleanArray(tombstones, v)

			// Remove values we already read
			v.Exclude(cur.readMin, cur.readMax)

			// If the block we decoded should have all of it's values included, mark it as read so we
			// don't use it again.
			if v.Len() > 0 {
				v.Include(minT, maxT)
				// Merge the remaining values with the existing
				v.Merge(values)
				*values = *v
			}
			cur.markRead(minT, maxT)
		}
	}

	first.markRead(minT, maxT)

	return values, err
}

func excludeTombstonesBooleanArray(t []TimeRange, values *tsdb.BooleanArray) {
	for i := range t {
		values.Exclude(t[i].Min, t[i].Max)
	}
}
