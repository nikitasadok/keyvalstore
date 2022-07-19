package memtable

import (
	"bytes"
	"encoding/binary"
	"errors"
	"io"
	"os"
)

type MemTable struct {
	tree     *Tree
	entryLog *os.File
}

func New(path string) (*MemTable, error) {
	f, err := os.OpenFile(path, os.O_WRONLY | os.O_APPEND | os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}

	return &MemTable{
		tree:     nil,
		entryLog: f,
	}, nil
}

func (m *MemTable) Insert(key string, val []byte) error {
	err := m.writeToEntryLog(key, val)
	if err != nil {
		return errors.New("cannot insert to entry log")
	}

	m.tree = insertRec(m.tree, key, val)
	// m.tree.Insert(key, val)

	return nil
}

func (m *MemTable) Search(key string) ([]byte, bool) {
	val := m.tree.Search(key)

	return val, val != nil
}

func (m *MemTable) Delete(key string) error {
	val := m.tree.Search(key)
	if val == nil {
		return errors.New("key doesn't exist")
	}

	err := m.writeToEntryLog(key, []byte("tombstone"))
	if err != nil {
		return errors.New("cannot insert to entry log")
	}

	m.tree.Insert(key, nil)
	return nil
}

func (m *MemTable) writeToEntryLog(key string, val []byte) error {
	_, err := m.entryLog.Write(m.serializeEntry(key, val))
	return err
}

func (m *MemTable) serializeEntry(key string, val []byte) []byte {
	keyLen := len(key)
	valLen := len(val)

	varIntKeyLen := uvarintlen(uint64(keyLen))
	varIntValLen := uvarintlen(uint64(valLen))
	buf := make([]byte, varIntValLen + varIntKeyLen + keyLen + valLen)

	offset := binary.PutUvarint(buf, uint64(keyLen))
	for i := 0; i < keyLen; i++ {
		buf[offset + i] = key[i]
	}
	offset += keyLen
	offset += binary.PutUvarint(buf[offset:], uint64(valLen))
	for i := 0; i < valLen; i++ {
		buf[offset + i] = val[i]
	}

	return buf
}

func (m *MemTable) loadFromFile(path string) error{
	file, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	buf := bytes.NewBuffer(file)
	for {
		keyLen, err := binary.ReadUvarint(buf)
		if err != nil {
			break
		}

		keyBytes := make([]byte, keyLen)
		for i := uint64(0); i < keyLen; i++ {
			keyBytes[i], err = buf.ReadByte()
			if err != nil {
				break
			}
		}

		valLen, err := binary.ReadUvarint(buf)
		if err != nil {
			break
		}

		valBytes := make([]byte ,valLen)
		for i := uint64(0); i < valLen; i++ {
			valBytes[i], err = buf.ReadByte()
			if err != nil {
				break
			}
		}

		m.tree.Insert(string(keyBytes), valBytes)
		if err != nil {
			break
		}
	}

	if err == io.EOF {
		err = nil
	}

	return err
}

func uvarintlen(x uint64) int {
	i := 0
	for x >= 0x80 {
		x >>= 7
		i++
	}

	return i + 1
}
