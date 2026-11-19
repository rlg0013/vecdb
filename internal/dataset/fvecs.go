package dataset

import (
	"encoding/binary"
	"fmt"
	"io"
	"math"
	"os"
)

// ReadFvecs reads vectors from an fvecs binary file.
//
// Each vector is stored as:
//   - 4-byte little-endian int32 dimension
//   - dimension number of 4-byte float32 values
//
// The returned vectors are stored as []float32.
func ReadFvecs(path string) ([][]float32, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("open fvecs file: %w", err)
	}
	defer file.Close()

	return ReadFvecsFrom(file)
}

// ReadFvecsFrom reads vectors from an existing reader.
// This makes the parser easier to test without creating files.

func ReadFvecsFrom(r io.Reader) ([][]float32, error) {
	var vectors [][]float32

	var dimensionBuf [4]byte

	for {
		_, err := io.ReadFull(r, dimensionBuf[:])

		if err == io.EOF {
			break
		}

		if err == io.ErrUnexpectedEOF {
			return nil, fmt.Errorf("incomplete dimesnion prefix: %w", err)
		}

		if err != nil {
			return nil, fmt.Errorf("read dimension: %w", err)
		}

		dimension := int32(binary.LittleEndian.Uint32(dimensionBuf[:]))

		if dimension <= 0 {
			return nil, fmt.Errorf("invalid vector dimension: %w", dimension)
		}

		if dimension > 1_000_000 {
			return nil, fmt.Errorf("vector dimension too large: %w", dimension)
		}

		vector := make([]float32, int(dimension))

		for i := range vector {
			var valueBuf [4]byte

			_, err := io.ReadFull(r, valueBuf[:])
			if err != nil {
				return nil, fmt.Errorf(
					"read vector value %d: %w", i, err,
				)
			}

			bits := binary.LittleEndian.Uint32(valueBuf[:])
			value := math.Float32frombits(bits)

			if math.IsNaN(float64(value)) || math.IsInf(float64(value), 0) {
				return nil, fmt.Errorf("Invalid float32 value at position %d", i)
			}

			vector[i] = value
		}

		vectors = append(vectors, vector)
	}
	return vectors, nil
}
