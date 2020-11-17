package codenotary___merkletree___space

import (
	"crypto/sha256"
	"fmt"
	"strings"

	"github.com/codenotary/merkletree"
)

/*
 * This `localMemStore` is the same as codenotary's internal `memStore`,
 * the code is just copied here for study and testing purposes.
 * A slight change on printing was added to get unique hashes (4 chars instead of two)
 */

func NewLocalMemStore() merkletree.Storer {
	return &LocalMemStore{
		data: make([][][sha256.Size]byte, 1),
	}
}

type LocalMemStore struct {
	data [][][sha256.Size]byte
}

func (m *LocalMemStore) Width() uint64 {
	return uint64(len(m.data[0]))
}

func (m *LocalMemStore) Set(layer uint8, index uint64, value [sha256.Size]byte) {

	for uint8(len(m.data)) <= layer {
		m.data = append(m.data, make([][sha256.Size]byte, 0, 256*256))
	}
	if uint64(len(m.data[layer])) == index {
		m.data[layer] = append(m.data[layer], value)
	} else {
		m.data[layer][index] = value
	}
}

func (m *LocalMemStore) Get(layer uint8, index uint64) *[sha256.Size]byte {

	if int(layer) >= len(m.data) || index >= uint64(len(m.data[layer])) {
		return nil
	}
	return &m.data[layer][index]
}

func (m *LocalMemStore) Print() {

	l := len(m.data)
	tab := ""
	for i := l - 1; i >= 0; i-- {
		fmt.Print(strings.Repeat("    ", (1<<i)-1))
		tab = strings.Repeat("    ", (1<<(i+1))-1)
		for _, v := range m.data[i] {
			fmt.Printf("%.4x%s", v[0:2], tab)
		}
		fmt.Println()
	}
}
