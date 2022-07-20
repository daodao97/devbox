package http

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	resty "github.com/go-resty/resty/v2"
	"github.com/spf13/cast"
)

type Http struct{}

type Options struct {
	Url        string                 `json:"url"`
	Method     string                 `json:"method"`
	Headers    map[string]interface{} `json:"headers"`
	Body       map[string]interface{} `json:"body"`
	QueryParam map[string]interface{} `json:"query_params"`
	FormData   map[string]string      `json:"form_data"`
}

type Response struct {
	Headers http.Header
	Body    string
}

func (h Http) Request(opt *Options) (*Response, error) {

	if opt.Url == "" {
		return nil, errors.New("request url is empty")
	}

	client := resty.New().SetTimeout(time.Second).R()
	for k, v := range opt.Headers {
		client.SetHeader(k, cast.ToString(v))
	}

	if opt.Body != nil {
		client.SetBody(opt.Body)
	}

	if opt.Method == "" {
		opt.Method = http.MethodGet
	}

	mm := map[string]func(url string) (*resty.Response, error){
		http.MethodGet:     client.Get,
		http.MethodPost:    client.Post,
		http.MethodPut:     client.Put,
		http.MethodHead:    client.Head,
		http.MethodPatch:   client.Patch,
		http.MethodOptions: client.Options,
		http.MethodDelete:  client.Delete,
	}

	req, ok := mm[opt.Method]
	if !ok {
		return nil, errors.New("not support method")
	}

	resp, err := req(opt.Url)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode() != http.StatusOK {
		return nil, fmt.Errorf("http code error %d", resp.StatusCode())
	}

	resp.Header()

	return &Response{
		Headers: resp.Header(),
		Body:    string(resp.Body()),
	}, nil
}
