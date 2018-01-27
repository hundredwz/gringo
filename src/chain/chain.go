// Copyright 2018 The Gringo Developers. All rights reserved.
// Use of this source code is governed by a GNU GENERAL PUBLIC LICENSE v3
// license that can be found in the LICENSE file.

package chain

import (
	"consensus"
	"sync"
	"container/list"
	"bytes"
	"time"
	"github.com/sirupsen/logrus"
)

// Testnet1 genesis block
var Testnet1 = consensus.Block{
	Header: consensus.BlockHeader{
		Version:         1,
		Height:          0,
		Previous:        bytes.Repeat([]byte{0xff}, consensus.BlockHashSize),
		Timestamp:       time.Date(2017, 11, 16, 20, 0, 0, 0, time.UTC),
		Difficulty:      10,
		TotalDifficulty: 10,

		UTXORoot:       bytes.Repeat([]byte{0x00}, 32),
		RangeProofRoot: bytes.Repeat([]byte{0x00}, 32),
		KernelRoot:     bytes.Repeat([]byte{0x00}, 32),

		Nonce: 28205,
		POW: consensus.Proof{
			Nonces: []uint32{
				0x21e, 0x7a2, 0xeae, 0x144e, 0x1b1c, 0x1fbd,
				0x203a, 0x214b, 0x293b, 0x2b74, 0x2bfa, 0x2c26,
				0x32bb, 0x346a, 0x34c7, 0x37c5, 0x4164, 0x42cc,
				0x4cc3, 0x55af, 0x5a70, 0x5b14, 0x5e1c, 0x5f76,
				0x6061, 0x60f9, 0x61d7, 0x6318, 0x63a1, 0x63fb,
				0x649b, 0x64e5, 0x65a1, 0x6b69, 0x70f8, 0x71c7,
				0x71cd, 0x7492, 0x7b11, 0x7db8, 0x7f29, 0x7ff8,
			},
		},
	},
}

// Testnet2 genesis block
var Testnet2 = consensus.Block{
	Header: consensus.BlockHeader{
		Version:   1,
		Height:    0,
		Previous:  bytes.Repeat([]byte{0xff}, consensus.BlockHashSize),
		Timestamp: time.Date(2017, 11, 16, 20, 0, 0, 0, time.UTC),
		//Difficulty:      10,
		//TotalDifficulty: 10,

		UTXORoot:       bytes.Repeat([]byte{0x00}, 32),
		RangeProofRoot: bytes.Repeat([]byte{0x00}, 32),
		KernelRoot:     bytes.Repeat([]byte{0x00}, 32),

		Nonce: 70081,
		POW: consensus.Proof{
			Nonces: []uint32{
				0x43ee48, 0x18d5a49, 0x2b76803, 0x3181a29, 0x39d6a8a, 0x39ef8d8,
				0x478a0fb, 0x69c1f9e, 0x6da4bca, 0x6f8782c, 0x9d842d7, 0xa051397,
				0xb56934c, 0xbf1f2c7, 0xc992c89, 0xce53a5a, 0xfa87225, 0x1070f99e,
				0x107b39af, 0x1160a11b, 0x11b379a8, 0x12420e02, 0x12991602, 0x12cc4a71,
				0x13d91075, 0x15c950d0, 0x1659b7be, 0x1682c2b4, 0x1796c62f, 0x191cf4c9,
				0x19d71ac0, 0x1b812e44, 0x1d150efe, 0x1d15bd77, 0x1d172841, 0x1d51e967,
				0x1ee1de39, 0x1f35c9b3, 0x1f557204, 0x1fbf884f, 0x1fcf80bf, 0x1fd59d40,
			},
		},
	},
}

// Mainnet genesis block
var Mainnet = consensus.Block{
	Header: consensus.BlockHeader{
		Version:         1,
		Height:          0,
		Previous:        bytes.Repeat([]byte{0xff}, consensus.BlockHashSize),
		Timestamp:       time.Date(2018, 8, 14, 0, 0, 0, 0, time.UTC),
		Difficulty:      1000,
		TotalDifficulty: 1000,

		UTXORoot:       bytes.Repeat([]byte{0x00}, 32),
		RangeProofRoot: bytes.Repeat([]byte{0x00}, 32),
		KernelRoot:     bytes.Repeat([]byte{0x00}, 32),

		Nonce: 28205,
		POW: consensus.Proof{
			Nonces: []uint32{
				0x21e, 0x7a2, 0xeae, 0x144e, 0x1b1c, 0x1fbd,
				0x203a, 0x214b, 0x293b, 0x2b74, 0x2bfa, 0x2c26,
				0x32bb, 0x346a, 0x34c7, 0x37c5, 0x4164, 0x42cc,
				0x4cc3, 0x55af, 0x5a70, 0x5b14, 0x5e1c, 0x5f76,
				0x6061, 0x60f9, 0x61d7, 0x6318, 0x63a1, 0x63fb,
				0x649b, 0x64e5, 0x65a1, 0x6b69, 0x70f8, 0x71c7,
				0x71cd, 0x7492, 0x7b11, 0x7db8, 0x7f29, 0x7ff8,
			},
		},
	},
}

type Chain struct {
	sync.RWMutex

	// Storage of blockchain
	storage Storage

	// genesis block
	genesis consensus.Block
	// current height of chain
	height uint64
	// current total difficulty
	totalDifficulty consensus.Difficulty

	// list of blockheaders
	blockHashChain *list.List
}

func New(genesis consensus.Block, storage Storage) *Chain {
	chain := Chain{
		storage:         storage,
		genesis:         genesis,
		height:          genesis.Header.Height,
		totalDifficulty: genesis.Header.TotalDifficulty,
		blockHashChain:  list.New(),
	}

	// init header list with genesis header
	chain.blockHashChain.PushFront(genesis.Hash())

	return &chain
}

// Genesis returns genesis block
func (c *Chain) Genesis() consensus.Block {
	return c.genesis
}

// TotalDifficulty returns current total difficulty
func (c *Chain) TotalDifficulty() consensus.Difficulty {
	c.RLock()
	defer c.RUnlock()

	return c.totalDifficulty
}

// Height returns current height
func (c *Chain) Height() uint64 {
	c.RLock()
	defer c.RUnlock()

	return c.height
}

// GetBlockHeaders returns block headers
func (c *Chain) GetBlockHeaders(loc consensus.Locator) []consensus.BlockHeader {
	// for safety
	if len(loc.Hashes) > consensus.MaxLocators {
		logrus.Error("locator hashes object is too big")
		loc.Hashes = loc.Hashes[:consensus.MaxLocators]
	}

	result := make([]consensus.BlockHeader, 0)
	c.RLock()
	defer c.RUnlock()

	for el := c.blockHashChain.Back(); el != nil; el = el.Prev() {
		blockHash, ok := el.Value.(consensus.Hash)
		if !ok {
			logrus.Fatal("unexpected errors with list of blockHashes")
		}

		for _, hash := range loc.Hashes {
			if bytes.Compare(blockHash, hash) == 0 {
				// Founded, send the next blocks hashes
				blockHash, ok := el.Next().Value.(consensus.Hash)
				if !ok {
					logrus.Fatal("unexpected errors with list of blockHashes")
				}

				blockID := BlockID{
					Hash: &blockHash,
					Height: nil,
				}

				// get blocks from
				blockList, err := c.storage.From(blockID, consensus.MaxBlockHeaders)
				if err != nil {
					logrus.Error(err)
				}

				// collect headers
				for _, block := range blockList {
					result = append(result, block.Header)
				}

				return result
			}
		}
	}

	return result
}

// GetBlock returns block by hash, if not found returns nil, nil
func (c *Chain) GetBlock(hash consensus.Hash) (*consensus.Block, error) {
	if hash == nil {
		return nil, nil
	}

	return c.storage.GetBlock(BlockID{
		Hash:   &hash,
		Height: nil,
	})
}

// GetBlockID returns block by hash, height or both
func (c *Chain) GetBlockID(b BlockID) (*consensus.Block, error) {
	return c.storage.GetBlock(b)
}

func (c *Chain) ProcessHeaders(headers []consensus.BlockHeader) error {
	return nil
}

func (c *Chain) ProcessBlock(block *consensus.Block) error {
	// befora locking block chain on change MUST lock blockHashesChain chain
	return nil
}
