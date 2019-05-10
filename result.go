package flyhttp

import (
	"encoding/json"
	"io/ioutil"
)

func (r Result) existError() bool {
	return r.err != nil || r.resp == nil
}

func (r Result) Bytes() ([]byte, error) {
	if r.resp != nil {
		defer r.resp.Body.Close()
	}
	if r.existError() {
		return nil, r.err
	}

	return ioutil.ReadAll(r.resp.Body)
}

func (r Result) String() (string, error) {
	if r.resp != nil {
		defer r.resp.Body.Close()
	}
	if r.existError() {
		return "", r.err
	}

	bytes, err := ioutil.ReadAll(r.resp.Body)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func (r Result) Json() (interface{}, error) {
	if r.resp != nil {
		defer r.resp.Body.Close()
	}
	if r.existError() {
		return "", r.err
	}

	bytes, err := ioutil.ReadAll(r.resp.Body)
	if err != nil {
		return nil, err
	}

	var v interface{}
	err = json.Unmarshal(bytes, &v)
	if err != nil {
		return nil, err
	}
	return v, nil
}

