package core

import (
	"container/list"
	"fmt"

	"github.com/obulpathi/boltcoin/core/types"
	"github.com/obulpathi/boltcoin/crypto"
	"github.com/obulpathi/boltcoin/ethdb"
	"github.com/obulpathi/boltcoin/ethutil"
	"github.com/obulpathi/boltcoin/event"
	"github.com/obulpathi/boltcoin/wire"
)

// Implement our EthTest Manager
type TestManager struct {
	// stateManager *StateManager
	eventMux *event.TypeMux

	db         ethutil.Database
	txPool     *TxPool
	blockChain *ChainManager
	Blocks     []*types.Block
}

func (s *TestManager) IsListening() bool {
	return false
}

func (s *TestManager) IsMining() bool {
	return false
}

func (s *TestManager) PeerCount() int {
	return 0
}

func (s *TestManager) Peers() *list.List {
	return list.New()
}

func (s *TestManager) ChainManager() *ChainManager {
	return s.blockChain
}

func (tm *TestManager) TxPool() *TxPool {
	return tm.txPool
}

// func (tm *TestManager) StateManager() *StateManager {
// 	return tm.stateManager
// }

func (tm *TestManager) EventMux() *event.TypeMux {
	return tm.eventMux
}
func (tm *TestManager) Broadcast(msgType wire.MsgType, data []interface{}) {
	fmt.Println("Broadcast not implemented")
}

func (tm *TestManager) ClientIdentity() wire.ClientIdentity {
	return nil
}
func (tm *TestManager) KeyManager() *crypto.KeyManager {
	return nil
}

func (tm *TestManager) Db() ethutil.Database {
	return tm.db
}

func NewTestManager() *TestManager {
	ethutil.ReadConfig(".ethtest", "/tmp/ethtest", "ETH")

	db, err := ethdb.NewMemDatabase()
	if err != nil {
		fmt.Println("Could not create mem-db, failing")
		return nil
	}
	ethutil.Config.Db = db

	testManager := &TestManager{}
	testManager.eventMux = new(event.TypeMux)
	testManager.db = db
	// testManager.txPool = NewTxPool(testManager)
	// testManager.blockChain = NewChainManager(testManager)
	// testManager.stateManager = NewStateManager(testManager)

	// Start the tx pool
	testManager.txPool.Start()

	return testManager
}
