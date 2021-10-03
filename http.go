package zink

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/lestrrat-go/jwx/jwa"
	"github.com/lestrrat-go/jwx/jwt"
	"github.com/pkg/errors"
)

type Handle struct {
	AccessID    string
	SecretKey   string
	AccessPoint string
}

const DefaultAccessPoint = "https://z.umn.edu/api/v1/urls"

func New(id string, secret string) *Handle {
	return &Handle{
		AccessID:    id,
		SecretKey:   secret,
		AccessPoint: DefaultAccessPoint,
	}
}

func (h *Handle) Do(ctx context.Context, reqs []Request) ([]Response, error) {
	payload := jwt.New()
	if err := payload.Set("urls", reqs); err != nil {
		return nil, errors.Wrapf(err, "failed to prepare the jwt payload")
	}

	token, err := jwt.Sign(payload, jwa.HS256, []byte(h.SecretKey))
	if err != nil {
		return nil, errors.Wrapf(err, "failed to sign the jwt payload")
	}

	httpReq, err := http.NewRequestWithContext(ctx, http.MethodPost, h.AccessPoint, nil)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to prepare the HTTP request")
	}

	httpReq.Header.Set("Authorization", fmt.Sprintf("%s:%s", h.AccessID, token))

	resp, err := http.DefaultClient.Do(httpReq)
	if err != nil {
		return nil, errors.Wrapf(err, "failed to make the HTTP request")
	}
	defer resp.Body.Close()

	var rawRes []rawResponse
	if err = json.NewDecoder(resp.Body).Decode(&rawRes); err != nil {
		return nil, errors.Wrapf(err, "failed to decode the HTTP response")
	}

	res, err := parseRawResponses(rawRes)
	if err != nil {
		return nil, err
	}

	return res, nil
}
