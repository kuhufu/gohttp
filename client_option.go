package flyhttp

import "net/http"

type ClientOption func(cli *GroupClient)

func WithHost(host string) ClientOption {
	return func(cli *GroupClient) {
		cli.host = host
	}
}

func WithHttpClient(client *http.Client) ClientOption {
	return func(cli *GroupClient) {
		cli.cli = client
	}
}

func WithBase(base string) ClientOption {
	return func(cli *GroupClient) {
		cli.base = base
	}
}

func WithHeader(header http.Header) ClientOption {
	return func(cli *GroupClient) {
		cli.header = header
	}
}

func SetHeader(key, val string) ClientOption {
	return func(cli *GroupClient) {
		cli.getOrCreateHeader().Set(key, val)
	}
}

func AddHeader(key, val string) ClientOption {
	return func(cli *GroupClient) {
		cli.getOrCreateHeader().Add(key, val)
	}
}
