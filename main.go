package main

import (
	"keyvaluestore/datastructs/memtable"
	"log"
)

func main(){
	mt, err := memtable.New("testlogfile")
	if err != nil {
		log.Fatalln(err)
	}

	err = mt.Insert("someTestKey", []byte("someTestVal"))
	if err != nil {
		log.Fatalln(err)
	}
}
