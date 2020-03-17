package downloader

import (
	"testing"
)

func TestChs(t *testing.T) {
	if "cccc" != chs(4, "c") {
		t.Errorf("They should be equal")
	}
}

/* THESE TEST NO LONGER APPLY TO THE CURRENT IMPLEMENTATION, THEY'LL BE REPLACED SOON
func TestNextChunkNumber(t *testing.T) {

	resumer := Resumer{}
	resumer.TotalElements = 100
	resumer.DoneElements = 0
	resumer.chunkSize = 100

	nextChunkNumber, skipNRows := resumer.nextChunkNumber()

	assert.Equal(t, int64(0), nextChunkNumber, "they should be equal")
	assert.Equal(t, int64(0), skipNRows, "they should be equal")

	///

	resumer.TotalElements = 100
	resumer.DoneElements = 0
	resumer.chunkSize = 10

	nextChunkNumber, skipNRows = resumer.nextChunkNumber()

	assert.Equal(t, int64(0), nextChunkNumber, "they should be equal")
	assert.Equal(t, int64(0), skipNRows, "they should be equal")

	///

	resumer.TotalElements = 100
	resumer.DoneElements = 10
	resumer.chunkSize = 10

	nextChunkNumber, skipNRows = resumer.nextChunkNumber()

	assert.Equal(t, int64(1), nextChunkNumber, "they should be equal")
	assert.Equal(t, int64(0), skipNRows, "they should be equal")

	///

	resumer.TotalElements = 100
	resumer.DoneElements = 9
	resumer.chunkSize = 10

	nextChunkNumber, skipNRows = resumer.nextChunkNumber()

	assert.Equal(t, int64(0), nextChunkNumber, "they should be equal")
	assert.Equal(t, int64(9), skipNRows, "they should be equal")

	///

	resumer.TotalElements = 100
	resumer.DoneElements = 11
	resumer.chunkSize = 10

	nextChunkNumber, skipNRows = resumer.nextChunkNumber()

	assert.Equal(t, int64(1), nextChunkNumber, "they should be equal")
	assert.Equal(t, int64(1), skipNRows, "they should be equal")

	///

	resumer.TotalElements = 99
	resumer.DoneElements = 98
	resumer.chunkSize = 10

	nextChunkNumber, skipNRows = resumer.nextChunkNumber()

	assert.Equal(t, int64(98), nextChunkNumber, "they should be equal")
	assert.Equal(t, int64(0), skipNRows, "they should be equal")

	///

	resumer.TotalElements = 99
	resumer.DoneElements = 99
	resumer.chunkSize = 10

	nextChunkNumber, skipNRows = resumer.nextChunkNumber()

	assert.Equal(t, int64(9), nextChunkNumber, "they should be equal")
	assert.Equal(t, int64(1), skipNRows, "they should be equal")
	assert.Equal(t, int64(1), resumer.chunkSize, "they should be equal")

}
*/
