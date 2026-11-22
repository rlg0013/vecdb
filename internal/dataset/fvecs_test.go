package dataset

import (
	"bytes"
	"encoding/binary"
	"math"
	"testing"
)

func encodeVector(values []float32) []byte {
	var buf bytes.Buffer

	_ = binary.Write(
		&buf,
		binary.LittleEndian,
		int32(len(values)),
	)

	for _, value := range values {
		_ = binary.Write(
			&buf,
			binary.LittleEndian,
			value,
		)
	}

	return buf.Bytes()
}

func TestReadFvecsFrom(t *testing.T) {
	expected := [][]float32{
		{0.1, 0.2, 0.3},
		{1.0, 2.0, 3.0},
	}

	var data []byte

	for _, vector := range expected {
		data = append(data, encodeVector(vector)...)
	}

	got, err := ReadFvecsFrom(bytes.NewReader(data))

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(got) != len(expected) {
		t.Fatalf(
			"expected %d vectors, got %d",
			len(expected),
			len(got),
		)
	}

	for i := range expected {
		if len(got[i]) != len(expected[i]) {
			t.Fatalf("dimension mismatch at vector %d", i)
		}

		for j := range expected[i] {
			if math.Abs(
				float64(got[i][j]-expected[i][j]),
			) > 1e-6 {
				t.Errorf(
					"vector %d position %d: expected %f, got %f",
					i,
					j,
					expected[i][j],
					got[i][j],
				)
			}
		}
	}
}

func TestReadFvecsEmpty(t *testing.T) {
	got, err := ReadFvecsFrom(bytes.NewReader(nil))

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(got) != 0 {
		t.Fatalf("expected empty results, got %d vectors", len(got))
	}
}

func TestReadFvecsInavlidDimensions(t *testing.T) {
	var buf bytes.Buffer

	_ = binary.Write(
		&buf,
		binary.LittleEndian,
		int32(0),
	)

	_, err := ReadFvecsFrom(&buf)

	if err == nil {
		t.Fatalf("expcted error for zero dimension")
	}
}

func TestReadFvecsIncompleteVector(t *testing.T) {
	var buf bytes.Buffer

	_ = binary.Write(
		&buf,
		binary.LittleEndian,
		int32(3),
	)

	_ = binary.Write(
		&buf,
		binary.LittleEndian,
		float32(1.0),
	)

	_, err := ReadFvecsFrom(&buf)

	if err == nil {
		t.Fatal("expected error for incomplete vector")
	}
}

func TestReadFvecsRejectNaN(t *testing.T) {
	data := encodeVector([]float32{
		float32(math.NaN()),
	})

	_, err := ReadFvecsFrom(bytes.NewReader(data))

	if err == nil {
		t.Fatalf("expected error for NaN")
	}
}
