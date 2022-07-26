package sstable

import (
	"bufio"
	"encoding/binary"
	"io"
	"os"
)

type SSTable struct {
	offsets []int
	keyVals *bufio.Reader
}

// offsets can be stored in memory
// on table init load in memory
// do binary-search like and find needed key
// 1) get middle index
// 2) check key at offset
// 3) do binary search depending on key val
func (t *SSTable) Search(key []byte) []byte {


	// numRecords offset_0 offset_1 offset_2 ... offset_n
}

func readKey(r *bufio.Reader) ([]byte, error) {
	keyLen, err := binary.ReadUvarint(r)
	if err != nil {
		return nil, err
	}

	buf := make([]byte, keyLen)
	for i := uint64(0); i < keyLen; i++ {
		b, err := r.ReadByte()
		if err != nil {
			return nil, err
		}

		buf[i] = b
	}

	return buf, nil
}
