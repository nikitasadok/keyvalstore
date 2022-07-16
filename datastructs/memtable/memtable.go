package memtable

import (
	"encoding/binary"
	"errors"
	"fmt"
	"keyvaluestore/datastructs/avltree"
	"os"
)

type MemTable struct {
	tree     *avltree.Tree
	entryLog *os.File
}

type record struct {
	key string
	val []byte
}

func (m *MemTable) Insert(key string, val []byte) error {
	err := m.writeToEntryLog(key, val)
	if err != nil {
		return errors.New("cannot insert to entry log")
	}
	m.tree = avltree.Insert(m.tree, key, val)

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

	m.tree = avltree.Insert(m.tree, key, nil)
	return nil
}

func (m *MemTable) writeToEntryLog(key string, val []byte) error {
	keyLen := len(key)
	valLen := len(val)
	buf := make([]byte, 0, 16+keyLen+valLen)
	binary.LittleEndian.PutUint64(buf[:8], uint64(keyLen))
	buf = append(buf, []byte(key)...)
    binary.LittleEndian.PutUint64(buf[len(buf):len(buf) + 8], uint64(valLen))
    buf = append(buf, val...)
	_, err := m.entryLog.Write(buf)

	return err
}

func (m *MemTable) serializeEntry(key string, val []byte) []byte {
	keyLen := len(key)
	valLen := len(val)

	varIntKeyLen := uvarintlen(uint64(keyLen))
	varIntValLen := uvarintlen(uint64(valLen))
	buf := make([]byte, varIntValLen + varIntKeyLen +keyLen+valLen)
	fmt.Println(len(buf))
	varKeySize := binary.PutUvarint(buf, uint64(keyLen))
	for i := 0; i < keyLen; i++ {
		buf[varKeySize+ i] = key[i]
	}

	fmt.Printf("%x\n", buf)
	varValSize := binary.PutUvarint(buf, uint64(valLen))
	for i := 0; i < valLen; i++ {
		buf[varValSize + i] = val[i]
	}
	fmt.Printf("%x\n", buf)
	return buf
}

func uvarintlen(x uint64) int {
	i := 0
	for x >= 0x80 {
		x >>= 7
		i++
	}

	return i + 1
}
