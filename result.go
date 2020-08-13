package gohttp

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Result struct {
	resp *http.Response
	err  error
}

func (r Result) Bytes() (data []byte, err error) {
	resp, err := r.resp, r.err
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return ioutil.ReadAll(resp.Body)
}

func (r Result) String() (data string, err error) {
	resp, err := r.resp, r.err
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (r Result) JSON(v interface{}) error {
	resp, err := r.resp, r.err
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, v)
	if err != nil {
		return err
	}
	return nil
}

func (r Result) Raw() (resp *http.Response, err error) {
	return r.resp, r.err
}

func (r Result) Err() (err error) {
	return r.err
}
