package zink

type Request struct {
	URL        string  `json:"url"`
	Keyword    *string `json:"keyword,omitempty"`
	Collection *string `json:"collection,omitempty"`
}

type rawResult struct {
	Status  string      `json:"status"`
	Message interface{} `json:"message"`
}

type rawResponse struct {
	Request
	Result rawResult `json:"result"`
}

type ResultSuccess struct {
	Message string
}

type ResultError struct {
	Message []string
}

type Result interface{}

type Response struct {
	Request
	Result Result
}
