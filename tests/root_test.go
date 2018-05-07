package tests

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"testing-in-golang/framework/swapi"
)

func Test_GetRoot_200(t *testing.T) {
	t.Parallel()
	body := swapi.GetUrls()
	assert.Contains(t, body.People, "http")
}
