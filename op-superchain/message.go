package superchain

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type MessageSafetyLabel int

const (
	Invalid MessageSafetyLabel = iota
	Safe
	Finalized
)

type MessageIdentifier struct {
	Origin      common.Address
	BlockNumber uint64
	LogIndex    uint64
	Timestamp   uint64
	ChainId     uint64
}

type MessagePayload struct {
	// changes: not the entire log
	Log *types.Log
}

func MessagePayloadBytes(log *types.Log) []byte {
	msg := []byte{}
	for _, topic := range log.Topics {
		msg = append(msg, topic.Bytes()...)
	}
	return append(msg, log.Data...)
}
