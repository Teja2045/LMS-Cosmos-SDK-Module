package keeper

import (
	"bytes"
	"encoding/gob"
	"fmt"
)

func DecodeList(byteArray []byte) []string {
	var list []string
	buf := bytes.NewBuffer(byteArray)
	dec := gob.NewDecoder(buf)
	err := dec.Decode(&list)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(list)
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
