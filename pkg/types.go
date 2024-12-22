package pkg

type ApiRoute struct {
	Nickname   string `json:"nickname"`
	Group      string `json:"group"`
	Route      string `json:"route"`
	HTTPMethod string `json:"http_method"`
}

type Config struct {
	RootApiPath    string `json:"root_api_path"`
	MaxEchoTimeout int    `json:"max_echo_timeout"`
}

type ResponseContent struct {
	StatusCode int                    `json:"statusCode"`
	Body       map[string]interface{} `json:"body"`
	Headers    map[string]string      `json:"headers"`
}
