// SPDX-FileCopyrightText: 2023 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package extension

import (
	"golang.org/x/crypto/cryptobyte"
)

// Connection ID is a DTLS extension that provides an alternative to IP address
// and port for session association.
//
// https://tools.ietf.org/html/rfc9146
type ConnectionID struct {
	// Zero-length connection ID indicates that connection ID will be sent but
	// there is no need to respond with one.
	CID []byte // variable length
}

// TypeValue returns the extension TypeValue
func (c ConnectionID) TypeValue() TypeValue {
	return ConnectionIDTypeValue
}

// Marshal encodes the extension
func (c *ConnectionID) Marshal() ([]byte, error) {
	var b cryptobyte.Builder
	b.AddUint16(uint16(c.TypeValue()))
	b.AddUint16LengthPrefixed(func(b *cryptobyte.Builder) {
		b.AddUint8LengthPrefixed(func(b *cryptobyte.Builder) {
			b.AddBytes(c.CID)
		})
	})
	return b.Bytes()
}

// Unmarshal populates the extension from encoded data
func (c *ConnectionID) Unmarshal(data []byte) error {
	val := cryptobyte.String(data)
	var extension uint16
	val.ReadUint16(&extension)
	if TypeValue(extension) != c.TypeValue() {
		return errInvalidExtensionType
	}

	var extData cryptobyte.String
	val.ReadUint16LengthPrefixed(&extData)

	var cid cryptobyte.String
	if !extData.ReadUint8LengthPrefixed(&cid) {
		return errInvalidSNIFormat
	}
	c.CID = make([]byte, len(cid))
	if !cid.CopyBytes(c.CID) {
		return errInvalidSNIFormat
	}
	return nil
}
