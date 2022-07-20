package http

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRequest(t *testing.T) {
	c := Http{}

	opt := &Options{
		Url:     "http://google.com",
		Method:  "POST",
		Body:    map[string]interface{}{},
		Headers: map[string]interface{}{},
	}

	ret, err := c.Request(opt)
	assert.Equal(t, nil, err)
	fmt.Println(ret)
}
