package gohttp

import (
	"net/http"
)

type GroupClient struct {
	cli    *http.Client //std http client
	header http.Header  //public header
	base   string       //base url
}

func New(opts ...ClientOption) *GroupClient {
	cli := &GroupClient{
		cli: http.DefaultClient,
	}

	for _, opt := range opts {
		opt(cli)
	}

	return cli
}

func (cli *GroupClient) copy() *GroupClient {
	newCli := *cli
	newCli.header = cli.header.Clone()

	return &newCli
}

func (cli *GroupClient) getOrCreateHeader() http.Header {
	if cli.header == nil {
		cli.header = http.Header{}
	}

	return cli.header
}

func (cli *GroupClient) Header() http.Header {

	return cli.getOrCreateHeader().Clone()
}

func (cli *GroupClient) Group(relativePath string, opts ...ClientOption) *GroupClient {
	newCli := cli.copy()
	newCli.base = newCli.base + relativePath

	for _, opt := range opts {
		opt(newCli)
	}

	return newCli
}

func (cli *GroupClient) buildUrl(url string) string {
	if cli.base == "" {
		return url
	} else {
		return cli.base + url
	}
}

func (cli *GroupClient) Do(method string, url string, opts ...RequestOption) (*http.Response, error) {
	req, err := http.NewRequest(method, cli.buildUrl(url), nil)
	if err != nil {
		return nil, err
	}

	req.Header = cli.header.Clone()

	for _, opt := range opts {
		opt(req)
	}

	resp, err := cli.cli.Do(req)
	if err != nil {
		return nil, err
	}

	return resp, err
}

func (cli *GroupClient) Get(url string, opts ...RequestOption) (*http.Response, error) {
	return cli.Do("GET", url, opts...)
}

func (cli *GroupClient) Post(url string, opts ...RequestOption) (*http.Response, error) {
	return cli.Do("POST", url, opts...)
}
