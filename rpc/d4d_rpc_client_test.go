package rpc

import (
	"fmt"
	"testing"
)

func Test_RPC(t *testing.T) {
	url := "127.0.0.1:8548"
	rpc := NewD4DRPCClient(url)
	var result interface{}
	rpc.call("monitor_pong", nil, &result)
	fmt.Printf("result is %#v\n", result)
}

func Test_CALLRPC(t *testing.T) {
	url := "http://127.0.0.1:8548"
	var result interface{}
	CallRPC(url, "monitor_pong", nil, &result)
	fmt.Printf("result is %#v\n", result)
}

func Test_CallRPCBatch(t *testing.T) {
	url := "http://127.0.0.1:8548"
	result1 := new(interface{})
	result2 := new(interface{})
	batch := []BatchElem{
		{
			Method: "monitor_ping",
			Args:   []interface{}{},
			Result: &result1,
		},
		{
			Method: "monitor_pong",
			Args:   []interface{}{},
			Result: &result2,
		},
	}
	CallRPCBatch(url, batch)
	fmt.Printf("result1 is %#v\n", *result1)
	fmt.Printf("result2 is %#v\n", *result2)
}
