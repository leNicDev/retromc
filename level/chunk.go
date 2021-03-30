package level

import (
	"bytes"
	"compress/zlib"
)

const (
	CHUNK_SIZE_X = 16
	CHUNK_SIZE_Y = 128
	CHUNK_SIZE_Z = 16
)

type Chunk struct {
	X     int32
	Y     int16
	Z     int32
	SizeX byte
	SizeY byte
	SizeZ byte
	Data  []byte
}

// compress chunk data using zlib
func (c *Chunk) CompressData() []byte {
	var buf bytes.Buffer
	writer := zlib.NewWriter(&buf)
	writer.Write(c.Data)
	writer.Close()
	return buf.Bytes()
}

// fill the chunk with air blocks
func (c *Chunk) fillAir() {
	// calculate amount of blocks in a chunk
	blocksAmount := CHUNK_SIZE_X * CHUNK_SIZE_Y * CHUNK_SIZE_Z

	blockTypes := make([]byte, blocksAmount)
	blockMetadata := make([]byte, divEvenOrRoundUp(blocksAmount, 2))
	blockLight := make([]byte, divEvenOrRoundUp(blocksAmount, 2))
	blockSkyLight := make([]byte, divEvenOrRoundUp(blocksAmount, 2))

	metaIndex := 0
	for blockIndex := 0; blockIndex < blocksAmount; blockIndex++ {
		// create new air block
		block := NewAirBlock()

		// set block type id
		blockTypes[blockIndex] = block.TypeId // set block type

		if blockIndex%2 == 0 { // lower bits
			blockMetadata[metaIndex/2] = block.Metadata
			blockLight[metaIndex/2] = block.Light
			blockSkyLight[metaIndex/2] = block.SkyLight
		} else { // upper bits
			blockMetadata[metaIndex] = mergeNibbles(blockMetadata[metaIndex], block.Metadata)
			blockLight[metaIndex] = mergeNibbles(blockLight[metaIndex], block.Light)
			blockSkyLight[metaIndex] = mergeNibbles(blockSkyLight[metaIndex], block.SkyLight)
			metaIndex++
		}
	}

	// concat data
	data := blockTypes
	data = append(data, blockMetadata...)
	data = append(data, blockLight...)
	data = append(data, blockSkyLight...)

	// update chunk data
	c.Data = data
}

func NewChunk() Chunk {
	chunk := Chunk{
		X:     0,
		Y:     0,
		Z:     0,
		SizeX: CHUNK_SIZE_X,
		SizeY: CHUNK_SIZE_Y,
		SizeZ: CHUNK_SIZE_Z,
	}
	chunk.fillAir()
	return chunk
}
