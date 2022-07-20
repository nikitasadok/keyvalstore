package memtable

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"os"
	"testing"
)

func TestMemTable_serializeEntry(t *testing.T) {
	type fields struct {
		tree     *Tree
		entryLog *os.File
	}
	type args struct {
		key string
		val []byte
	}
	tests := []struct {
		name   string
		fields fields
		args    args
		want []byte
	}{
		{
			name: "one letter key and val",
			args: args{
				key: "k",
				val: []byte("v"),
			},
			want: []byte{1, 'k', 1, 'v'},
		},
		{
			name: "key and val greater than 127 chars",
			args: args{
				key: generateRandomString(128),
				val: []byte(generateRandomString(128)),
			},
		},

		// TODO: Add test cases.
	}
	tests[1].want = serialization256KeyVal(tests[1].args.key, tests[1].args.val)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MemTable{
				tree:     tt.fields.tree,
				entryLog: tt.fields.entryLog,
			}
			got := m.serializeEntry(tt.args.key, tt.args.val)
			assert.Equal(t, tt.want, got)
		})
	}
}

func generateRandomString(size int) string {
	var letterRunes = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
		res := make([]byte, size)
		for j := 0; j < size; j++ {
			res[j] = letterRunes[rand.Intn(len(letterRunes))%size]
		}

	return string(res)
}

func serialization256KeyVal(key string, val []byte) []byte {
	serialized := []byte{}
	serialized = append(serialized, 80)
	serialized = append(serialized, 0)

	serialized = append(serialized, []byte(key)...)

	serialized = append(serialized, 1)
	serialized = append(serialized, 0)

	return append(serialized, val...)
}

func TestMemTable_deserialize(t *testing.T) {
	type fields struct {
		tree     *Tree
		entryLog *os.File
	}
	type args struct {
		buf *bytes.Buffer
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		wantTree *Tree
	}{
		{
			name: "simple valid key val",
			fields: fields{
				tree:     nil,
				entryLog: nil,
			},
			args: args{bytes.NewBuffer([]byte{1, 'k', 1, 'v'})},
			wantErr: false,
			wantTree: fillTree(nil, []treeEntry{
				{
					key: "k",
					val: []byte("v"),
				},
			}),
		},
		{
			name:    "a little bigger example",
			fields:  fields{},
			args:    args{bytes.NewBuffer([]byte{4, 'r', 'o', 'o', 't', 3, 'v', 'a', 'l', 4, 'l', 'k', 'e', 'y',
				4, 'l', 'v', 'a', 'l', 5, 'v', 'r', 'k','e','y', 4, 'r', 'v', 'a', 'l'})},
			wantErr: false,
			wantTree: fillTree(nil ,[]treeEntry{
				{
					key: "root",
					val: []byte("val"),
				},
				{
					key: "lkey",
					val: []byte("lval"),
				},
				{
					key: "vrkey",
					val: []byte("rval"),
				},
			}),
		},
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MemTable{
				tree:     tt.fields.tree,
				entryLog: tt.fields.entryLog,
			}
			err := m.deserialize(tt.args.buf)
			assert.Nil(t, err)
			assert.Equal(t, tt.wantTree, m.tree)
		})
	}
}