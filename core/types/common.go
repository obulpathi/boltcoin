package types

import (
	"math/big"

	"github.com/obulpathi/boltcoin/state"
	"github.com/obulpathi/boltcoin/wire"
)

type BlockProcessor interface {
	Process(*Block) (*big.Int, state.Messages, error)
}

type Broadcaster interface {
	Broadcast(wire.MsgType, []interface{})
}
