// Copyright 2021 Kirill Scherba <kirill@scherba.ru>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// ByteSlice package help binary marshal/unmarshal small strings and byte slice
package bslice

import (
	"bytes"
	"encoding/binary"
)

// ByteSlice help binary marshal/ubmarshal byte slice
type ByteSlice struct{}

// WriteSlice return binary representation of byte slice
func (b ByteSlice) WriteSlice(buf *bytes.Buffer, data []byte) (err error) {
	if err = binary.Write(buf, binary.LittleEndian, uint16(len(data))); err != nil {
		return
	}
	err = binary.Write(buf, binary.LittleEndian, data)
	return
}

// ReadSlice read binary byte slice created with WriteSlice function
func (b ByteSlice) ReadSlice(buf *bytes.Buffer) (data []byte, err error) {
	var l uint16
	if err = binary.Read(buf, binary.LittleEndian, &l); err != nil {
		return
	}
	data = make([]byte, l)
	err = binary.Read(buf, binary.LittleEndian, data)
	return
}

// ReadString read binary string created with WriteSlice function
func (b ByteSlice) ReadString(buf *bytes.Buffer) (data string, err error) {
	d, err := b.ReadSlice(buf)
	if err != nil {
		return
	}
	data = string(d)
	return
}

// WriteStringSlice return binary representation of string slice
func (b ByteSlice) WriteStringSlice(buf *bytes.Buffer, data []string) (err error) {
	if err = binary.Write(buf, binary.LittleEndian, uint16(len(data))); err != nil {
		return
	}
	for i := range data {
		if err = b.WriteSlice(buf, []byte(data[i])); err != nil {
			return
		}
	}

	return
}

// ReadString read binary string slice created with WriteStringSlice function
func (b ByteSlice) ReadStringSlice(buf *bytes.Buffer) (data []string, err error) {
	var l uint16
	if err = binary.Read(buf, binary.LittleEndian, &l); err != nil {
		return
	}
	for i := 0; i < int(l); i++ {
		var d []byte
		if d, err = b.ReadSlice(buf); err != nil {
			return
		}
		data = append(data, string(d))
	}
	return
}
