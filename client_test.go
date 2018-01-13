package npmdl

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func client() *Client {
	return New()
}

func TestClientErrorBadRequest(t *testing.T) {
	c := client()

	_, err := c.call("/point/asdf/express")

	assert.Error(t, err)

	e := err.(*Error)
	assert.Equal(t, e.Error(), "invalid date")
	assert.Equal(t, "Bad Request", e.Status)
	assert.Equal(t, 400, e.StatusCode)
}
