package zenvia

import (
	"net/http"
)

const (
	PRODUCTION = "https://api.zenvia.com"
	SANDBOX    = "https://TODO"
)

var (
	Client HTTPClient
)

type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

type ZenviaAccess struct {
	api      string
	from     string
	to       string
	apiToken string
}

func ZenviaConfig() ZenviaAccess {
	return ZenviaAccess{
		api: PRODUCTION,
	}
}

func (zenviaAccess ZenviaAccess) ApiToken(apiToken string) ZenviaAccess {
	zenviaAccess.apiToken = apiToken

	return zenviaAccess
}

func (zenviaAccess ZenviaAccess) From(from string) ZenviaAccess {
	zenviaAccess.from = from

	return zenviaAccess
}

func (zenviaAccess ZenviaAccess) To(to string) ZenviaAccess {
	zenviaAccess.to = to

	return zenviaAccess
}

func (zenviaAccess ZenviaAccess) Sandbox() ZenviaAccess {
	zenviaAccess.api = SANDBOX

	return zenviaAccess
}

type ZenviaAccessRepository interface {
	SendSMS(message string)
}
