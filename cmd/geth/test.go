// 测试rlp编码解码
package main

import (
	"fmt"

	"github.com/ethereum/go-ethereum/rlp"
)

type RecorderMessage struct {
	Peer1_deliver_time string
	Peer2_receive_time string
	Peer2_deliver_time string
	Peer1_receive_time string
}

func main1() {
	var a RecorderMessage
	a.Peer1_deliver_time = "123"
	a.Peer2_deliver_time = "456"
	_, r, _ := rlp.EncodeToReader(a)
	fmt.Println(r)
	fmt.Println(a)
	var b RecorderMessage
	if err := rlp.Decode(r, &b); err != nil {
		fmt.Println("Error decoding message:", err)
	}
	fmt.Println(b)
	fmt.Println(b.Peer1_deliver_time)
	fmt.Println(b.Peer1_receive_time)
	fmt.Println(b.Peer2_deliver_time)
	fmt.Println(b.Peer1_receive_time)
}
