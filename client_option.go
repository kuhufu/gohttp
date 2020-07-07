package flyhttp

import "net/http"

type ClientOption func(cli *GroupClient)

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

func AddCookie(key, val string) ClientOption {
	return func(cli *GroupClient) {
		header := cli.getOrCreateHeader()
		if c := header.Get("Cookie"); c != "" {
			header.Set("Cookie", c+"; "+val)
		} else {
			header.Set("Cookie", val)
		}
	}
}
