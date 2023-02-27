package keeper

import (
	"bytes"
	"encoding/gob"
)

func DecodeList(byteArray []byte) []string {
	var list []string
	buf := bytes.NewBuffer(byteArray)
	dec := gob.NewDecoder(buf)
	err := dec.Decode(&list)
	if err != nil {
		panic(err)
	}
	return list
}

func EncodeList(list []string) []byte {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(list)
	if err != nil {
		panic(err)
	}
	return buf.Bytes()
}
