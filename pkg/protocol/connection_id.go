// SPDX-FileCopyrightText: 2023 The Pion community <https://pion.ly>
// SPDX-License-Identifier: MIT

package protocol

// ConnectionIDData messages are carried by the record layer and wrap an
// encrypted inner payload that contains the real content type and data with
// possible padding.
type ConnectionIDData struct {
	Length   uint16
	RealType ContentType
	Zeros    uint8
	Data     []byte
}

// ContentType returns the ContentType of this content
func (c ConnectionIDData) ContentType() ContentType {
	return ContentTypeApplicationData
}

// Marshal encodes the ConnectionIDData to binary
func (c *ConnectionIDData) Marshal() ([]byte, error) {
	// Size of buffer is content type and data plus size of data.
	out := make([]byte, 1+int(c.Zeros)+len(c.Data))
	out[0] = byte(c.RealType)
	out = append(out, make([]byte, c.Zeros)...)
	return append(out, c.Data...), nil
}

// Unmarshal populates the ConnectionIDData from binary
func (c *ConnectionIDData) Unmarshal(data []byte) error {
	c.RealType = ContentType(data[0])
	c.Zeros = data[1]
	c.Data = append([]byte{}, data[1:]...)
	return nil
}
