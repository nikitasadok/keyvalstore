package database

import (
	"keyvaluestore/datastructs/memtable"
	"os"
)

type Store struct {
	memTable *memtable.MemTable
	ssTables []*os.File

}
