// Copyright 2018 The Gringo Developers. All rights reserved.
// Use of this source code is governed by a GNU GENERAL PUBLIC LICENSE v3
// license that can be found in the LICENSE file.

package p2p

import (
	"consensus"
	"bytes"
	"encoding/binary"
	"github.com/sirupsen/logrus"
	"io"
)

type Locator struct {
	Hashes []consensus.Hash
}

// Bytes implements Message interface
func (h *Locator) Bytes() []byte {
	buff := new(bytes.Buffer)

	// FIXME: should check the bounds & set the limits
	if err := binary.Write(buff, binary.BigEndian, uint8(len(h.Hashes))); err != nil {
		logrus.Fatal(err)
	}

	for _, hash := range h.Hashes {
		if _, err := buff.Write(hash); err != nil {
			logrus.Fatal(err)
		}
	}

	return buff.Bytes()
}

// Read implements Message interface
func (h *Locator) Read(r io.Reader) error {

	var count uint8
	if err := binary.Read(r, binary.BigEndian, &count); err != nil {
		return err
	}

	h.Hashes = make([]consensus.Hash, count)
	for i := 0; i < int(count); i++ {
		h.Hashes[i] = make([]byte, consensus.BlockHashSize)

		if _, err := io.ReadFull(r, h.Hashes[i]); err != nil {
			return err
		}
	}

	return nil
}