package flyhttp

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func (r Result) existError() bool {
	return r.err != nil || r.resp == nil
}

func (r Result) Bytes() (data []byte, err error) {
	if r.resp != nil {
		defer r.resp.Body.Close()
	}
	if r.existError() {
		return nil, r.err
	}

	return ioutil.ReadAll(r.resp.Body)
}

func (r Result) String() (data string, err error) {
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

func (r Result) Json() (data interface{}, err error) {
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

func (r Result) Resp() (resp *http.Response, err error) {
	resp, err = r.resp, r.err
	return
}
