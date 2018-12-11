package rpc

import (
	"fmt"
)

// D4DRPCClient d4d rpc client for contract service
type D4DRPCClient struct {
	url    string
	scheme string
	Debug  bool
}

// New create new json_rpc client with given url
func newRPC(url string, options ...func(rpc *D4DRPCClient)) *D4DRPCClient {
	rpc := &D4DRPCClient{
		url:    url,
		scheme: "http",
	}
	for _, option := range options {
		option(rpc)
	}
	return rpc
}

// NewD4DRPCClient create new json_rpc
func NewD4DRPCClient(url string, options ...func(rpc *D4DRPCClient)) *D4DRPCClient {
	return newRPC(url, options...)
}

func (rpc *D4DRPCClient) call(serviceMethod string, args interface{}, reply interface{}) error {
	rawURL := fmt.Sprintf("%s://%s", rpc.scheme, rpc.url)
	conn, err := Dial(rawURL)
	defer func() {
		if conn != nil {
			conn.Close()
		}
	}()

	if err != nil {
		return err
	}

	err = conn.Call(&reply, serviceMethod, args)
	if err != nil {
		return err
	}
	if rpc.Debug {
		log.Debug("%s\nRequest: %v\nResponse: %v\n", serviceMethod, args, &reply)
	}
	return nil
}

func callRPC(rawURL string, serviceMethod string, args interface{}, reply interface{}) error {
	conn, err := Dial(rawURL)
	defer func() {
		if conn != nil {
			conn.Close()
		}
	}()

	if err != nil {
		return err
	}

	err = conn.Call(&reply, serviceMethod, args)
	if err != nil {
		return err
	}
	log.Debug("callRPC response from %v", rawURL)
	return nil
}

// CallRPC call rpc without struct
func CallRPC(rawURL string, serviceMethod string, args interface{}, reply interface{}) error {
	return callRPC(rawURL, serviceMethod, args, reply)
}

// CallRPCBatch call rpc batch without struct
func CallRPCBatch(rawURL string, batch []BatchElem) error {
	conn, err := Dial(rawURL)
	defer func() {
		if conn != nil {
			conn.Close()
		}
	}()
	if err != nil {
		return err
	}
	return conn.BatchCall(batch)
}
