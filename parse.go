package zink

import (
	"github.com/pkg/errors"
)

var (
	errIllformedMessage = errors.New("ill-formed message")
	errUnexpectedStatus = errors.New("unexpected status")
)

func parseRawResult(raw rawResult) (Result, error) {
	switch raw.Status {
	case "success":
		message, ok := raw.Message.(string)
		if !ok {
			return nil, errors.Wrapf(errIllformedMessage, "failed to parse %v", raw.Message)
		}
		return &ResultSuccess{
			Message: message,
		}, nil

	case "error":
		rawMessage, ok := raw.Message.([]interface{})
		if !ok {
			return nil, errors.Wrapf(errIllformedMessage, "failed to parse %v", raw.Message)
		}

		message := make([]string, len(rawMessage))
		for i, m := range rawMessage {
			message[i], ok = m.(string)
			if !ok {
				return nil, errors.Wrapf(errIllformedMessage, "failed to parse %v", m)
			}
		}

		return &ResultError{
			Message: message,
		}, nil

	default:
		return nil, errors.Wrapf(errUnexpectedStatus, "failed to parse %v", raw.Status)
	}
}

func parseRawResponses(raw []rawResponse) ([]Response, error) {
	res := make([]Response, len(raw))
	for i, r := range raw {
		result, err := parseRawResult(r.Result)
		if err != nil {
			return nil, err
		}

		res[i] = Response{
			Request: r.Request,
			Result:  result,
		}
	}

	return res, nil
}
