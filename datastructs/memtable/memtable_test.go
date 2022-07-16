package memtable

import (
	"github.com/stretchr/testify/assert"
	"keyvaluestore/datastructs/avltree"
	"os"
	"testing"
)

func TestMemTable_serializeEntry(t *testing.T) {
	type fields struct {
		tree     *avltree.Tree
		entryLog *os.File
	}
	type args struct {
		key string
		val []byte
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name: "one letter key and val",
			args: args{
				key: "k",
				val: []byte("v"),
			},
			want: 4,
		},

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &MemTable{
				tree:     tt.fields.tree,
				entryLog: tt.fields.entryLog,
			}
			got := m.serializeEntry(tt.args.key, tt.args.val)
			assert.Equal(t, tt.want, len(got))
		})
	}
}