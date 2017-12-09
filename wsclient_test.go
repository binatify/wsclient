package wsclient

import (
	"testing"

	"github.com/golib/assert"
)

func TestNewClient(t *testing.T) {
	assertion := assert.New(t)

	client := NewClient(nil)

	assertion.NotNil(client)
}

func TestToPayload(t *testing.T) {
	assertion := assert.New(t)

	type serviceStation struct{}

	var in serviceStation

	r, err := toPayload(in)
	assertion.Nil(err)
	assertion.NotNil(r)
}

func TestDo(t *testing.T) {
	assertion := assert.New(t)

	type serviceStation struct{}

	var in serviceStation

	client := NewClient(nil)

	_, err := client.Do(in)

	assertion.Nil(err)
}
