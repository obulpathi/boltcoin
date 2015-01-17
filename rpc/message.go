package rpc

import "github.com/obulpathi/boltcoin/ethutil"

type Message struct {
	Call string        `json:"call"`
	Args []interface{} `json:"args"`
	Id   int           `json:"_id"`
	Data interface{}   `json:"data"`
}

func (self *Message) Arguments() *ethutil.Value {
	return ethutil.NewValue(self.Args)
}
