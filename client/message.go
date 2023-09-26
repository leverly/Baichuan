package client

type ReqMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type RequestMessage struct {
	Id        string
	Model     string       `json:"model"`
	Messages  []ReqMessage `json:"messages"`
	Parameter ReqParameter `json:"parameters,omitempty"`
}

type ReqParameter struct {
	Temperature         float64 `json:"temperature,omitempty"`
	Top_k               int     `json:"top_k,omitempty"`
	Top_p               float64 `json:"top_p,omitempty"`
	With_search_enhance bool    `json:"with_search_enhance,omitempty"`
}

type ResponseMessage struct {
	Code  int        `json:"code"`
	Msg   string     `json:"msg"`
	Data  RespData   `json:"data"`
	Usage TokenUsage `json:"usage"`
}

type RespMessage struct {
	Role          string `json:"role"`
	Content       string `json:"content"`
	Finish_reason string `json:"finish_reason"`
}

type RespData struct {
	Messages []RespMessage `json:"messages"`
}

type TokenUsage struct {
	Prompt_tokens int `json:"prompt_tokens"`
	Answer_tokens int `json:"answer_tokens"`
	Total_tokens  int `json:"total_tokens"`
}
