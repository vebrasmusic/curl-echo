package pkg

type ApiRoute struct {
	Nickname   string `json:"nickname"`
	Group      string `json:"group"`
	Route      string `json:"route"`
	HTTPMethod string `json:"http_method"`
}
