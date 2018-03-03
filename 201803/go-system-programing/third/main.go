package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

func main() {
	endian()
}

func endian() {
	// 10000 x ビッグエンディアン
	data := []byte{0x0, 0x0, 0x27, 0x10}
	var i int32
	binary.Read(bytes.NewReader(data), binary.BigEndian, &i)
	fmt.Printf("data: %d\n", i)
}
