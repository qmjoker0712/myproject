package ginrpc

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

const (
	jsonrpcVersion           = "2.0"
	serviceMethodSeparator   = "_"
	subscribeMethodSuffix    = "_subscribe"
	unsubscribeMethodSuffix  = "_unsubscribe"
	notificationMethodSuffix = "_subscription"
)

var (
	MethodNotFoundError = errors.New("Method not found:")
)

type jsonError struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

type jsonRequest struct {
	Method  string          `json:"method"`
	Version string          `json:"jsonrpc"`
	Id      json.RawMessage `json:"id,omitempty"`
	Payload json.RawMessage `json:"params,omitempty"`
}

type jsonSuccessResponse struct {
	Version string      `json:"jsonrpc"`
	Id      interface{} `json:"id,omitempty"`
	Result  interface{} `json:"result"`
}

type jsonErrResponse struct {
	Version string      `json:"jsonrpc"`
	Id      interface{} `json:"id,omitempty"`
	Error   string      `json:"error"`
}

func CreateResponse(id interface{}, reply interface{}) interface{} {
	return &jsonSuccessResponse{Version: jsonrpcVersion, Id: id, Result: reply}
}

// CreateErrorResponse will create a JSON-RPC error response with the given id and error.
func CreateErrorResponse(id interface{}, err error) interface{} {
	return &jsonErrResponse{Version: jsonrpcVersion, Id: id, Error: err.Error()}
}

// parseRequest will parse a single request from the given RawMessage.
// It will return the parsed request and error code
func parseRequest(incomingMsg *jsonRequest) (*rpcRequest, error) {
	// if subscribe method
	if strings.HasSuffix(incomingMsg.Method, subscribeMethodSuffix) {
		return nil, fmt.Errorf("Unsupport websocket now!")
	}
	if strings.HasSuffix(incomingMsg.Method, unsubscribeMethodSuffix) {
		return nil, fmt.Errorf("Unsupport websocket now!")
	}

	elems := strings.Split(incomingMsg.Method, serviceMethodSeparator)
	if len(elems) != 2 {
		return nil, fmt.Errorf(MethodNotFoundError.Error(), incomingMsg.Method)
	}

	if len(incomingMsg.Payload) == 0 {
		return &rpcRequest{
			service: elems[0],
			method:  elems[1],
			id:      &incomingMsg.Id,
		}, nil
	}

	return &rpcRequest{
		service: elems[0],
		method:  elems[1],
		id:      &incomingMsg.Id,
		params:  incomingMsg.Payload,
	}, nil
}

// ParseRequestArguments tries to parse the given params (json.RawMessage) with the given
// types. It returns the parsed values or an error when the parsing failed.
func parseRequestArguments(argTypes []reflect.Type, params interface{}) ([]reflect.Value, error) {
	if args, ok := params.(json.RawMessage); !ok {
		return nil, fmt.Errorf("Invalid params supplied")
	} else {
		return parsePositionalArguments(args, argTypes)
	}
}

// parsePositionalArguments tries to parse the given args to an array of values with the
// given types. It returns the parsed values or an error when the args could not be
// parsed. Missing optional arguments are returned as reflect.Zero values.
func parsePositionalArguments(rawArgs json.RawMessage, types []reflect.Type) ([]reflect.Value, error) {
	// Read beginning of the args array.
	dec := json.NewDecoder(bytes.NewReader(rawArgs))
	if tok, _ := dec.Token(); tok != json.Delim('[') {
		return nil, fmt.Errorf("non-array args")
	}
	// Read args.
	args := make([]reflect.Value, 0, len(types))
	for i := 0; dec.More(); i++ {
		if i >= len(types) {
			return nil, fmt.Errorf("too many arguments, want at most %d", len(types))
		}
		argval := reflect.New(types[i])
		if err := dec.Decode(argval.Interface()); err != nil {
			return nil, fmt.Errorf("invalid argument %d: %v", i, err)
		}
		if argval.IsNil() && types[i].Kind() != reflect.Ptr {
			return nil, fmt.Errorf("missing value for required argument %d", i)
		}
		args = append(args, argval.Elem())
	}
	// Read end of args array.
	if _, err := dec.Token(); err != nil {
		return nil, err
	}
	// Set any missing args to nil.
	for i := len(args); i < len(types); i++ {
		if types[i].Kind() != reflect.Ptr {
			return nil, fmt.Errorf("missing value for required argument %d", i)
		}
		args = append(args, reflect.Zero(types[i]))
	}
	return args, nil
}
