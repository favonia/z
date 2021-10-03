package zink

type Request struct {
	URL        string  `json:"url"`
	Keyword    *string `json:"keyword,omitempty"`
	Collection *string `json:"collection,omitempty"`
}

type Requests = []Request

type Status string

type rawResult struct {
	Status  string      `json:"status"`
	Message interface{} `json:"message"`
}

type rawResponse struct {
	Request
	Result rawResult `json:"result"`
}

type rawResponses = []rawResponse

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

type Responses = []Response
