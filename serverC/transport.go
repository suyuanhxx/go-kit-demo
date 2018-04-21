package main

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/go-kit/kit/endpoint"
	"fmt"
)

func makeGetUserByUserNameEndpoint(svc UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(getUserByUserNameRequest)
		v := svc.GetUserByUserName(req.S)
		return v, nil
	}
}

func decodeGetUserByUserNameRequest(_ context.Context, r *http.Request) (interface{}, error) {
	var request getUserByUserNameRequest
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		return nil, err
	}
	return request, nil
}

func decodeGetUserByUserNameResponse(_ context.Context, r *http.Response) (interface{}, error) {
	fmt.Println(r.Body)
	var response []byte
	if err := json.NewDecoder(r.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil

}

//func encodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
//	return json.NewEncoder(w).Encode(response)
//}

func encodeResponse(_ context.Context, obj interface{}) (json.RawMessage, error) {
	b, err := json.Marshal(obj)
	if err != nil {
		return nil, fmt.Errorf("couldn't marshal response: %s", err)
	}
	return b, nil
}

func encodeRequest(_ context.Context, r *http.Request, request interface{}) error {
	var buf bytes.Buffer
	if err := json.NewEncoder(&buf).Encode(request); err != nil {
		return err
	}
	r.Body = ioutil.NopCloser(&buf)
	return nil
}

type getUserByUserNameRequest struct {
	S string `json:"s"`
}

//type getUserByUserNameRequestResponse struct {
//	V   string `json:"v"`
//	Err string `json:"err,omitempty"`
//}
