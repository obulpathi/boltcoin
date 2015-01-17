package ar

import (
	"math/big"

	"github.com/obulpathi/boltcoin/trie"
)

type Block interface {
	Trie() *trie.Trie
	Diff() *big.Int
}
